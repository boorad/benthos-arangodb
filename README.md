# Benthos ArangoDB Plugin

Interact with ArangoDB in Benthos pipelines

## Build

Add this dependency to your project
```shell
go get github.com/boorad/benthos-arangodb
```

Author a main file that calls `service.Run()` and imports this plugin

```go
package main

import (
	"github.com/benthosdev/benthos/v4/public/service"

	// Import all standard Benthos components
	_ "github.com/benthosdev/benthos/v4/public/components/all"

	// Add this plugin package here
	_ "github.com/boorad/benthos-arangodb"
)

func main() {
	service.RunCLI(context.Background())
}
```

Finally, build your custom main function:

```shell
go build
```

Alternatively build it as a Docker image with:

```shell
go mod vendor
docker build . -t benthos-arangodb
```

## Configuration Examples

Create vertex documents in a collection named for the value in `this.objects`:

```yaml
arangodb:
  url: ${ARANGODB_URL}
  username: ${ARANGODB_USERNAME}
  password: ${ARANGODB_PASSWORD}
  database: ${ARANGODB_DATABASE}
  collection_map: |
    root = this.objects
  operation: create
  params:
    overwrite: true

```

Create edge documents in `edges` collection:

```yaml
mapping: |
  root = {
    "_key":  "123",
    "_from": "systems/benthos",
    "_to":   "systems/arangodb",
    "plugin_type":  "processor"
  }

arangodb:
  url: ${ARANGODB_URL}
  username: ${ARANGODB_USERNAME}
  password: ${ARANGODB_PASSWORD}
  database: ${ARANGODB_DATABASE}
  collection: edges
  operation: create
  params:
    overwrite: true
```

Query via AQL:

```yaml
arangodb:
  url: ${ARANGODB_URL}
  username: ${ARANGODB_USERNAME}
  password: ${ARANGODB_PASSWORD}
  database: ${ARANGODB_DATABASE}
  operation: query
  query:
    aql: |
      FOR s IN systems
        FILTER s.type IN @types
        RETURN s
    vars_map: |
      root.types = ["input", "processor", "output"]

```
