package mock

import (
	"context"

	"github.com/arangodb/go-driver"
)

type MockGraph struct{}

func (g *MockGraph) Name() string                     { return "MockGraph" }
func (g *MockGraph) Remove(ctx context.Context) error { return nil }
func (g *MockGraph) IsSmart() bool                    { return true }
func (g *MockGraph) IsSatellite() bool                { return false }
func (g *MockGraph) IsDisjoint() bool                 { return false }

// GraphEdgeCollections
func (g *MockGraph) EdgeCollection(ctx context.Context, name string) (driver.Collection, driver.VertexConstraints, error) {
	return &MockCollection{}, driver.VertexConstraints{}, nil
}
func (g *MockGraph) EdgeCollectionExists(ctx context.Context, name string) (bool, error) {
	return true, nil
}
func (g *MockGraph) EdgeCollections(ctx context.Context) ([]driver.Collection, []driver.VertexConstraints, error) {
	return []driver.Collection{}, []driver.VertexConstraints{}, nil
}
func (g *MockGraph) CreateEdgeCollection(ctx context.Context, collection string, constraints driver.VertexConstraints) (driver.Collection, error) {
	return &MockCollection{}, nil
}
func (g *MockGraph) CreateEdgeCollectionWithOptions(ctx context.Context, collection string, constraints driver.VertexConstraints, options driver.CreateEdgeCollectionOptions) (driver.Collection, error) {
	return &MockCollection{}, nil
}
func (g *MockGraph) SetVertexConstraints(ctx context.Context, collection string, constraints driver.VertexConstraints) error {
	return nil
}

// GraphVertexCollections
func (g *MockGraph) VertexCollection(ctx context.Context, name string) (driver.Collection, error) {
	return &MockCollection{}, nil
}
func (g *MockGraph) VertexCollectionExists(ctx context.Context, name string) (bool, error) {
	return true, nil
}
func (g *MockGraph) VertexCollections(ctx context.Context) ([]driver.Collection, error) {
	return []driver.Collection{}, nil
}
func (g *MockGraph) CreateVertexCollection(ctx context.Context, collection string) (driver.Collection, error) {
	return &MockCollection{}, nil
}
func (g *MockGraph) CreateVertexCollectionWithOptions(ctx context.Context, collection string, options driver.CreateVertexCollectionOptions) (driver.Collection, error) {
	return &MockCollection{}, nil
}

func (g *MockGraph) ID() string                               { return "" }
func (g *MockGraph) Key() driver.DocumentID                   { return "" }
func (g *MockGraph) Rev() string                              { return "" }
func (g *MockGraph) EdgeDefinitions() []driver.EdgeDefinition { return []driver.EdgeDefinition{} }
func (g *MockGraph) SmartGraphAttribute() string              { return "" }
func (g *MockGraph) MinReplicationFactor() int                { return 0 }
func (g *MockGraph) NumberOfShards() int                      { return 0 }
func (g *MockGraph) OrphanCollections() []string              { return []string{} }
func (g *MockGraph) ReplicationFactor() int                   { return 0 }
func (g *MockGraph) WriteConcern() int                        { return 0 }
