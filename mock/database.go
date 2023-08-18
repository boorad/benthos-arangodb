package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockDatabase struct{}

func (db *MockDatabase) Name() string {
	return "MockDB"
}

func (db *MockDatabase) Info(ctx context.Context) (driver.DatabaseInfo, error) {
	return driver.DatabaseInfo{}, nil
}

func (db *MockDatabase) EngineInfo(ctx context.Context) (driver.EngineInfo, error) {
	return driver.EngineInfo{}, nil
}

func (db *MockDatabase) Remove(ctx context.Context) error { return nil }

// DatabaseCollections
func (db *MockDatabase) Collection(ctx context.Context, name string) (driver.Collection, error) {
	return &MockCollection{}, nil
}
func (db *MockDatabase) CollectionExists(ctx context.Context, name string) (bool, error) {
	return true, nil
}
func (db *MockDatabase) Collections(ctx context.Context) ([]driver.Collection, error) {
	return []driver.Collection{}, nil
}
func (db *MockDatabase) CreateCollection(ctx context.Context, name string, options *driver.CreateCollectionOptions) (driver.Collection, error) {
	return &MockCollection{}, nil
}

// View functions
func (db *MockDatabase) View(ctx context.Context, name string) (driver.View, error) {
	return &MockView{}, nil
}
func (db *MockDatabase) ViewExists(ctx context.Context, name string) (bool, error) {
	return true, nil
}
func (db *MockDatabase) Views(ctx context.Context) ([]driver.View, error) {
	return []driver.View{}, nil
}
func (db *MockDatabase) CreateArangoSearchView(ctx context.Context, name string, options *driver.ArangoSearchViewProperties) (driver.ArangoSearchView, error) {
	return &MockArangoSearchView{}, nil
}
func (db *MockDatabase) CreateArangoSearchAliasView(ctx context.Context, name string, options *driver.ArangoSearchAliasViewProperties) (driver.ArangoSearchViewAlias, error) {
	return &MockArangoSearchViewAlias{}, nil
}

// Graph functions
func (db *MockDatabase) Graph(ctx context.Context, name string) (driver.Graph, error) {
	return &MockGraph{}, nil
}
func (db *MockDatabase) GraphExists(ctx context.Context, name string) (bool, error) {
	return true, nil
}
func (db *MockDatabase) Graphs(ctx context.Context) ([]driver.Graph, error) {
	return []driver.Graph{}, nil
}
func (db *MockDatabase) CreateGraph(ctx context.Context, name string, options *driver.CreateGraphOptions) (driver.Graph, error) {
	return &MockGraph{}, nil
}
func (db *MockDatabase) CreateGraphV2(ctx context.Context, name string, options *driver.CreateGraphOptions) (driver.Graph, error) {
	return &MockGraph{}, nil
}

// Pregel functions
func (db *MockDatabase) StartJob(ctx context.Context, options driver.PregelJobOptions) (string, error) {
	return "", nil
}
func (db *MockDatabase) GetJob(ctx context.Context, id string) (*driver.PregelJob, error) {
	return &driver.PregelJob{}, nil
}
func (db *MockDatabase) GetJobs(ctx context.Context) ([]*driver.PregelJob, error) {
	return []*driver.PregelJob{}, nil
}
func (db *MockDatabase) CancelJob(ctx context.Context, id string) error {
	return nil
}

// Query functions
func (db *MockDatabase) ExplainQuery(ctx context.Context, query string, bindVars map[string]interface{}, opts *driver.ExplainQueryOptions) (driver.ExplainQueryResult, error) {
	return driver.ExplainQueryResult{}, nil
}

// Streaming trasnaction functions
func (db *MockDatabase) BeginTransaction(ctx context.Context, cols driver.TransactionCollections, opts *driver.BeginTransactionOptions) (driver.TransactionID, error) {
	return "", nil
}
func (db *MockDatabase) CommitTransaction(ctx context.Context, tid driver.TransactionID, opts *driver.CommitTransactionOptions) error {
	return nil
}
func (db *MockDatabase) AbortTransaction(ctx context.Context, tid driver.TransactionID, opts *driver.AbortTransactionOptions) error {
	return nil
}

func (db *MockDatabase) TransactionStatus(ctx context.Context, tid driver.TransactionID) (driver.TransactionStatusRecord, error) {
	return driver.TransactionStatusRecord{}, nil
}

// ArangoSearch Analyers API
func (db *MockDatabase) EnsureAnalyzer(ctx context.Context, analyzer driver.ArangoSearchAnalyzerDefinition) (bool, driver.ArangoSearchAnalyzer, error) {
	return true, &MockArangoSearchAnalyzer{}, nil
}
func (db *MockDatabase) Analyzer(ctx context.Context, name string) (driver.ArangoSearchAnalyzer, error) {
	return &MockArangoSearchAnalyzer{}, nil
}
func (db *MockDatabase) Analyzers(ctx context.Context) ([]driver.ArangoSearchAnalyzer, error) {
	return []driver.ArangoSearchAnalyzer{}, nil
}

func (db *MockDatabase) Query(ctx context.Context, query string, bindVars map[string]interface{}) (driver.Cursor, error) {
	return &MockCursor{}, nil
}
func (db *MockDatabase) ValidateQuery(ctx context.Context, query string) error {
	return nil
}
func (db *MockDatabase) OptimizerRulesForQueries(ctx context.Context) ([]driver.QueryRule, error) {
	return []driver.QueryRule{}, nil
}
func (db *MockDatabase) Transaction(ctx context.Context, action string, options *driver.TransactionOptions) (interface{}, error) {
	return "", nil
}
