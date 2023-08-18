package arangodb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	client "github.com/boorad/benthos-arangodb/client"
	arango "github.com/arangodb/go-driver"

	"github.com/benthosdev/benthos/v4/public/bloblang"
	"github.com/benthosdev/benthos/v4/public/service"
)

const (
	// client config
	configKeyUrl           = "url"
	configKeyUsername      = "username"
	configKeyPassword      = "password"
	configKeyDatabase      = "database"
	configKeyCollection    = "collection"
	configKeyCollectionMap = "collection_map"

	// processor config
	configKeyOperation    = "operation"
	configKeyParams       = "params"
	configKeyFilterMap    = "filter_map"
	configKeyDocumentMap  = "document_map"
	configKeyQuery        = "query"
	configKeyQueryAql     = "aql"
	configKeyQueryVarsMap = "vars_map"
)

func init() {

	configSpec := arangoProcessorConfigSpec()

	constructor := func(conf *service.ParsedConfig, mgr *service.Resources) (service.BatchProcessor, error) {
		var (
			client arango.Client
			config *Config
			db     arango.Database
			err    error
		)

		ctx := context.Background() // TODO: with timeout?

		if config, err = getProcessorConfig(conf); err != nil {
			return nil, err
		}

		if client, err = config.Client.Client(); err != nil {
			return nil, err
		}

		if db, err = client.Database(ctx, config.Client.Database); err != nil {
			return nil, err
		}

		return newArangoDBProcessor(config, client, db, mgr.Logger(), mgr.Metrics())
	}

	if err := service.RegisterBatchProcessor("arangodb", configSpec, constructor); err != nil {
		panic(err)
	}
}

func arangoProcessorConfigSpec() *service.ConfigSpec {
	return service.NewConfigSpec().
		Field(service.NewStringField(configKeyUrl).
			Description("The URL of the target ArangoDB instance.")).
		Field(service.NewStringField(configKeyUsername).
			Description("The username for the target ArangoDB instance.")).
		Field(service.NewStringField(configKeyPassword).
			Description("The password for the target ArangoDB instance.")).
		Field(service.NewStringField(configKeyDatabase).
			Description("The database to use from the target ArangoDB instance.")).
		Field(service.NewStringField(configKeyCollection).
			Description("The collection to use from the target ArangoDB instance.").
			Optional()).
		Field(service.NewBloblangField(configKeyCollectionMap).
			Description("The collection map to use from the target ArangoDB instance. This can dynamically replace the collection name from values of the current doc (this).").
			Optional()).
		Field(service.NewStringEnumField(configKeyOperation, Operations...).
			Description(fmt.Sprintf("The operation to perform: %v", Operations))).
		Field(service.NewAnyMapField(configKeyParams).
			Description("Operation-specific parameters.  ex: the create operator's `overwrite: true` parameter").
			Optional()).
		Field(service.NewBloblangField(configKeyFilterMap).
			Description("A mapping representing the filter for the ArangoDB command.  NOT CURRENTLY OPERATIONAL").
			Optional()).
		Field(service.NewBloblangField(configKeyDocumentMap).
			Description("A mapping representing the document that will be used for the current operation. Maps fields in the current document to target ArangoDB fields.").
			Optional()).
		Field(service.NewObjectField(configKeyQuery,
			service.NewStringField(configKeyQueryAql).
				Description("AQL Query").
				Optional(),
			service.NewBloblangField(configKeyQueryVarsMap).
				Description("Add optional variables for query binding.  This is a mapping with access to the current doc (this).").
				Optional(),
		).
			Description("AQL and Variables for the query operator").
			Optional())
}

type arangoDBProcessor struct {
	config Config
	client arango.Client
	db     arango.Database
	logger *service.Logger
}

func newArangoDBProcessor(
	config *Config,
	client arango.Client,
	db arango.Database,
	logger *service.Logger,
	_metrics *service.Metrics,
) (*arangoDBProcessor, error) {

	return &arangoDBProcessor{
		config: *config,
		client: client,
		db:     db,
		logger: logger,
	}, nil
}

// Document and Variables struct
// (doc is for all operators, vars are for query operator)
type DV struct {
	doc  *service.Message
	vars *service.Message
}

