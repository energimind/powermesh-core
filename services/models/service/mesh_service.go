package service

import (
	"context"
	"time"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/models"
)

// meshStore defines the interface for a mesh store.
type meshStore interface {
	meshOperations
	nodeOperations
	relationOperations
}

// meshOperations defines the operations on meshes.
type meshOperations interface {
	CreateMesh(ctx context.Context, mesh models.Mesh) error
	UpdateMesh(ctx context.Context, mesh models.Mesh) error
	DeleteMesh(ctx context.Context, modelID string) error
	GetMesh(ctx context.Context, modelID string) (models.Mesh, error)
}

// nodeOperations defines the operations on nodes.
type nodeOperations interface {
	CreateNode(ctx context.Context, modelID string, node models.Node) error
	UpdateNode(ctx context.Context, modelID string, node models.Node) error
	DeleteNode(ctx context.Context, modelID, nodeID string) error
	GetNode(ctx context.Context, modelID, nodeID string) (models.Node, error)
}

// relationOperations defines the operations on relations.
type relationOperations interface {
	CreateRelation(ctx context.Context, modelID string, relation models.Relation) error
	UpdateRelation(ctx context.Context, modelID string, relation models.Relation) error
	DeleteRelation(ctx context.Context, modelID, relationID string) error
	GetRelation(ctx context.Context, modelID, relationID string) (models.Relation, error)
}

// meshListener defines the external mesh event modelListener.
type meshListener interface {
	HandleMeshEvent(ctx context.Context, event models.MeshEvent) error
}

// MeshService implements the mesh service.
//
// It implements the models.MeshService interface.
//
// We do not wrap the errors returned by the store because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type MeshService struct {
	idGen    idGenerator
	store    meshStore
	listener meshListener
	now      func() time.Time
}

// NewMeshService creates a new mesh service.
func NewMeshService(idGen idGenerator, store meshStore, listener meshListener) *MeshService {
	return &MeshService{
		idGen:    idGen,
		store:    store,
		listener: listener,
		now:      time.Now,
	}
}

// CreateMesh implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) CreateMesh(
	ctx context.Context,
	actor access.Actor,
	modelID string,
	data models.MeshData,
) (models.Mesh, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Mesh{}, err
	}

	mesh := meshFromData(modelID, data)

	if err := s.store.CreateMesh(ctx, mesh); err != nil {
		return models.Mesh{}, err
	}

	if err := s.fireMeshEvent(ctx, actor, models.MeshCreated, mesh); err != nil {
		return models.Mesh{}, err
	}

	return mesh, nil
}

// UpdateMesh implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) UpdateMesh(
	ctx context.Context,
	actor access.Actor,
	modelID string,
	data models.MeshData,
) (models.Mesh, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Mesh{}, err
	}

	mesh := meshFromData(modelID, data)

	if err := s.store.UpdateMesh(ctx, mesh); err != nil {
		return models.Mesh{}, err
	}

	if err := s.fireMeshEvent(ctx, actor, models.MeshUpdated, mesh); err != nil {
		return models.Mesh{}, err
	}

	return mesh, nil
}

// DeleteMesh implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) DeleteMesh(
	ctx context.Context,
	actor access.Actor,
	modelID string,
) error {
	if err := validateModelID(modelID); err != nil {
		return err
	}

	if err := s.store.DeleteMesh(ctx, modelID); err != nil {
		return err
	}

	if err := s.fireMeshEvent(ctx, actor, models.MeshDeleted, models.Mesh{ModelID: modelID}); err != nil {
		return err
	}

	return nil
}

// GetMesh implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) GetMesh(
	ctx context.Context,
	modelID string,
) (models.Mesh, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Mesh{}, err
	}

	mesh, err := s.store.GetMesh(ctx, modelID)
	if err != nil {
		return models.Mesh{}, err
	}

	return mesh, nil
}

// CreateNode implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) CreateNode(
	ctx context.Context,
	actor access.Actor,
	modelID string,
	data models.NodeData,
) (models.Node, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Node{}, err
	}

	if err := validateNodeData(data); err != nil {
		return models.Node{}, err
	}

	node := nodeFromData(s.idGen.GenerateID(), data)

	if err := s.store.CreateNode(ctx, modelID, node); err != nil {
		return models.Node{}, err
	}

	updates := models.Mesh{
		ModelID: modelID,
		Nodes:   map[string]models.Node{node.ID: node},
	}

	if err := s.fireMeshContentsEvent(ctx, actor, models.MeshContentsCreated, updates, models.Mesh{}); err != nil {
		return models.Node{}, err
	}

	return node, nil
}

// UpdateNode implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) UpdateNode(
	ctx context.Context,
	actor access.Actor,
	modelID, nodeID string,
	data models.NodeData,
) (models.Node, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Node{}, err
	}

	if err := validateNodeID(nodeID); err != nil {
		return models.Node{}, err
	}

	if err := validateNodeData(data); err != nil {
		return models.Node{}, err
	}

	node := nodeFromData(nodeID, data)

	if err := s.store.UpdateNode(ctx, modelID, node); err != nil {
		return models.Node{}, err
	}

	updates := models.Mesh{
		ModelID: modelID,
		Nodes:   map[string]models.Node{node.ID: node},
	}

	if err := s.fireMeshContentsEvent(ctx, actor, models.MeshContentsUpdated, updates, models.Mesh{}); err != nil {
		return models.Node{}, err
	}

	return node, nil
}

