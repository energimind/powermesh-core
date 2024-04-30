package service

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/models"
	"github.com/stretchr/testify/require"
)

func TestMeshService_CreateMesh(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		data          models.MeshData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			data:    validMeshData,
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			data:       validMeshData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			data:          validMeshData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			modelID:   validModelID,
			data:      validMeshData,
			wantEvent: models.MeshCreated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			mesh, err := svc.CreateMesh(context.Background(), test.actor, test.modelID, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, mesh)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, mesh)
			}
		})
	}
}

func TestMeshService_UpdateMesh(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		data          models.MeshData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			data:    validMeshData,
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			data:       validMeshData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			data:          validMeshData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			modelID:   validModelID,
			data:      validMeshData,
			wantEvent: models.MeshUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			mesh, err := svc.UpdateMesh(context.Background(), test.actor, test.modelID, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, mesh)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, mesh)
			}
		})
	}
}

func TestMeshService_MergeMesh(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		data          models.MeshData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			data:    validMeshData,
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			data:       validMeshData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			data:          validMeshData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			modelID:   validModelID,
			data:      validMeshData,
			wantEvent: models.MeshUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			err := svc.MergeMesh(context.Background(), test.actor, test.modelID, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMeshService_DeleteMesh(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			modelID:   validModelID,
			wantEvent: models.MeshDeleted,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			err := svc.DeleteMesh(context.Background(), test.actor, test.modelID)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMeshService_GetMesh(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		modelID    string
		storeError bool
		wantErr    error
	}{
		"invalid-modelID": {
			modelID: "",
			wantErr: errorz.ValidationError{},
		},
		"not-found": {
			modelID: "missing",
			wantErr: errorz.NotFoundError{},
		},
		"store-error": {
			modelID:    validModelID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			modelID: validModelID,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)

			svc := NewMeshService(newTestIDGenerator(), ts, nil)

			mesh, err := svc.GetMesh(context.Background(), test.modelID)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, mesh)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, mesh)
			}
		})
	}
}

func TestMeshService_CreateNode(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		data          models.NodeData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			wantErr: errorz.ValidationError{},
		},
		"invalid-data": {
			actor:   adminActor,
			modelID: validModelID,
			data:    models.NodeData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			data:       validNodeData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			data:          validNodeData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			modelID:   validModelID,
			data:      validNodeData,
			wantEvent: models.MeshUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			node, err := svc.CreateNode(context.Background(), test.actor, test.modelID, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, node)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, node)
			}
		})
	}
}

func TestMeshService_UpdateNode(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		nodeID        string
		data          models.NodeData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			wantErr: errorz.ValidationError{},
		},
		"invalid-nodeID": {
			actor:   adminActor,
			modelID: validModelID,
			nodeID:  "",
			wantErr: errorz.ValidationError{},
		},
		"invalid-data": {
			actor:   adminActor,
			modelID: validModelID,
			nodeID:  validNodeID,
			data:    models.NodeData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			nodeID:     validNodeID,
			data:       validNodeData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			nodeID:        validNodeID,
			data:          validNodeData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			modelID:   validModelID,
			nodeID:    validNodeID,
			data:      validNodeData,
			wantEvent: models.MeshUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			node, err := svc.UpdateNode(context.Background(), test.actor, test.modelID, test.nodeID, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, node)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, node)
			}
		})
	}
}

