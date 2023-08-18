package arangodb

import (
	"context"
	"testing"
	"time"

	"github.com/boorad/benthos-arangodb/mock"
	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestArangoProcessorCreateWithDocumentMap(t *testing.T) {

	inputMsg := service.NewMessage([]byte(`{"key_a": "value_a", "key_b": "value_b"}`))

	tests := []struct {
		name string
		conf string
		want map[string]interface{}
	}{
		{
			name: "no document_map",
			conf: `
url: url
username: username
password: password
database: database
collection: test
operation: create
params:
  return_new: true`,
			want: map[string]interface{}{"key_a": "value_a", "key_b": "value_b"},
		},
		{
			name: "document_map remove key",
			conf: `
url: url
username: username
password: password
database: database
operation: create
collection: test
params:
  return_new: true
document_map: |
  root = this
  root.key_b = deleted()`,
			want: map[string]interface{}{"key_a": "value_a"},
		},
		{
			name: "document_map new key",
			conf: `
url: url
username: username
password: password
database: database
operation: create
collection: test
params:
  return_new: true
document_map: |
  root = this
  root.key_c = "value_c"`,
			want: map[string]interface{}{"key_a": "value_a", "key_b": "value_b", "key_c": "value_c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))
			defer cancel()

			spec := arangoProcessorConfigSpec()
			env := service.NewEnvironment()
			resources := service.MockResources()

			parsedConfig, err := spec.ParseYAML(tt.conf, env)
			require.NoError(t, err)

			config, err := getProcessorConfig(parsedConfig)
			require.NoError(t, err)

			processor, err := newArangoDBProcessor(config, &mock.MockClient{}, &mock.MockDatabase{}, resources.Logger(), resources.Metrics())
			require.NoError(t, err)

			batch := service.MessageBatch{inputMsg}
			out, err := processor.ProcessBatch(ctx, batch)
			require.NoError(t, err)

			// tests have single message in single batch
			m, err := out[0][0].AsStructured()
			require.NoError(t, err)

			// extract the NEW doc and compare
			n := m.(map[string]interface{})
			if diff := cmp.Diff(tt.want, n["NEW"]); diff != "" {
				t.Errorf("expected result not equal:\n%v", diff)
			}
		})
	}
}
