package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockConnection struct{}
type MockRequest struct{}
type MockResponse struct{}

// Connection
func (c *MockConnection) NewRequest(method, path string) (driver.Request, error) {
	return &MockRequest{}, nil
}
func (c *MockConnection) Do(ctx context.Context, req driver.Request) (driver.Response, error) {
	return &MockResponse{}, nil
}
func (c *MockConnection) Unmarshal(data driver.RawObject, result interface{}) error { return nil }
func (c *MockConnection) Endpoints() []string                                       { return []string{} }
func (c *MockConnection) UpdateEndpoints(endpoints []string) error                  { return nil }
func (c *MockConnection) SetAuthentication(driver.Authentication) (driver.Connection, error) {
	return &MockConnection{}, nil
}
func (c *MockConnection) Protocols() driver.ProtocolSet { return driver.ProtocolSet{} }

// Request
func (r *MockRequest) SetQuery(key, value string) driver.Request { return &MockRequest{} }
func (r *MockRequest) SetBody(body ...interface{}) (driver.Request, error) {
	return &MockRequest{}, nil
}
func (r *MockRequest) SetBodyArray(bodyArray interface{}, mergeArray []map[string]interface{}) (driver.Request, error) {
	return &MockRequest{}, nil
}
func (r *MockRequest) SetBodyImportArray(bodyArray interface{}) (driver.Request, error) {
	return &MockRequest{}, nil
}
func (r *MockRequest) SetHeader(key, value string) driver.Request { return &MockRequest{} }
func (r *MockRequest) Written() bool                              { return true }
func (r *MockRequest) Clone() driver.Request                      { return &MockRequest{} }
func (r *MockRequest) Path() string                               { return "" }
func (r *MockRequest) Method() string                             { return "" }

// response
func (r *MockResponse) StatusCode() int                                  { return 200 }
func (r *MockResponse) Endpoint() string                                 { return "" }
func (r *MockResponse) CheckStatus(validStatusCodes ...int) error        { return nil }
func (r *MockResponse) Header(key string) string                         { return "" }
func (r *MockResponse) ParseBody(field string, result interface{}) error { return nil }
func (r *MockResponse) ParseArrayBody() ([]driver.Response, error)       { return []driver.Response{}, nil }
