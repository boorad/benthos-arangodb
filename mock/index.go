package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockIndex struct{}

func (i *MockIndex) Name() string                     { return "MockIndex" }
func (i *MockIndex) ID() string                       { return "" }
func (i *MockIndex) UserName() string                 { return "" }
func (i *MockIndex) Type() driver.IndexType           { return "" }
func (i *MockIndex) Remove(ctx context.Context) error { return nil }
func (i *MockIndex) Fields() []string                 { return []string{} }
func (i *MockIndex) Unique() bool                     { return true }
func (i *MockIndex) Deduplicate() bool                { return false }
func (i *MockIndex) Sparse() bool                     { return false }
func (i *MockIndex) GeoJSON() bool                    { return false }
func (i *MockIndex) InBackground() bool               { return false }
func (i *MockIndex) Estimates() bool                  { return false }
func (i *MockIndex) MinLength() int                   { return 0 }
func (i *MockIndex) ExpireAfter() int                 { return 0 }
func (i *MockIndex) LegacyPolygons() bool             { return false }
func (i *MockIndex) CacheEnabled() bool               { return false }
func (i *MockIndex) StoredValues() []string           { return []string{} }
func (i *MockIndex) InvertedIndexOptions() driver.InvertedIndexOptions {
	return driver.InvertedIndexOptions{}
}
