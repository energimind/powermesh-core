package service

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/permissions"
	"github.com/stretchr/testify/require"
)

func TestPermissionService_CreateRoleBinding(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		data          permissions.RoleBindingData
		storeError    bool
		listenerError bool
		wantEvent     permissions.EventType
		wantErr       error
	}{
		"invalid-roleBindingData": {
			actor:   adminActor,
			data:    permissions.RoleBindingData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			data:       validRoleBindingData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			data:          validRoleBindingData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			data:      validRoleBindingData,
			wantEvent: permissions.RoleBindingCreated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)

			svc := NewPermissionService(newTestIDGenerator(), ts, WithListener(tl))

			rb, err := svc.CreateRoleBinding(context.Background(), test.actor, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, rb)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, rb)
				require.NotEmpty(t, rb.ID)
			}

			if test.wantEvent != "" {
				requireEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestPermissionService_UpdateRoleBinding(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		id            string
		data          permissions.RoleBindingData
		storeError    bool
		listenerError bool
		wantEvent     permissions.EventType
		wantErr       error
	}{
		"invalid-id": {
			actor:   adminActor,
			id:      "",
			data:    validRoleBindingData,
			wantErr: errorz.ValidationError{},
		},
		"invalid-roleBindingData": {
			actor:   adminActor,
			id:      validRoleBindingID,
			data:    permissions.RoleBindingData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			id:         validRoleBindingID,
			data:       validRoleBindingData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			id:            validRoleBindingID,
			data:          validRoleBindingData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			id:        validRoleBindingID,
			data:      validRoleBindingData,
			wantEvent: permissions.RoleBindingUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)

			svc := NewPermissionService(newTestIDGenerator(), ts, WithListener(tl))

			rb, err := svc.UpdateRoleBinding(context.Background(), test.actor, test.id, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, rb)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, rb)
				require.NotEmpty(t, rb.ID)
			}

			if test.wantEvent != "" {
				requireEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestPermissionService_DeleteRoleBinding(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		id            string
		storeError    bool
		listenerError bool
		wantEvent     permissions.EventType
		wantErr       error
	}{
		"invalid-id": {
			actor:   adminActor,
			id:      "",
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			id:         validRoleBindingID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			id:            validRoleBindingID,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			id:        validRoleBindingID,
			wantEvent: permissions.RoleBindingDeleted,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)

			svc := NewPermissionService(newTestIDGenerator(), ts, WithListener(tl))

			err := svc.DeleteRoleBinding(context.Background(), test.actor, test.id)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
			} else {
				require.NoError(t, err)
			}

			if test.wantEvent != "" {
				requireEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestPermissionService_DeleteRoleBindingsByResource(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		query         permissions.RoleBindingQuery
		storeError    bool
		listenerError bool
		wantEvent     permissions.EventType
		wantErr       error
	}{
		"invalid-query": {
			actor:   adminActor,
			query:   permissions.RoleBindingQuery{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			query:      validRoleBindingQuery,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			query:         validRoleBindingQuery,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			query:     validRoleBindingQuery,
			wantEvent: permissions.RoleBindingDeleted,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)

			svc := NewPermissionService(newTestIDGenerator(), ts, WithListener(tl))

			err := svc.DeleteRoleBindingsByResource(context.Background(), test.actor, test.query.ResourceID, test.query.ResourceType)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
			} else {
				require.NoError(t, err)
			}

			if test.wantEvent != "" {
				requireEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestPermissionService_GetRoleBinding(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		query         permissions.RoleBindingQuery
		storeError    bool
		listenerError bool
		wantErr       error
	}{
		"invalid-query": {
			actor:   adminActor,
			query:   permissions.RoleBindingQuery{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			query:      validRoleBindingQuery,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			actor: adminActor,
			query: validRoleBindingQuery,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)

			svc := NewPermissionService(newTestIDGenerator(), ts, WithListener(tl))

			rb, err := svc.GetRoleBinding(context.Background(), test.query)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, rb)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, rb)
				require.NotEmpty(t, rb.ID)
			}
		})
	}
}

func TestPermissionService_GetRoleBindingsByOwner(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		ownerID       string
		storeError    bool
		listenerError bool
		wantErr       error
	}{
		"invalid-ownerID": {
			actor:   adminActor,
			ownerID: "",
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			ownerID:    validOwnerID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			actor:   adminActor,
			ownerID: validOwnerID,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)

			svc := NewPermissionService(newTestIDGenerator(), ts, WithListener(tl))

			rbs, err := svc.GetRoleBindingsByOwner(context.Background(), test.ownerID)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, rbs)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, rbs)
			}
		})
	}
}

func TestPermissionService_GetAccessibleResources(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		query         permissions.AccessibleResourcesQuery
		storeError    bool
		listenerError bool
		wantErr       error
	}{
		"invalid-query": {
			actor:   adminActor,
			query:   permissions.AccessibleResourcesQuery{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			query:      validAccessibleResourcesQuery,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			actor: adminActor,
			query: validAccessibleResourcesQuery,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)

			svc := NewPermissionService(newTestIDGenerator(), ts, WithListener(tl))

			resources, err := svc.GetAccessibleResources(context.Background(), test.query)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Nil(t, resources)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, resources)
			}
		})
	}
}
