package service

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/users"
	"github.com/stretchr/testify/require"
)

var (
	adminActor    = access.Actor{Role: access.RoleAdmin}
	validUserID   = "1" // must match generated ID from testIDGenerator
	validUserData = users.UserData{
		Username: "user1",
		Email:    "user1@somewhere.com",
	}
	validUser = users.User{
		ID:       validUserID,
		Username: validUserData.Username,
		Email:    validUserData.Email,
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
	eventFired  users.Event
}

// Ensure that the testListener implements the listener interface.
var _ listener = (*testListener)(nil)

func newTestListener(forcedError bool) *testListener {
	var err error

	if forcedError {
		err = errors.New("forced-error")
	}

	return &testListener{forcedError: err}
}

func (l *testListener) HandleUserEvent(_ context.Context, event users.Event) error {
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

func (s *testStore) CreateUser(
	_ context.Context,
	user users.User,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.Equal(s.t, validUser, user)

	return nil
}

func (s *testStore) UpdateUser(
	_ context.Context,
	user users.User,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.Equal(s.t, validUser, user)

	return nil
}

func (s *testStore) DeleteUser(
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

func (s *testStore) GetUser(
	_ context.Context,
	id string,
) (users.User, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return users.User{}, s.forcedError
	}

	require.NotEmpty(s.t, id)

	if id == validUserID {
		return validUser, nil
	}

	return users.User{}, errorz.NewNotFoundError("user not found")
}

func (s *testStore) GetUsersByIDs(
	_ context.Context,
	ids []string,
) ([]users.User, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return nil, s.forcedError
	}

	require.NotEmpty(s.t, ids)

	found := make([]users.User, 0, len(ids))

	for _, id := range ids {
		if id == validUserID {
			found = append(found, validUser)
		}
	}

	return found, nil
}

func (s *testStore) GetUserByUsername(
	_ context.Context,
	username string,
) (users.User, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return users.User{}, s.forcedError
	}

	require.NotEmpty(s.t, username)

	return users.User{ID: validUserID, Username: username}, nil
}

func requireEventFired(t *testing.T, wantEvent users.EventType, listener *testListener) {
	t.Helper()

	eventFired := listener.eventFired

	require.NotEmpty(t, eventFired)

	ue, ok := users.ExtractUserEvent(eventFired)

	require.True(t, ok)

	require.Equal(t, wantEvent, ue.Type)
	require.NotEmpty(t, ue.Actor)
	require.NotEmpty(t, ue.User)
	require.NotEmpty(t, ue.Timestamp)
}
