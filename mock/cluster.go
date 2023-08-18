package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockCluster struct{}

func (c *MockCluster) Health(ctx context.Context) (driver.ClusterHealth, error) {
	return driver.ClusterHealth{}, nil
}
func (c *MockCluster) DatabaseInventory(ctx context.Context, db driver.Database) (driver.DatabaseInventory, error) {
	return driver.DatabaseInventory{}, nil
}
func (c *MockCluster) MoveShard(ctx context.Context, col driver.Collection, shard driver.ShardID, fromServer, toServer driver.ServerID) error {
	return nil
}
func (c *MockCluster) CleanOutServer(ctx context.Context, serverID string) error { return nil }
func (c *MockCluster) ResignServer(ctx context.Context, serverID string) error   { return nil }
func (c *MockCluster) IsCleanedOut(ctx context.Context, serverID string) (bool, error) {
	return true, nil
}
func (c *MockCluster) RemoveServer(ctx context.Context, serverID driver.ServerID) error { return nil }
