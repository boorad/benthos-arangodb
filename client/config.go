package arangodb

import "github.com/benthosdev/benthos/v4/public/bloblang"

// Config is a config struct for an arango connection.
type Config struct {
	Url           string             `json:"url" yaml:"url"`
	Username      string             `json:"username" yaml:"username"`
	Password      string             `json:"password" yaml:"password"`
	Database      string             `json:"database" yaml:"database"`
	Collection    string             `json:"collection" yaml:"collection"`
	CollectionMap *bloblang.Executor `json:"collection_map" yaml:"collection_map"`
}
