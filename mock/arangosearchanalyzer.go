package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockArangoSearchAnalyzer struct{}

func (sa *MockArangoSearchAnalyzer) Name() string                          { return "MockArangoSearchAnalyzer" }
func (sa *MockArangoSearchAnalyzer) Type() driver.ArangoSearchAnalyzerType { return "" }
func (sa *MockArangoSearchAnalyzer) UniqueName() string                    { return "MockArangoSearchAnalyzer" }
func (sa *MockArangoSearchAnalyzer) Definition() driver.ArangoSearchAnalyzerDefinition {
	return driver.ArangoSearchAnalyzerDefinition{}
}
func (sa *MockArangoSearchAnalyzer) Properties() driver.ArangoSearchAnalyzerProperties {
	return driver.ArangoSearchAnalyzerProperties{}
}
func (sa *MockArangoSearchAnalyzer) Database() driver.Database                    { return &MockDatabase{} }
func (sa *MockArangoSearchAnalyzer) Remove(ctx context.Context, force bool) error { return nil }
