package mock

import (
	"context"
	"time"

	"github.com/arangodb/go-driver"
)

type MockCursor struct{}
type MockQueryStatistics struct{}
type MockQueryExtra struct{}

func (c *MockCursor) Close() error  { return nil }
func (c *MockCursor) HasMore() bool { return false }
func (c *MockCursor) ReadDocument(ctx context.Context, result interface{}) (driver.DocumentMeta, error) {
	return driver.DocumentMeta{}, nil
}
func (qe *MockCursor) RetryReadDocument(ctx context.Context, result interface{}) (driver.DocumentMeta, error) {
	return driver.DocumentMeta{}, nil
}
func (c *MockCursor) Count() int64                       { return 0 }
func (c *MockCursor) Statistics() driver.QueryStatistics { return &MockQueryStatistics{} }
func (c *MockCursor) Extra() driver.QueryExtra           { return &MockQueryExtra{} }

func (qs *MockQueryStatistics) WritesExecuted() int64        { return 0 }
func (qs *MockQueryStatistics) WritesIgnored() int64         { return 0 }
func (qs *MockQueryStatistics) ScannedFull() int64           { return 0 }
func (qs *MockQueryStatistics) ScannedIndex() int64          { return 0 }
func (qs *MockQueryStatistics) Filtered() int64              { return 0 }
func (qs *MockQueryStatistics) FullCount() int64             { return 0 }
func (qs *MockQueryStatistics) ExecutionTime() time.Duration { return 0 }

func (qe *MockQueryExtra) GetStatistics() driver.QueryStatistics { return &MockQueryStatistics{} }
func (qe *MockQueryExtra) GetProfileRaw() ([]byte, bool, error)  { return []byte{}, true, nil }
func (qe *MockQueryExtra) GetPlanRaw() ([]byte, bool, error)     { return []byte{}, true, nil }