func TestMeshService_DeleteNode(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		nodeID        string
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			wantErr: errorz.ValidationError{},
		},
		"invalid-nodeID": {
			actor:   adminActor,
			modelID: validModelID,
			nodeID:  "",
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			nodeID:     validNodeID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			nodeID:        validNodeID,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			modelID:   validModelID,
			nodeID:    validNodeID,
			wantEvent: models.MeshUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			err := svc.DeleteNode(context.Background(), test.actor, test.modelID, test.nodeID)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMeshService_GetNode(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		modelID    string
		nodeID     string
		storeError bool
		wantErr    error
	}{
		"invalid-modelID": {
			modelID: "",
			nodeID:  validNodeID,
			wantErr: errorz.ValidationError{},
		},
		"invalid-nodeID": {
			modelID: validModelID,
			nodeID:  "",
			wantErr: errorz.ValidationError{},
		},
		"not-found": {
			modelID: validModelID,
			nodeID:  "missing",
			wantErr: errorz.NotFoundError{},
		},
		"store-error": {
			modelID:    validModelID,
			nodeID:     validNodeID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			modelID: validModelID,
			nodeID:  validNodeID,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)

			svc := NewMeshService(newTestIDGenerator(), ts, nil)

			node, err := svc.GetNode(context.Background(), test.modelID, test.nodeID)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, node)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, node)
			}
		})
	}
}

func TestMeshService_CreateRelation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		data          models.RelationData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			wantErr: errorz.ValidationError{},
		},
		"invalid-data": {
			actor:   adminActor,
			modelID: validModelID,
			data:    models.RelationData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			data:       validRelationData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			data:          validRelationData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			modelID:   validModelID,
			data:      validRelationData,
			wantEvent: models.MeshUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			relation, err := svc.CreateRelation(context.Background(), test.actor, test.modelID, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, relation)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMeshService_UpdateRelation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		relationID    string
		data          models.RelationData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			wantErr: errorz.ValidationError{},
		},
		"invalid-relationID": {
			actor:      adminActor,
			modelID:    validModelID,
			relationID: "",
			wantErr:    errorz.ValidationError{},
		},
		"invalid-data": {
			actor:      adminActor,
			modelID:    validModelID,
			relationID: validRelationID,
			data:       models.RelationData{},
			wantErr:    errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			relationID: validRelationID,
			data:       validRelationData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			relationID:    validRelationID,
			data:          validRelationData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:      adminActor,
			modelID:    validModelID,
			relationID: validRelationID,
			data:       validRelationData,
			wantEvent:  models.MeshUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			relation, err := svc.UpdateRelation(context.Background(), test.actor, test.modelID, test.relationID, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, relation)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMeshService_DeleteRelation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		modelID       string
		relationID    string
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelID": {
			actor:   adminActor,
			modelID: "",
			wantErr: errorz.ValidationError{},
		},
		"invalid-relationID": {
			actor:      adminActor,
			modelID:    validModelID,
			relationID: "",
			wantErr:    errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			modelID:    validModelID,
			relationID: validRelationID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"modelListener-error": {
			actor:         adminActor,
			modelID:       validModelID,
			relationID:    validRelationID,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:      adminActor,
			modelID:    validModelID,
			relationID: validRelationID,
			wantEvent:  models.MeshUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)
			tl := newTestMeshListener(test.listenerError)

			svc := NewMeshService(newTestIDGenerator(), ts, tl)

			err := svc.DeleteRelation(context.Background(), test.actor, test.modelID, test.relationID)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMeshService_GetRelation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		modelID    string
		relationID string
		storeError bool
		wantErr    error
	}{
		"invalid-modelID": {
			modelID:    "",
			relationID: validRelationID,
			wantErr:    errorz.ValidationError{},
		},
		"invalid-relationID": {
			modelID:    validModelID,
			relationID: "",
			wantErr:    errorz.ValidationError{},
		},
		"not-found": {
			modelID:    validModelID,
			relationID: "missing",
			wantErr:    errorz.NotFoundError{},
		},
		"store-error": {
			modelID:    validModelID,
			relationID: validRelationID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			modelID:    validModelID,
			relationID: validRelationID,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestMeshStore(t, test.storeError)

			svc := NewMeshService(newTestIDGenerator(), ts, nil)

			relation, err := svc.GetRelation(context.Background(), test.modelID, test.relationID)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, relation)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, relation)
			}
		})
	}
}
