package mock

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/arangodb/go-driver"
)

type MockCollection struct{}

func (c *MockCollection) Name() string                                                { return "MockCollection" }
func (c *MockCollection) Database() driver.Database                                   { return &MockDatabase{} }
func (c *MockCollection) Status(ctx context.Context) (driver.CollectionStatus, error) { return 0, nil }
func (c *MockCollection) Count(ctx context.Context) (int64, error)                    { return 0, nil }
func (c *MockCollection) Statistics(ctx context.Context) (driver.CollectionStatistics, error) {
	return driver.CollectionStatistics{}, nil
}
func (c *MockCollection) Checksum(ctx context.Context, withRevisions bool, withData bool) (driver.CollectionChecksum, error) {
	return driver.CollectionChecksum{}, nil
}
func (c *MockCollection) Revision(ctx context.Context) (string, error) { return "", nil }
func (c *MockCollection) Properties(ctx context.Context) (driver.CollectionProperties, error) {
	return driver.CollectionProperties{}, nil
}
func (c *MockCollection) SetProperties(ctx context.Context, options driver.SetCollectionPropertiesOptions) error {
	return nil
}
func (c *MockCollection) Shards(ctx context.Context, details bool) (driver.CollectionShards, error) {
	return driver.CollectionShards{}, nil
}
func (c *MockCollection) Load(ctx context.Context) error     { return nil }
func (c *MockCollection) Unload(ctx context.Context) error   { return nil }
func (c *MockCollection) Remove(ctx context.Context) error   { return nil }
func (c *MockCollection) Truncate(ctx context.Context) error { return nil }

// Index functions
func (c *MockCollection) Index(ctx context.Context, name string) (driver.Index, error) {
	return &MockIndex{}, nil
}
func (c *MockCollection) IndexExists(ctx context.Context, name string) (bool, error) {
	return true, nil
}
func (c *MockCollection) Indexes(ctx context.Context) ([]driver.Index, error) {
	return []driver.Index{}, nil
}
func (c *MockCollection) EnsureFullTextIndex(ctx context.Context, fields []string, options *driver.EnsureFullTextIndexOptions) (driver.Index, bool, error) {
	return &MockIndex{}, true, nil
}
func (c *MockCollection) EnsureGeoIndex(ctx context.Context, fields []string, options *driver.EnsureGeoIndexOptions) (driver.Index, bool, error) {
	return &MockIndex{}, true, nil
}
func (c *MockCollection) EnsureHashIndex(ctx context.Context, fields []string, options *driver.EnsureHashIndexOptions) (driver.Index, bool, error) {
	return &MockIndex{}, true, nil
}
func (c *MockCollection) EnsurePersistentIndex(ctx context.Context, fields []string, options *driver.EnsurePersistentIndexOptions) (driver.Index, bool, error) {
	return &MockIndex{}, true, nil
}
func (c *MockCollection) EnsureSkipListIndex(ctx context.Context, fields []string, options *driver.EnsureSkipListIndexOptions) (driver.Index, bool, error) {
	return &MockIndex{}, true, nil
}
func (c *MockCollection) EnsureTTLIndex(ctx context.Context, field string, expireAfter int, options *driver.EnsureTTLIndexOptions) (driver.Index, bool, error) {
	return &MockIndex{}, true, nil
}
func (c *MockCollection) EnsureZKDIndex(ctx context.Context, fields []string, options *driver.EnsureZKDIndexOptions) (driver.Index, bool, error) {
	return &MockIndex{}, true, nil
}
func (c *MockCollection) EnsureInvertedIndex(ctx context.Context, options *driver.InvertedIndexOptions) (driver.Index, bool, error) {
	return &MockIndex{}, true, nil
}

// Collection documents
func (c *MockCollection) DocumentExists(ctx context.Context, key string) (bool, error) {
	return true, nil
}
func (c *MockCollection) ReadDocument(ctx context.Context, key string, result interface{}) (driver.DocumentMeta, error) {
	return driver.DocumentMeta{}, nil
}
func (c *MockCollection) ReadDocuments(ctx context.Context, keys []string, results interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	return driver.DocumentMetaSlice{}, driver.ErrorSlice{}, nil
}
func (c *MockCollection) CreateDocument(ctx context.Context, document interface{}) (driver.DocumentMeta, error) {
	// create deterministic key and id
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%v", document)))
	key := fmt.Sprintf("mock-key-%.7s", hex.EncodeToString(h.Sum(nil)))
	return driver.DocumentMeta{
		Key:    key,
		ID:     driver.NewDocumentID(c.Name(), key),
		Rev:    "1",
		OldRev: "",
	}, nil
}
func (c *MockCollection) CreateDocuments(ctx context.Context, documents interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	const (
		keyReturnNew driver.ContextKey = "arangodb-returnNew"
	)

	// documents must be array
	docArry := documents.([]interface{})
	docMeta := make([]driver.DocumentMeta, len(docArry))

	// handle WithReturnNew (by just returning supplied doc)
	var newDocs []interface{}
	if val := ctx.Value(keyReturnNew); val != nil {
		newDocs = val.([]interface{})
	}

	for i, d := range docArry {
		m, _ := c.CreateDocument(ctx, d)
		docMeta[i] = m
		if newDocs != nil {
			newDocs[i] = d
		}
	}
	return docMeta, driver.ErrorSlice{}, nil
}
func (c *MockCollection) UpdateDocument(ctx context.Context, key string, update interface{}) (driver.DocumentMeta, error) {
	return driver.DocumentMeta{}, nil
}
func (c *MockCollection) UpdateDocuments(ctx context.Context, keys []string, updates interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	return driver.DocumentMetaSlice{}, driver.ErrorSlice{}, nil
}
func (c *MockCollection) ReplaceDocument(ctx context.Context, key string, document interface{}) (driver.DocumentMeta, error) {
	return driver.DocumentMeta{}, nil
}
func (c *MockCollection) ReplaceDocuments(ctx context.Context, keys []string, documents interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	return driver.DocumentMetaSlice{}, driver.ErrorSlice{}, nil
}
func (c *MockCollection) RemoveDocument(ctx context.Context, key string) (driver.DocumentMeta, error) {
	return driver.DocumentMeta{}, nil
}
func (c *MockCollection) RemoveDocuments(ctx context.Context, keys []string) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	return driver.DocumentMetaSlice{}, driver.ErrorSlice{}, nil
}
func (c *MockCollection) ImportDocuments(ctx context.Context, documents interface{}, options *driver.ImportDocumentOptions) (driver.ImportDocumentStatistics, error) {
	return driver.ImportDocumentStatistics{}, nil
}
