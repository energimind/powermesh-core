package service

import (
	"context"
	"errors"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/models"
	"github.com/stretchr/testify/require"
)

var (
	validNodeID     = "1" // must match generated ID from testIDGenerator
	validRelationID = "1" // must match generated ID from testIDGenerator
	validMeshData   = models.MeshData{
		Code: "code1",
	}
	validMesh = models.Mesh{
		ModelID: validModelID,
		Code:    validMeshData.Code,
	}
	validNodeData = models.NodeData{
		Kind: "kind1",
		Code: "code1",
		Props: models.PropBag{
			"section1": models.PropSection{
				"prop1": "value1",
			},
		},
	}
	validNode = models.Node{
		ID:    validNodeID,
		Kind:  validNodeData.Kind,
		Code:  validNodeData.Code,
		Props: validNodeData.Props,
	}
	validRelationData = models.RelationData{
		Kind: "kind1",
		From: "node1",
		To:   "node2",
		Props: models.PropBag{
			"section2": models.PropSection{
				"prop2": "value2",
			},
		},
	}
	validRelation = models.Relation{
		ID:    validRelationID,
		Kind:  validRelationData.Kind,
		From:  validRelationData.From,
		To:    validRelationData.To,
		Props: validRelationData.Props,
	}
)

type testMeshListener struct {
	forcedError error
	eventFired  models.MeshEvent
}

// Ensure that the testMeshListener implements the meshListener interface.
var _ meshListener = (*testMeshListener)(nil)

func newTestMeshListener(forcedError bool) *testMeshListener {
	var err error

	if forcedError {
		err = errors.New("forced-error")
	}

	return &testMeshListener{
		forcedError: err,
	}
}

func (l *testMeshListener) HandleMeshEvent(_ context.Context, event models.MeshEvent) error {
	if l.forcedError != nil {
		return l.forcedError
	}

	l.eventFired = event

	return nil
}

type testMeshStore struct {
	t           *testing.T
	forcedError error
}

// Ensure that the testMeshStore implements the meshStore interface.
var _ meshStore = (*testMeshStore)(nil)

func newTestMeshStore(t *testing.T, forcedError bool) *testMeshStore {
	var err error

	if forcedError {
		err = errorz.NewStoreError("forced-error")
	}

	return &testMeshStore{
		t:           t,
		forcedError: err,
	}
}

func (s *testMeshStore) CreateMesh(
	_ context.Context,
	mesh models.Mesh,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.Equal(s.t, validMesh, mesh)

	return nil
}

func (s *testMeshStore) UpdateMesh(
	_ context.Context,
	mesh models.Mesh,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.Equal(s.t, validMesh, mesh)

	return nil
}

func (s *testMeshStore) MergeMesh(
	_ context.Context,
	mesh models.Mesh,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.Equal(s.t, validMesh, mesh)

	return nil
}

func (s *testMeshStore) DeleteMesh(
	_ context.Context,
	modelID string,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, modelID)

	return nil
}

func (s *testMeshStore) GetMesh(
	_ context.Context,
	modelID string,
) (models.Mesh, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return models.Mesh{}, s.forcedError
	}

	require.NotEmpty(s.t, modelID)

	if modelID == validModelID {
		return models.Mesh{ModelID: modelID}, nil
	}

	return models.Mesh{}, errorz.NewNotFoundError("mesh %v not found", modelID)
}

func (s *testMeshStore) CreateNode(
	_ context.Context,
	modelID string,
	model models.Node,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, modelID)
	require.Equal(s.t, validNode, model)

	return nil
}

func (s *testMeshStore) UpdateNode(
	_ context.Context,
	modelID string,
	model models.Node,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, modelID)
	require.Equal(s.t, validNode, model)

	return nil
}

func (s *testMeshStore) DeleteNode(
	_ context.Context,
	modelID, nodeID string,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, modelID)
	require.NotEmpty(s.t, nodeID)

	return nil
}

func (s *testMeshStore) GetNode(
	_ context.Context,
	modelID, nodeID string,
) (models.Node, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return models.Node{}, s.forcedError
	}

	require.NotEmpty(s.t, modelID)
	require.NotEmpty(s.t, nodeID)

	if nodeID == validNodeID {
		return models.Node{ID: nodeID}, nil
	}

	return models.Node{}, errorz.NewNotFoundError("node %v not found", nodeID)
}

func (s *testMeshStore) CreateRelation(
	_ context.Context,
	modelID string,
	relation models.Relation,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, modelID)
	require.Equal(s.t, validRelation, relation)

	return nil
}

func (s *testMeshStore) UpdateRelation(
	_ context.Context,
	modelID string,
	relation models.Relation,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, modelID)
	require.Equal(s.t, validRelation, relation)

	return nil
}

func (s *testMeshStore) DeleteRelation(
	_ context.Context,
	modelID, relationID string,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, modelID)
	require.NotEmpty(s.t, relationID)

	return nil
}

func (s *testMeshStore) GetRelation(
	_ context.Context,
	modelID, relationID string,
) (models.Relation, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return models.Relation{}, s.forcedError
	}

	require.NotEmpty(s.t, modelID)
	require.NotEmpty(s.t, relationID)

	if relationID == validRelationID {
		return models.Relation{ID: relationID}, nil
	}

	return models.Relation{}, errorz.NewNotFoundError("relation %v not found", relationID)
}
