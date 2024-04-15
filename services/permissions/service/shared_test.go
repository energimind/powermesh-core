package service

import (
	"context"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/permissions"
	"github.com/stretchr/testify/require"
)

var (
	adminActor           = access.Actor{Role: access.RoleAdmin}
	guestActor           = access.Actor{Role: access.RoleGuest}
	validRoleBindingID   = "rb1"
	validRoleBindingData = permissions.RoleBindingData{
		UserID:     "user1",
		ObjectID:   "object1",
		ObjectType: permissions.ObjectTypeModel,
		Role:       access.RoleAdmin,
	}
	validRoleBindingQuery = permissions.RoleBindingQuery{
		UserID:   "user1",
		ObjectID: "object1",
	}
	validAccessibleObjectsQuery = permissions.AccessibleObjectsQuery{
		UserID:     "user1",
		ObjectType: permissions.ObjectTypeModel,
	}
)

type testIDGenerator struct {
	idCounter atomic.Int64
}

func newTestIDGenerator() *testIDGenerator {
	return &testIDGenerator{}
}

func (g *testIDGenerator) GenerateID() string {
	return strconv.FormatInt(g.idCounter.Add(1), 10)
}

type testListener struct {
	forcedError error
	eventFired  permissions.Event
}

func newTestListener(forcedError bool) *testListener {
	var err error

	if forcedError {
		err = errorz.NewGatewayError("forced-error")
	}

	return &testListener{forcedError: err}
}

func (l *testListener) HandlePermissionEvent(_ context.Context, event permissions.Event) error {
	if l.forcedError != nil {
		return l.forcedError
	}

	l.eventFired = event

	return nil
}

type testStore struct {
	t           *testing.T
	forcedError error
}

func newTestStore(t *testing.T, forcedError bool) *testStore {
	var err error

	if forcedError {
		err = errorz.NewStoreError("forced-error")
	}

	return &testStore{
		t:           t,
		forcedError: err,
	}
}

func (s *testStore) CreateRoleBinding(_ context.Context, id string, data permissions.RoleBindingData) (permissions.RoleBinding, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return permissions.RoleBinding{}, s.forcedError
	}

	require.NotEmpty(s.t, id)
	require.Equal(s.t, validRoleBindingData, data)

	return permissions.RoleBinding{ID: id}, nil
}

func (s *testStore) UpdateRoleBinding(_ context.Context, id string, data permissions.RoleBindingData) (permissions.RoleBinding, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return permissions.RoleBinding{}, s.forcedError
	}

	require.NotEmpty(s.t, id)
	require.Equal(s.t, validRoleBindingData, data)

	return permissions.RoleBinding{ID: id}, nil
}

func (s *testStore) DeleteRoleBinding(_ context.Context, id string) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, id)

	return nil
}

func (s *testStore) GetRoleBinding(_ context.Context, query permissions.RoleBindingQuery) (permissions.RoleBinding, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return permissions.RoleBinding{}, s.forcedError
	}

	require.Equal(s.t, validRoleBindingQuery, query)

	return permissions.RoleBinding{ID: validRoleBindingID}, nil
}

func (s *testStore) GetAccessibleObjects(_ context.Context, query permissions.AccessibleObjectsQuery) ([]string, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return nil, s.forcedError
	}

	require.Equal(s.t, validAccessibleObjectsQuery, query)

	return []string{"object1"}, nil
}

func requireEventFired(t *testing.T, wantEvent permissions.EventType, listener *testListener) {
	t.Helper()

	require.NotEmpty(t, listener.eventFired)
	require.Equal(t, wantEvent, listener.eventFired.Type)
	require.NotEmpty(t, listener.eventFired.Actor)
	require.NotEmpty(t, listener.eventFired.RoleBinding)
	require.NotEmpty(t, listener.eventFired.Timestamp)
}
