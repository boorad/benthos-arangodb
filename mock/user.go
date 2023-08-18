package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockUser struct{}

func (u *MockUser) Name() string                                                  { return "MockUser" }
func (u *MockUser) IsActive() bool                                                { return true }
func (u *MockUser) IsPasswordChangeNeeded() bool                                  { return false }
func (u *MockUser) Extra(result interface{}) error                                { return nil }
func (u *MockUser) Remove(ctx context.Context) error                              { return nil }
func (u *MockUser) Update(ctx context.Context, options driver.UserOptions) error  { return nil }
func (u *MockUser) Replace(ctx context.Context, options driver.UserOptions) error { return nil }
func (u *MockUser) AccessibleDatabases(ctx context.Context) ([]driver.Database, error) {
	return []driver.Database{}, nil
}
func (u *MockUser) SetDatabaseAccess(ctx context.Context, db driver.Database, access driver.Grant) error {
	return nil
}
func (u *MockUser) GetDatabaseAccess(ctx context.Context, db driver.Database) (driver.Grant, error) {
	return "", nil
}
func (u *MockUser) RemoveDatabaseAccess(ctx context.Context, db driver.Database) error { return nil }
func (u *MockUser) SetCollectionAccess(ctx context.Context, col driver.AccessTarget, access driver.Grant) error {
	return nil
}
func (u *MockUser) GetCollectionAccess(ctx context.Context, col driver.AccessTarget) (driver.Grant, error) {
	return "", nil
}
func (u *MockUser) RemoveCollectionAccess(ctx context.Context, col driver.AccessTarget) error {
	return nil
}
func (u *MockUser) GrantReadWriteAccess(ctx context.Context, db driver.Database) error { return nil }
func (u *MockUser) RevokeAccess(ctx context.Context, db driver.Database) error         { return nil }
