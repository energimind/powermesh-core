package service

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/permissions"
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
		"access-denied": {
			actor:   guestActor,
			data:    validRoleBindingData,
			wantErr: errorz.AccessDeniedError{},
		},
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
			wantErr:       errorz.GatewayError{},
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

			svc := NewPermissionService(newTestIDGenerator(), ts, tl)

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
		"access-denied": {
			actor:   guestActor,
			data:    validRoleBindingData,
			wantErr: errorz.AccessDeniedError{},
		},
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
			wantErr:       errorz.GatewayError{},
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

			svc := NewPermissionService(newTestIDGenerator(), ts, tl)

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
		"access-denied": {
			actor:   guestActor,
			id:      validRoleBindingID,
			wantErr: errorz.AccessDeniedError{},
		},
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
			wantErr:       errorz.GatewayError{},
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

			svc := NewPermissionService(newTestIDGenerator(), ts, tl)

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

func TestPermissionService_GetRoleBinding(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		query         permissions.RoleBindingQuery
		storeError    bool
		listenerError bool
		wantErr       error
	}{
		"access-denied": {
			actor:   guestActor,
			query:   validRoleBindingQuery,
			wantErr: errorz.AccessDeniedError{},
		},
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

			svc := NewPermissionService(newTestIDGenerator(), ts, tl)

			rb, err := svc.GetRoleBinding(context.Background(), test.actor, test.query)

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

func TestPermissionService_GetAccessibleObjects(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		query         permissions.AccessibleObjectsQuery
		storeError    bool
		listenerError bool
		wantErr       error
	}{
		"access-denied": {
			actor:   guestActor,
			query:   validAccessibleObjectsQuery,
			wantErr: errorz.AccessDeniedError{},
		},
		"invalid-query": {
			actor:   adminActor,
			query:   permissions.AccessibleObjectsQuery{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			query:      validAccessibleObjectsQuery,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			actor: adminActor,
			query: validAccessibleObjectsQuery,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)

			svc := NewPermissionService(newTestIDGenerator(), ts, tl)

			objects, err := svc.GetAccessibleObjects(context.Background(), test.actor, test.query)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Nil(t, objects)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, objects)
			}
		})
	}
}
