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
	return q.UpdateOne(s.meshes, toStoreMesh).
		Key(meshKey).
		Exec(ctx, mesh.ModelID, mesh)
}

// MergeMesh implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) MergeMesh(ctx context.Context, mesh models.Mesh) error {
	return q.MergeFields(s.meshes).
		Key(meshKey).
		Exec(ctx, mesh.ModelID, mergeMeshUpdate(mesh))
}

// DeleteMesh implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) DeleteMesh(ctx context.Context, modelID string) error {
	return q.DeleteOne(s.meshes).
		Key(meshKey).
		Exec(ctx, modelID)
}

// GetMesh implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) GetMesh(ctx context.Context, modelID string) (models.Mesh, error) {
	return q.GetOne(s.meshes, fromStoreMesh).
		Key(meshKey).
		Exec(ctx, modelID)
}

// CreateNode implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) CreateNode(ctx context.Context, modelID string, node models.Node) error {
	return q.EmbeddedPush(s.meshes, fieldNodes, toStoreNode).
		Key(meshKey).
		Exec(ctx, modelID, node)
}

// UpdateNode implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) UpdateNode(ctx context.Context, modelID string, node models.Node) error {
	return q.EmbeddedUpdate(s.meshes, fieldNodes, fieldID, toStoreNode).
		Key(meshKey).
		Exec(ctx, modelID, node.ID, node)
}

// DeleteNode implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) DeleteNode(ctx context.Context, modelID, nodeID string) error {
	return q.EmbeddedPull(s.meshes, fieldNodes, fieldID).
		Key(meshKey).
		Exec(ctx, modelID, nodeID)
}

// GetNode implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) GetNode(ctx context.Context, modelID, nodeID string) (models.Node, error) {
	return q.EmbeddedGetOne(s.meshes, fieldNodes, fieldID, extractFirstNode).
		Key(meshKey).
		Exec(ctx, modelID, nodeID)
}

// CreateRelation implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) CreateRelation(ctx context.Context, modelID string, relation models.Relation) error {
	return q.EmbeddedPush(s.meshes, fieldRelations, toStoreRelation).
		Key(meshKey).
		Exec(ctx, modelID, relation)
}

// UpdateRelation implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) UpdateRelation(ctx context.Context, modelID string, relation models.Relation) error {
	return q.EmbeddedUpdate(s.meshes, fieldRelations, fieldID, toStoreRelation).
		Key(meshKey).
		Exec(ctx, modelID, relation.ID, relation)
}

// DeleteRelation implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) DeleteRelation(ctx context.Context, modelID, relationID string) error {
	return q.EmbeddedPull(s.meshes, fieldRelations, fieldID).
		Key(meshKey).
		Exec(ctx, modelID, relationID)
}

// GetRelation implements the mesh store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshStore) GetRelation(ctx context.Context, modelID, relationID string) (models.Relation, error) {
	return q.EmbeddedGetOne(s.meshes, fieldRelations, fieldID, extractFirstRelation).
		Key(meshKey).
		Exec(ctx, modelID, relationID)
}
