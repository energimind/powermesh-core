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
	validModelData = models.ModelData{
		Code: "code1",
		Name: "name1",
	}
	validModel = models.Model{
		ID:   validModelID,
		Code: validModelData.Code,
		Name: validModelData.Name,
	}
)

type testModelListener struct {
	forcedError error
	eventFired  models.ModelEvent
}

// Ensure that the testModelListener implements the modelListener interface.
var _ modelListener = (*testModelListener)(nil)

func newTestModelListener(forcedError bool) *testModelListener {
	var err error

	if forcedError {
		err = errors.New("forced-error")
	}

	return &testModelListener{
		forcedError: err,
	}
}

func (l *testModelListener) HandleModelEvent(_ context.Context, event models.ModelEvent) error {
	if l.forcedError != nil {
		return l.forcedError
	}

	l.eventFired = event

	return nil
}

type testModelStore struct {
	t           *testing.T
	forcedError error
}

// Ensure that the testModelStore implements the modelStore interface.
var _ modelStore = (*testModelStore)(nil)

func newTestModelStore(t *testing.T, forcedError bool) *testModelStore {
	var err error

	if forcedError {
		err = errorz.NewStoreError("forced-error")
	}

	return &testModelStore{
		t:           t,
		forcedError: err,
	}
}

func (s *testModelStore) CreateModel(
	_ context.Context,
	model models.Model,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, model.ID)
	require.Equal(s.t, validModel, model)

	return nil
}

func (s *testModelStore) UpdateModel(
	_ context.Context,
	model models.Model,
) error {
	s.t.Helper()

	if s.forcedError != nil {
		return s.forcedError
	}

	require.NotEmpty(s.t, model.ID)
	require.Equal(s.t, validModel, model)

	return nil
}

func (s *testModelStore) DeleteModel(
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

func (s *testModelStore) GetModel(
	_ context.Context,
	id string,
) (models.Model, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return models.Model{}, s.forcedError
	}

	require.NotEmpty(s.t, id)

	if id == validModelID {
		return models.Model{ID: id}, nil
	}

	return models.Model{}, errorz.NewNotFoundError("model %v not found", id)
}

func (s *testModelStore) GetModelsByIDs(
	_ context.Context,
	ids []string,
) ([]models.Model, error) {
	s.t.Helper()

	if s.forcedError != nil {
		return nil, s.forcedError
	}

	require.NotEmpty(s.t, ids)

	found := make([]models.Model, 0, len(ids))

	for _, id := range ids {
		if id == validModelID {
			found = append(found, validModel)
		}
	}

	return found, nil
}

func requireModelEventFired(t *testing.T, wantEvent models.EventType, listener *testModelListener) {
	t.Helper()

	require.NotEmpty(t, listener.eventFired)
	require.Equal(t, wantEvent, listener.eventFired.Type)
	require.NotEmpty(t, listener.eventFired.Actor)
	require.NotEmpty(t, listener.eventFired.Model)
	require.NotEmpty(t, listener.eventFired.Timestamp)
}
