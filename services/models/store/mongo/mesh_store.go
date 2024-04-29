package mongo

import (
	"context"

	q "github.com/energimind/powermesh-core/mongoquery"
	"github.com/energimind/powermesh-core/services/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collMeshes = "meshes"
	meshKey    = "modelId"
)

// MeshStore is a MongoDB store for meshes.
//
// We do not wrap the errors returned by mongoquery utilities because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type MeshStore struct {
	meshes *mongo.Collection
}

// NewMeshStore creates a new MongoDB mesh store.
func NewMeshStore(db *mongo.Database) *MeshStore {
	return &MeshStore{
		meshes: db.Collection(collMeshes),
	}
}

// CreateMesh implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) CreateMesh(ctx context.Context, mesh models.Mesh) error {
	return q.CreateOne(s.meshes, toStoreMesh).Exec(ctx, mesh)
}

// UpdateMesh implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) UpdateMesh(ctx context.Context, mesh models.Mesh) error {
	return q.UpdateOne(s.meshes, toStoreMesh).Key(meshKey).Exec(ctx, mesh.ModelID, mesh)
}

// DeleteMesh implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) DeleteMesh(ctx context.Context, modelID string) error {
	return q.DeleteOne(s.meshes).Key(meshKey).Exec(ctx, modelID)
}

// GetMesh implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) GetMesh(ctx context.Context, modelID string) (models.Mesh, error) {
	return q.GetOne(s.meshes, fromStoreMesh).Key(meshKey).Exec(ctx, modelID)
}
