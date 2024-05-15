package service

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/users"
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
			svc := NewUserService(ts, newTestIDGenerator(), WithListener(tl))

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
			svc := NewUserService(ts, newTestIDGenerator(), WithListener(tl))

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
			svc := NewUserService(ts, newTestIDGenerator(), WithListener(tl))

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

func TestUserService_GetUser(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id         string
		storeError bool
		wantErr    error
	}{
		"invalid-id": {
			id:      "",
			wantErr: errorz.ValidationError{},
		},
		"not-found": {
			id:      "missing",
			wantErr: errorz.NotFoundError{},
		},
		"store-error": {
			id:         validUserID,
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			id: validUserID,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			svc := NewUserService(ts, newTestIDGenerator(), WithListener(newTestListener(false)))

			user, err := svc.GetUser(context.Background(), test.id)

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

func TestUserService_GetUsersByIDs(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		ids        []string
		storeError bool
		want       []users.User
		wantErr    error
	}{
		"empty-ids": {
			ids:  []string{},
			want: []users.User{},
		},
		"invalid-ids": {
			ids:     []string{""},
			wantErr: errorz.ValidationError{},
		},
		"store-error": {
			ids:        []string{validUserID},
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			ids:  []string{validUserID},
			want: []users.User{validUser},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			svc := NewUserService(ts, newTestIDGenerator(), WithListener(newTestListener(false)))

			found, err := svc.GetUsersByIDs(context.Background(), test.ids)

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

func TestUserService_GetUserByUsername(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		username   string
		storeError bool
		wantErr    error
	}{
		"invalid-username": {
			username: "",
			wantErr:  errorz.ValidationError{},
		},
		"store-error": {
			username:   "user1",
			storeError: true,
			wantErr:    errorz.StoreError{},
		},
		"success": {
			username: "user1",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ts := newTestStore(t, test.storeError)
			svc := NewUserService(ts, newTestIDGenerator(), WithListener(newTestListener(false)))

			user, err := svc.GetUserByUsername(context.Background(), test.username)

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

func TestUserService_fireUserEvent_noListener(t *testing.T) {
	t.Parallel()

	svc := NewUserService(newTestStore(t, false), newTestIDGenerator())

	require.NotPanics(t, func() {
		_ = svc.fireUserEvent(
			context.Background(),
			adminActor,
			users.UserCreated,
			users.User{},
		)
	})
}