// DeleteNode implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) DeleteNode(
	ctx context.Context,
	actor access.Actor,
	modelID, nodeID string,
) error {
	if err := validateModelID(modelID); err != nil {
		return err
	}

	if err := validateNodeID(nodeID); err != nil {
		return err
	}

	if err := s.store.DeleteNode(ctx, modelID, nodeID); err != nil {
		return err
	}

	deletes := models.Mesh{
		ModelID: modelID,
		Nodes:   map[string]models.Node{nodeID: {ID: nodeID}},
	}

	if err := s.fireMeshContentsEvent(ctx, actor, models.MeshContentsDeleted, models.Mesh{}, deletes); err != nil {
		return err
	}

	return nil
}

// GetNode implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) GetNode(
	ctx context.Context,
	modelID, nodeID string,
) (models.Node, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Node{}, err
	}

	if err := validateNodeID(nodeID); err != nil {
		return models.Node{}, err
	}

	node, err := s.store.GetNode(ctx, modelID, nodeID)
	if err != nil {
		return models.Node{}, err
	}

	return node, nil
}

// CreateRelation implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) CreateRelation(
	ctx context.Context,
	actor access.Actor,
	modelID string,
	data models.RelationData,
) (models.Relation, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Relation{}, err
	}

	if err := validateRelationData(data); err != nil {
		return models.Relation{}, err
	}

	relation := relationFromData(s.idGen.GenerateID(), data)

	if err := s.store.CreateRelation(ctx, modelID, relation); err != nil {
		return models.Relation{}, err
	}

	updates := models.Mesh{
		ModelID:   modelID,
		Relations: map[string]models.Relation{relation.ID: relation},
	}

	if err := s.fireMeshContentsEvent(ctx, actor, models.MeshContentsCreated, updates, models.Mesh{}); err != nil {
		return models.Relation{}, err
	}

	return relation, nil
}

// UpdateRelation implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) UpdateRelation(
	ctx context.Context,
	actor access.Actor,
	modelID, relationID string,
	data models.RelationData,
) (models.Relation, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Relation{}, err
	}

	if err := validateRelationID(relationID); err != nil {
		return models.Relation{}, err
	}

	if err := validateRelationData(data); err != nil {
		return models.Relation{}, err
	}

	relation := relationFromData(relationID, data)

	if err := s.store.UpdateRelation(ctx, modelID, relation); err != nil {
		return models.Relation{}, err
	}

	updates := models.Mesh{
		ModelID:   modelID,
		Relations: map[string]models.Relation{relation.ID: relation},
	}

	if err := s.fireMeshContentsEvent(ctx, actor, models.MeshContentsUpdated, updates, models.Mesh{}); err != nil {
		return models.Relation{}, err
	}

	return relation, nil
}

// DeleteRelation implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) DeleteRelation(
	ctx context.Context,
	actor access.Actor,
	modelID, relationID string,
) error {
	if err := validateModelID(modelID); err != nil {
		return err
	}

	if err := validateRelationID(relationID); err != nil {
		return err
	}

	if err := s.store.DeleteRelation(ctx, modelID, relationID); err != nil {
		return err
	}

	deletes := models.Mesh{
		ModelID:   modelID,
		Relations: map[string]models.Relation{relationID: {ID: relationID}},
	}

	if err := s.fireMeshContentsEvent(ctx, actor, models.MeshContentsDeleted, models.Mesh{}, deletes); err != nil {
		return err
	}

	return nil
}

// GetRelation implements the models.MeshService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *MeshService) GetRelation(
	ctx context.Context,
	modelID, relationID string,
) (models.Relation, error) {
	if err := validateModelID(modelID); err != nil {
		return models.Relation{}, err
	}

	if err := validateRelationID(relationID); err != nil {
		return models.Relation{}, err
	}

	relation, err := s.store.GetRelation(ctx, modelID, relationID)
	if err != nil {
		return models.Relation{}, err
	}

	return relation, nil
}

// fireMeshEvent fires a mesh event.
func (s *MeshService) fireMeshEvent(
	ctx context.Context,
	actor access.Actor,
	eventType models.EventType,
	mesh models.Mesh,
) error {
	event := models.MeshEvent{
		EventHeader: models.EventHeader{
			Type:      eventType,
			Actor:     actor,
			Timestamp: s.now(),
		},
		Updates: mesh,
	}

	if err := s.listener.HandleMeshEvent(ctx, event); err != nil {
		return errorz.NewInternalError("%s event handler failed: %v", eventType, err)
	}

	return nil
}

// fireMeshContentsEvent fires a mesh contents event.
func (s *MeshService) fireMeshContentsEvent(
	ctx context.Context,
	actor access.Actor,
	eventType models.EventType,
	updates, deletes models.Mesh,
) error {
	event := models.MeshEvent{
		EventHeader: models.EventHeader{
			Type:      eventType,
			Actor:     actor,
			Timestamp: s.now(),
		},
		Updates: updates,
		Deletes: deletes,
	}

	if err := s.listener.HandleMeshEvent(ctx, event); err != nil {
		return errorz.NewInternalError("%s event handler failed: %v", eventType, err)
	}

	return nil
}
