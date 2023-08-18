package arangodb

import (
	"fmt"

	arango "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// Client returns a new arangodb client based on the configuration parameters.
func (a Config) Client() (arango.Client, error) {

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{a.Url},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create arangodb connection: %v", err)
	}

	clientConfig := arango.ClientConfig{
		Connection:     conn,
		Authentication: arango.BasicAuthentication(a.Username, a.Password),
	}

	client, err := arango.NewClient(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create arangodb client: %v", err)
	}

	return client, nil

}
