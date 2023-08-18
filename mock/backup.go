package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockBackup struct{}

func (b *MockBackup) Create(ctx context.Context, opt *driver.BackupCreateOptions) (driver.BackupID, driver.BackupCreateResponse, error) {
	return "", driver.BackupCreateResponse{}, nil
}
func (b *MockBackup) Delete(ctx context.Context, id driver.BackupID) error { return nil }
func (b *MockBackup) Restore(ctx context.Context, id driver.BackupID, opt *driver.BackupRestoreOptions) error {
	return nil
}
func (b *MockBackup) List(ctx context.Context, opt *driver.BackupListOptions) (map[driver.BackupID]driver.BackupMeta, error) {
	return map[driver.BackupID]driver.BackupMeta{}, nil
}
func (b *MockBackup) Upload(ctx context.Context, id driver.BackupID, remoteRepository string, config interface{}) (driver.BackupTransferJobID, error) {
	return "", nil
}
func (b *MockBackup) Download(ctx context.Context, id driver.BackupID, remoteRepository string, config interface{}) (driver.BackupTransferJobID, error) {
	return "", nil
}
func (b *MockBackup) Progress(ctx context.Context, job driver.BackupTransferJobID) (driver.BackupTransferProgressReport, error) {
	return driver.BackupTransferProgressReport{}, nil
}
func (b *MockBackup) Abort(ctx context.Context, job driver.BackupTransferJobID) error { return nil }