func (p *arangoDBProcessor) ProcessBatch(ctx context.Context, batch service.MessageBatch) ([]service.MessageBatch, error) {

	var outBatch []service.MessageBatch

	// key: collection name, value: array of document & values structs on which to perform the operation
	docs := make(map[string][]*DV)
	vars := service.NewMessage(nil)

	for i, msg := range batch {
		var (
			collName string
			err      error
		)

		doc := msg

		if isDocumentAllowed(p.config.Operation) && p.config.DocumentMap != nil {
			res, err := batch.BloblangQuery(i, p.config.DocumentMap)
			if err != nil {
				p.logger.Debugf("document_map mapping failed: %v", err)
				msg.SetError(err)
				continue
			}
			doc = res
		} else if p.config.Operation == OperationQuery && p.config.Query.VarsMap != nil {
			res, err := batch.BloblangQuery(i, p.config.Query.VarsMap)
			if err != nil {
				p.logger.Debugf("query vars_map mapping failed: %v", err)
				msg.SetError(err)
				continue
			}
			vars = res
		}

		if p.config.Client.Collection != "" || p.config.Client.CollectionMap != nil {
			if collName, err = p.getArangoCollectionName(ctx, i, batch); err != nil {
				msg.SetError(err)
				continue
			}
		}
		docs[collName] = append(docs[collName], &DV{doc, vars})
	}

	for collName, collDVs := range docs {
		res, err := p.processCollectionDocs(ctx, collName, collDVs, p.config.Operation)
		if err != nil {
			p.logger.Errorf("error performing operation '%s' on collection '%s': %s", p.config.Operation, collName, err)
			// msg.SetError(err)
			continue
		}
		outBatch = append(outBatch, res)
	}
	p.logger.Debugf("batch complete: %d", len(batch))
	return outBatch, nil
}

func (p *arangoDBProcessor) Close(ctx context.Context) error {
	return nil
}

type createResp struct {
	Key    string            `json:"_key,omitempty"`
	ID     arango.DocumentID `json:"_id,omitempty"`
	Rev    string            `json:"_rev,omitempty"`
	OldRev string            `json:"_oldRev,omitempty"`
	NEW    any
}

func (p *arangoDBProcessor) processCollectionDocs(
	ctx context.Context, collName string, dvs []*DV, op string) (service.MessageBatch, error) {
	var (
		coll   arango.Collection
		cursor arango.Cursor
		meta   arango.DocumentMetaSlice
		docs   []any
		vars   []any
		errs   arango.ErrorSlice
		err    error
	)

	if collName != "" {
		coll, err = p.getArangoCollection(ctx, collName)
		if err != nil {
			p.logger.Errorf("error getting arangodb collection: %s", collName)
			return nil, err
		}
	}

	// Get structured doc & vars for arangodb driver.
	// If one fails, they both should, as they're tied together.
	for _, dv := range dvs {
		d, derr := dv.doc.AsStructured()
		v, verr := dv.vars.AsStructured()
		if derr == nil &&
			(verr == nil || verr.Error() == "target message part does not exist") {
			docs = append(docs, d)
			vars = append(vars, v)
		} else {
			if derr != nil {
				p.logger.Errorf("error getting structured doc: %s", derr)
			}
			if verr != nil {
				p.logger.Errorf("error getting structured vars: %s", verr)
			}
		}
	}

	// for holding old or new docs to be returned
	newDocs := make([]any, len(dvs))
	oldDocs := make([]any, len(dvs))

	var returnDocs []any

	switch op {
	case OperationCreate:
		createCtx := p.getCreateContext(ctx, newDocs, oldDocs)
		meta, errs, err = coll.CreateDocuments(createCtx, docs)
		for i, doc := range meta {
			resp := &createResp{
				Key:    doc.Key,
				ID:     doc.ID,
				Rev:    doc.Rev,
				OldRev: doc.OldRev,
			}
			if newDocs[i] != nil {
				resp.NEW = newDocs[i]
			}
			returnDocs = append(returnDocs, resp)
		}
	case OperationQuery:
		for i := range docs {
			var (
				varsMap map[string]interface{}
				ok bool
			)
			if vars[i] != nil { // queries can have no bindVars
				varsMap, ok = vars[i].(map[string]interface{})
				if !ok {
					p.logger.Errorf("query - error casting bindVars")
				}
			}
			if cursor, err = p.db.Query(ctx, p.config.Query.Aql, varsMap); err != nil {
				return nil, err
			}
			defer cursor.Close()
			queryDocs := make([]any, 0)

			for cursor.HasMore() {
				var doc interface{}
				_, err := cursor.ReadDocument(ctx, &doc)
				if err != nil {
					p.logger.Errorf("query - error reading cursor: %s", err)
				} else {
					queryDocs = append(queryDocs, doc)
				}
			}
			returnDocs = append(returnDocs, queryDocs)
		}
	default:
		err = fmt.Errorf("operation not implemented: %s", p.config.Operation)
		return nil, err
	}

	if err != nil {
		e := err.(arango.ArangoError)
		err = errors.New(e.ErrorMessage)
		return nil, err
	}

	var batch service.MessageBatch
	for _, doc := range returnDocs {
		b, err := json.Marshal(doc)
		if err != nil {
			p.logger.Errorf("error marshalling doc %v", doc)
		}
		msg := service.NewMessage(b)
		batch = append(batch, msg)
	}

	for i, e := range errs {
		if e != nil {
			p.logger.Errorf("arango error - batch index: %d - error: %s - doc: %v", i, errs[i].Error(), dvs[i])
			batch[i].SetError(errs[i])
		}
	}

	p.logger.Trace(fmt.Sprintf("operation successful: %s on %s", op, collName))
	return batch, nil
}

