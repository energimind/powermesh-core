package service

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/models"
	"github.com/stretchr/testify/require"
)

func TestModelService_CreateModel(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		data          models.ModelData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-modelData": {
			actor:   adminActor,
			data:    models.ModelData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			data:       validModelData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			data:          validModelData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			data:      validModelData,
			wantEvent: models.ModelCreated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestModelStore(t, test.storeError)
			tl := newTestModelListener(test.listenerError)

			svc := NewModelService(newTestIDGenerator(), ts, tl)

			model, err := svc.CreateModel(context.Background(), test.actor, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, model)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, model)
				require.NotEmpty(t, model.ID)
			}

			if test.wantEvent != "" {
				requireModelEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestModelService_UpdateModel(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		id            string
		data          models.ModelData
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-id": {
			actor:   adminActor,
			id:      "",
			data:    validModelData,
			wantErr: errorz.ValidationError{},
		},
		"invalid-modelData": {
			actor:   adminActor,
			id:      validModelID,
			data:    models.ModelData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			id:         validModelID,
			data:       validModelData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			id:            validModelID,
			data:          validModelData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			id:        validModelID,
			data:      validModelData,
			wantEvent: models.ModelUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestModelStore(t, test.storeError)
			tl := newTestModelListener(test.listenerError)

			svc := NewModelService(newTestIDGenerator(), ts, tl)

			model, err := svc.UpdateModel(context.Background(), test.actor, test.id, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, model)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, model)
				require.NotEmpty(t, model.ID)
			}

			if test.wantEvent != "" {
				requireModelEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestModelService_DeleteModel(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		id            string
		storeError    bool
		listenerError bool
		wantEvent     models.EventType
		wantErr       error
	}{
		"invalid-id": {
			actor:   adminActor,
			id:      "",
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			id:         validModelID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			id:            validModelID,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			id:        validModelID,
			wantEvent: models.ModelDeleted,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestModelStore(t, test.storeError)
			tl := newTestModelListener(test.listenerError)

			svc := NewModelService(newTestIDGenerator(), ts, tl)

			err := svc.DeleteModel(context.Background(), test.actor, test.id)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
			} else {
				require.NoError(t, err)
			}

			if test.wantEvent != "" {
				requireModelEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestModelService_GetModel(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id       string
		storeErr bool
		wantErr  error
	}{
		"invalid-id": {
			id:       "",
			storeErr: false,
			wantErr:  errorz.ValidationError{},
		},
		"not-found": {
			id:      missingModelID,
			wantErr: errorz.NotFoundError{},
		},
		"store-error": {
			id:       validModelID,
			storeErr: true,
			wantErr:  errorz.StoreError{},
		},
		"success": {
			id: validModelID,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestModelStore(t, test.storeErr)

			svc := NewModelService(newTestIDGenerator(), ts, newTestModelListener(false))

			model, err := svc.GetModel(context.Background(), test.id)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, model)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, model)
				require.NotEmpty(t, model.ID)
			}
		})
	}
}

func TestModelService_GetModelsByIDs(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		ids      []string
		storeErr bool
		want     []models.Model
		wantErr  error
	}{
		"empty-ids": {
			ids:  []string{},
			want: []models.Model{},
		},
		"invalid-ids": {
			ids:      []string{""},
			storeErr: false,
			wantErr:  errorz.ValidationError{},
		},
		"store-error": {
			ids:      []string{validModelID},
			storeErr: true,
			wantErr:  errorz.StoreError{},
		},
		"success": {
			ids:  []string{validModelID},
			want: []models.Model{validModel},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestModelStore(t, test.storeErr)

			svc := NewModelService(newTestIDGenerator(), ts, newTestModelListener(false))

			found, err := svc.GetModelsByIDs(context.Background(), test.ids)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, found)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.want, found)
			}
		})
	}
}
