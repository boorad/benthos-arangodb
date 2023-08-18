package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockView struct{}
type MockArangoSearchView struct{}
type MockArangoSearchViewAlias struct{}

func (v *MockView) Name() string          { return "MockView" }
func (v *MockView) Type() driver.ViewType { return "" }
func (v *MockView) ArangoSearchView() (driver.ArangoSearchView, error) {
	return &MockArangoSearchView{}, nil
}
func (v *MockView) ArangoSearchViewAlias() (driver.ArangoSearchViewAlias, error) {
	return &MockArangoSearchViewAlias{}, nil
}
func (v *MockView) Database() driver.Database                        { return &MockDatabase{} }
func (v *MockView) Rename(ctx context.Context, newName string) error { return nil }
func (v *MockView) Remove(ctx context.Context) error                 { return nil }

// ArangoSearchView
func (v *MockArangoSearchView) Name() string          { return "MockView" }
func (v *MockArangoSearchView) Type() driver.ViewType { return "" }
func (v *MockArangoSearchView) ArangoSearchView() (driver.ArangoSearchView, error) {
	return &MockArangoSearchView{}, nil
}
func (v *MockArangoSearchView) ArangoSearchViewAlias() (driver.ArangoSearchViewAlias, error) {
	return &MockArangoSearchViewAlias{}, nil
}
func (v *MockArangoSearchView) Database() driver.Database                        { return &MockDatabase{} }
func (v *MockArangoSearchView) Rename(ctx context.Context, newName string) error { return nil }
func (v *MockArangoSearchView) Remove(ctx context.Context) error                 { return nil }
func (v *MockArangoSearchView) Properties(ctx context.Context) (driver.ArangoSearchViewProperties, error) {
	return driver.ArangoSearchViewProperties{}, nil
}
func (v *MockArangoSearchView) SetProperties(ctx context.Context, options driver.ArangoSearchViewProperties) error {
	return nil
}

// ArangoSearchAliasView
func (v *MockArangoSearchViewAlias) Name() string          { return "MockView" }
func (v *MockArangoSearchViewAlias) Type() driver.ViewType { return "" }
func (v *MockArangoSearchViewAlias) ArangoSearchView() (driver.ArangoSearchView, error) {
	return &MockArangoSearchView{}, nil
}
func (v *MockArangoSearchViewAlias) ArangoSearchViewAlias() (driver.ArangoSearchViewAlias, error) {
	return &MockArangoSearchViewAlias{}, nil
}
func (v *MockArangoSearchViewAlias) Database() driver.Database                        { return &MockDatabase{} }
func (v *MockArangoSearchViewAlias) Rename(ctx context.Context, newName string) error { return nil }
func (v *MockArangoSearchViewAlias) Remove(ctx context.Context) error                 { return nil }
func (v *MockArangoSearchViewAlias) Properties(ctx context.Context) (driver.ArangoSearchAliasViewProperties, error) {
	return driver.ArangoSearchAliasViewProperties{}, nil
}
func (v *MockArangoSearchViewAlias) SetProperties(ctx context.Context, options driver.ArangoSearchAliasViewProperties) (driver.ArangoSearchAliasViewProperties, error) {
	return driver.ArangoSearchAliasViewProperties{}, nil
}