// return an arango Collection instance based on collection field or collection_map query result
func (p *arangoDBProcessor) getArangoCollectionName(ctx context.Context, i int, batch service.MessageBatch) (string, error) {
	var (
		m   *service.Message
		b   []byte
		err error
	)

	if p.config.Client.Collection != "" {
		return p.config.Client.Collection, nil
	}

	if p.config.Client.CollectionMap == nil {
		// both are nil
		if p.config.Operation == OperationQuery {
			// that's okay for queries
			return "", nil
		}
		return "", fmt.Errorf("'collection' and 'collection_map' cannot both be empty for operation '%s'", p.config.Operation)
	}

	if m, err = batch.BloblangQuery(i, p.config.Client.CollectionMap); err != nil {
		p.logger.Debug("collection_map failed")
		return "", err
	}
	if b, err = m.AsBytes(); err != nil {
		p.logger.Debug("error getting collection_map bytes")
		return "", err
	}
	return string(b), nil
}

func (p *arangoDBProcessor) getArangoCollection(ctx context.Context, name string) (arango.Collection, error) {
	coll, err := p.db.Collection(ctx, name)
	if err != nil {
		p.logger.Debug("error getting collection")
		return nil, err
	}

	return coll, nil
}

func (p *arangoDBProcessor) getCreateContext(ctx context.Context, newDocs, oldDocs []any) context.Context {
	newCtx := ctx
	var (
		v   any
		err error
	)
	for option, value := range p.config.Params {
		if v, err = value.FieldAny(); err != nil {
			// TODO: log error on Params value?
			continue
		}
		switch option {
		case "overwrite":
			if v == true {
				newCtx = arango.WithOverwrite(newCtx)
			}
		case "return_new":
			if v == true {
				newCtx = arango.WithReturnNew(newCtx, newDocs)
			}
		case "return_old":
			if v == true {
				newCtx = arango.WithReturnOld(newCtx, oldDocs)
			}
		default:
			p.logger.Errorf("unrecognized parameter: %s", option)
		}
	}
	return newCtx
}

func getQueryConfig(conf *service.ParsedConfig) (*Query, error) {
	var (
		Aql     string
		VarsMap *bloblang.Executor
		err     error
	)

	if conf.Contains("query", "aql") {
		if Aql, err = conf.FieldString("query", "aql"); err != nil {
			return nil, err
		}
	}

	if conf.Contains("query", "vars_map") {
		if VarsMap, err = conf.FieldBloblang("query", "vars_map"); err != nil {
			return nil, err
		}
	}

	return &Query{
		Aql,
		VarsMap,
	}, nil
}

func getProcessorConfig(conf *service.ParsedConfig) (*Config, error) {

	c := &Config{}

	var err error
	if c.Client, err = getClientConfig(conf); err != nil {
		return nil, err
	}

	if c.Operation, err = conf.FieldString(configKeyOperation); err != nil {
		return nil, err
	}

	if c.Params, err = conf.FieldAnyMap(configKeyParams); err != nil {
		return nil, err
	}

	if c.Query, err = getQueryConfig(conf); err != nil {
		return nil, err
	}

	if conf.Contains(configKeyDocumentMap) {
		if c.DocumentMap, err = conf.FieldBloblang(configKeyDocumentMap); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func getClientConfig(conf *service.ParsedConfig) (*client.Config, error) {

	var (
		c   = &client.Config{}
		err error
	)

	if c.Url, err = conf.FieldString(configKeyUrl); err != nil {
		return nil, err
	}

	if c.Username, err = conf.FieldString(configKeyUsername); err != nil {
		return nil, err
	}

	if c.Password, err = conf.FieldString(configKeyPassword); err != nil {
		return nil, err
	}

	if c.Database, err = conf.FieldString(configKeyDatabase); err != nil {
		return nil, err
	}

	if conf.Contains(configKeyCollection) {
		c.Collection, _ = conf.FieldString(configKeyCollection)
	}

	if conf.Contains(configKeyCollectionMap) {
		c.CollectionMap, _ = conf.FieldBloblang(configKeyCollectionMap)
	}

	return c, nil
}
