package service

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/users"
	"github.com/stretchr/testify/require"
)

func TestUserService_CreateUser(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		data          users.UserData
		storeError    bool
		listenerError bool
		wantEvent     users.EventType
		wantErr       error
	}{
		"access-denied": {
			actor:   guestActor,
			data:    validUserData,
			wantErr: errorz.AccessDeniedError{},
		},
		"invalid-userData": {
			actor:   adminActor,
			data:    users.UserData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			data:       validUserData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			data:          validUserData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			data:      validUserData,
			wantEvent: users.UserCreated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)
			svc := NewUserService(newTestIDGenerator(), ts, tl)

			user, err := svc.CreateUser(context.Background(), test.actor, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, user)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, user)
				require.NotEmpty(t, user.ID)
			}

			if test.wantEvent != "" {
				requireEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		id            string
		data          users.UserData
		storeError    bool
		listenerError bool
		wantEvent     users.EventType
		wantErr       error
	}{
		"access-denied": {
			actor:   guestActor,
			id:      validUserID,
			data:    validUserData,
			wantErr: errorz.AccessDeniedError{},
		},
		"invalid-id": {
			actor:   adminActor,
			id:      "",
			data:    validUserData,
			wantErr: errorz.ValidationError{},
		},
		"invalid-userData": {
			actor:   adminActor,
			id:      validUserID,
			data:    users.UserData{},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			id:         validUserID,
			data:       validUserData,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			id:            validUserID,
			data:          validUserData,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			id:        validUserID,
			data:      validUserData,
			wantEvent: users.UserUpdated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)
			svc := NewUserService(newTestIDGenerator(), ts, tl)

			user, err := svc.UpdateUser(context.Background(), test.actor, test.id, test.data)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, user)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, user)
				require.NotEmpty(t, user.ID)
			}

			if test.wantEvent != "" {
				requireEventFired(t, test.wantEvent, tl)
			} else {
				require.Empty(t, tl.eventFired)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor         access.Actor
		id            string
		storeError    bool
		listenerError bool
		wantEvent     users.EventType
		wantErr       error
	}{
		"access-denied": {
			actor:   guestActor,
			id:      validUserID,
			wantErr: errorz.AccessDeniedError{},
		},
		"invalid-id": {
			actor:   adminActor,
			id:      "",
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			id:         validUserID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"listener-error": {
			actor:         adminActor,
			id:            validUserID,
			listenerError: true,
			wantErr:       errorz.InternalError{},
		},
		"success": {
			actor:     adminActor,
			id:        validUserID,
			wantEvent: users.UserDeleted,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			tl := newTestListener(test.listenerError)
			svc := NewUserService(newTestIDGenerator(), ts, tl)

			err := svc.DeleteUser(context.Background(), test.actor, test.id)

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

func TestUserService_GetUserByUsername(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		actor      access.Actor
		username   string
		storeError bool
		wantErr    error
	}{
		"access-denied": {
			actor:    guestActor,
			username: "user1",
			wantErr:  errorz.AccessDeniedError{},
		},
		"invalid-username": {
			actor:    adminActor,
			username: "",
			wantErr:  errorz.ValidationError{},
		},
		"store-error": {
			actor:      adminActor,
			username:   "user1",
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			actor:    adminActor,
			username: "user1",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			svc := NewUserService(newTestIDGenerator(), ts, newTestListener(false))

			user, err := svc.GetUserByUsername(context.Background(), test.actor, test.username)

			if test.wantErr != nil {
				require.Error(t, err)
				require.IsType(t, test.wantErr, err)
				require.Empty(t, user)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, user)
				require.NotEmpty(t, user.ID)
			}
		})
	}
}
