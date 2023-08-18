package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockClient struct{}
type MockFoxxService struct{}

func (c *MockClient) SynchronizeEndpoints(ctx context.Context) error                 { return nil }
func (c *MockClient) SynchronizeEndpoints2(ctx context.Context, dbname string) error { return nil }
func (c *MockClient) Connection() driver.Connection                                  { return &MockConnection{} }

// Database functions
func (c *MockClient) Database(ctx context.Context, name string) (driver.Database, error) {
	return &MockDatabase{}, nil
}
func (c *MockClient) DatabaseExists(ctx context.Context, name string) (bool, error) { return true, nil }
func (c *MockClient) Databases(ctx context.Context) ([]driver.Database, error) {
	return []driver.Database{}, nil
}
func (c *MockClient) AccessibleDatabases(ctx context.Context) ([]driver.Database, error) {
	return []driver.Database{}, nil
}
func (c *MockClient) CreateDatabase(ctx context.Context, name string, options *driver.CreateDatabaseOptions) (driver.Database, error) {
	return &MockDatabase{}, nil
}

// User functions
func (c *MockClient) User(ctx context.Context, name string) (driver.User, error) {
	return &MockUser{}, nil
}
func (c *MockClient) UserExists(ctx context.Context, name string) (bool, error) { return true, nil }
func (c *MockClient) Users(ctx context.Context) ([]driver.User, error)          { return []driver.User{}, nil }
func (c *MockClient) CreateUser(ctx context.Context, name string, options *driver.UserOptions) (driver.User, error) {
	return &MockUser{}, nil
}

// Cluster functions
func (c *MockClient) Cluster(ctx context.Context) (driver.Cluster, error) { return &MockCluster{}, nil }

// Individual server information functions
func (c *MockClient) Version(ctx context.Context) (driver.VersionInfo, error) {
	return driver.VersionInfo{}, nil
}
func (c *MockClient) ServerRole(ctx context.Context) (driver.ServerRole, error) { return "", nil }
func (c *MockClient) ServerID(ctx context.Context) (string, error)              { return "", nil }

// Server/cluster administration functions
func (c *MockClient) ServerMode(ctx context.Context) (driver.ServerMode, error)       { return "", nil }
func (c *MockClient) SetServerMode(ctx context.Context, mode driver.ServerMode) error { return nil }
func (c *MockClient) Shutdown(ctx context.Context, removeFromCluster bool) error      { return nil }
func (c *MockClient) Metrics(ctx context.Context) ([]byte, error)                     { return []byte{}, nil }
func (c *MockClient) MetricsForSingleServer(ctx context.Context, serverID string) ([]byte, error) {
	return []byte{}, nil
}
func (c *MockClient) Statistics(ctx context.Context) (driver.ServerStatistics, error) {
	return driver.ServerStatistics{}, nil
}
func (c *MockClient) ShutdownV2(ctx context.Context, removeFromCluster, graceful bool) error {
	return nil
}
func (c *MockClient) ShutdownInfoV2(ctx context.Context) (driver.ShutdownInfo, error) {
	return driver.ShutdownInfo{}, nil
}
func (c *MockClient) Logs(ctx context.Context) (driver.ServerLogs, error) {
	return driver.ServerLogs{}, nil
}
func (c *MockClient) GetLogLevels(ctx context.Context, opts *driver.LogLevelsGetOptions) (driver.LogLevels, error) {
	return driver.LogLevels{}, nil
}
func (c *MockClient) SetLogLevels(ctx context.Context, logLevels driver.LogLevels, opts *driver.LogLevelsSetOptions) error { return nil }

// Replication functions
func (c *MockClient) Replication() driver.Replication { return &MockReplication{} }

// Backup functions
func (c *MockClient) Backup() driver.ClientBackup { return &MockBackup{} }
func (c *MockClient) Foxx() driver.FoxxService    { return &MockFoxxService{} }

func (f *MockFoxxService) InstallFoxxService(ctx context.Context, zipFile string, options driver.FoxxCreateOptions) error {
	return nil
}
func (f *MockFoxxService) UninstallFoxxService(ctx context.Context, options driver.FoxxDeleteOptions) error {
	return nil
}
