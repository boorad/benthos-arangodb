package mock

import (
	"context"
	"time"

	"github.com/arangodb/go-driver"
)

type MockBatch struct{}
type MockReplication struct{}

func (b *MockBatch) BatchID() string                                     { return "MockBatch" }
func (b *MockBatch) LastTick() driver.Tick                               { return "" }
func (b *MockBatch) Extend(ctx context.Context, ttl time.Duration) error { return nil }
func (b *MockBatch) Delete(ctx context.Context) error                    { return nil }

func (r *MockReplication) CreateBatch(ctx context.Context, db driver.Database, serverID int64, ttl time.Duration) (driver.Batch, error) {
	return &MockBatch{}, nil
}
func (r *MockReplication) DatabaseInventory(ctx context.Context, db driver.Database) (driver.DatabaseInventory, error) {
	return driver.DatabaseInventory{}, nil
}
func (r *MockReplication) GetRevisionTree(ctx context.Context, db driver.Database, batchId, collection string) (driver.RevisionTree, error) {
	return driver.RevisionTree{}, nil
}
func (r *MockReplication) GetRevisionsByRanges(ctx context.Context, db driver.Database, batchId, collection string, minMaxRevision []driver.RevisionMinMax, resume driver.RevisionUInt64) (driver.RevisionRanges, error) {
	return driver.RevisionRanges{}, nil
}
func (r *MockReplication) GetRevisionDocuments(ctx context.Context, db driver.Database, batchId, collection string, revisions driver.Revisions) ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}
