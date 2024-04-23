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
	validOwnerID         = "userOwner1"
	validRoleBindingID   = "1" // must match generated ID from testIDGenerator
	validRoleBindingData = permissions.RoleBindingData{
		OwnerID:      validOwnerID,
		UserID:       "user1",
		ResourceID:   "resource1",
		ResourceType: permissions.ResourceTypeModel,
		Role:         access.RoleAdmin,
	}
	validRoleBinding = permissions.RoleBinding{
		ID:           validRoleBindingID,
		OwnerID:      validRoleBindingData.OwnerID,
		UserID:       validRoleBindingData.UserID,
		ResourceID:   validRoleBindingData.ResourceID,
		ResourceType: validRoleBindingData.ResourceType,
		Role:         validRoleBindingData.Role,
	}
	validRoleBindingQuery = permissions.RoleBindingQuery{
		UserID:     "user1",
		ResourceID: "resource1",
	}
	validAccessibleResourcesQuery = permissions.AccessibleResourcesQuery{
		UserID:       "user1",
		ResourceType: permissions.ResourceTypeModel,
	}
)

type testIDGenerator struct {
	idCounter atomic.Int64
}

// Ensure that the testIDGenerator implements the idGenerator interface.
var _ idGenerator = (*testIDGenerator)(nil)

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

// Ensure that the testListener implements the listener interface.
var _ listener = (*testListener)(nil)

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

// Ensure that the testStore implements the store interface.
var _ store = (*testStore)(nil)

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

func (s *testStore) CreateRoleBinding(
	_ context.Context,
	roleBinding permissions.RoleBinding,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.Equal(s.t, validRoleBinding, roleBinding)

	return nil
}

func (s *testStore) UpdateRoleBinding(
	_ context.Context,
	roleBinding permissions.RoleBinding,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.Equal(s.t, validRoleBinding, roleBinding)

	return nil
}

func (s *testStore) DeleteRoleBinding(
	_ context.Context,
	id string,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, id)

	return nil
}

func (s *testStore) GetRoleBinding(
	_ context.Context,
	query permissions.RoleBindingQuery,
) (permissions.RoleBinding, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return permissions.RoleBinding{}, s.forcedError
	}

	require.Equal(s.t, validRoleBindingQuery, query)

	return permissions.RoleBinding{ID: validRoleBindingID}, nil
}

func (s *testStore) GetRoleBindingsByOwner(
	_ context.Context,
	ownerID string,
) ([]permissions.RoleBinding, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return nil, s.forcedError
	}

	require.Equal(s.t, validOwnerID, ownerID)

	return []permissions.RoleBinding{{ID: validRoleBindingID}}, nil
}

func (s *testStore) GetAccessibleResources(
	_ context.Context,
	query permissions.AccessibleResourcesQuery,
) ([]string, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return nil, s.forcedError
	}

	require.Equal(s.t, validAccessibleResourcesQuery, query)

	return []string{"resource1"}, nil
}

func requireEventFired(t *testing.T, wantEvent permissions.EventType, listener *testListener) {
	t.Helper()

	eventFired := listener.eventFired

	require.NotEmpty(t, eventFired)

	rbe, ok := permissions.ExtractRoleBindingEvent(eventFired)

	require.True(t, ok)

	require.Equal(t, wantEvent, rbe.Type)
	require.NotEmpty(t, rbe.Actor)
	require.NotEmpty(t, rbe.RoleBinding)
	require.NotEmpty(t, rbe.Timestamp)
}
