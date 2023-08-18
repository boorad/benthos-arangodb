package arangodb

import (
	client "github.com/boorad/benthos-arangodb/client"
	"github.com/benthosdev/benthos/v4/public/bloblang"
	"github.com/benthosdev/benthos/v4/public/service"
)

type Config struct {
	Client      *client.Config
	Operation   string                           `json:"operation" yaml:"operation"`
	Params      map[string]*service.ParsedConfig `json:"params" yaml:"params"`
	FilterMap   *bloblang.Executor               `json:"filter_map" yaml:"filter_map"`
	DocumentMap *bloblang.Executor               `json:"document_map" yaml:"document_map"`
	Query       *Query                           `json:"query" yaml:"query"`
}

type Query struct {
	Aql     string
	VarsMap *bloblang.Executor
}

const (
	OperationCreate string = "create"
	OperationRead   string = "read"
	OperationUpdate string = "update"
	OperationDelete string = "delete"
	OperationQuery  string = "query"
)

var Operations = []string{
	OperationCreate,
	OperationRead,
	OperationUpdate,
	OperationDelete,
	OperationQuery,
}

func isDocumentAllowed(op string) bool {
	switch op {
	case OperationCreate, OperationUpdate:
		return true
	default:
		return false
	}
}
