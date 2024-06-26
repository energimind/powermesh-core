package service

import (
	"context"
	"time"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/models"
)

// modelStore defines the external model store.
type modelStore interface {
	CreateModel(ctx context.Context, model models.Model) error
	UpdateModel(ctx context.Context, model models.Model) error
	DeleteModel(ctx context.Context, id string) error
	GetModel(ctx context.Context, id string) (models.Model, error)
	GetModelsByIDs(ctx context.Context, ids []string) ([]models.Model, error)
}

// modelListener defines the external model event modelListener.
type modelListener interface {
	HandleModelEvent(ctx context.Context, event models.ModelEvent) error
}

// ModelService implements the model service.
//
// It implements the models.ModelService interface.
//
// We do not wrap the errors returned by the store because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type ModelService struct {
	idGen    idGenerator
	store    modelStore
	listener modelListener
	now      func() time.Time
}

// Ensure ModelService implements the models.ModelService interface.
var _ models.ModelService = (*ModelService)(nil)

// NewModelService creates a new model service.
func NewModelService(store modelStore, idGen idGenerator, opts ...ModelServiceOption) *ModelService {
	svc := &ModelService{
		idGen: idGen,
		store: store,
		now:   time.Now,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

// CreateModel implements the models.ModelService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelService) CreateModel(
	ctx context.Context,
	actor access.Actor,
	data models.ModelData,
) (models.Model, error) {
	if err := validateModelData(data); err != nil {
		return models.Model{}, err
	}

	model := modelFromData(s.idGen.GenerateID(), data)

	if err := s.store.CreateModel(ctx, model); err != nil {
		return models.Model{}, err
	}

	if err := s.fireModelEvent(ctx, actor, models.ModelCreated, model); err != nil {
		return models.Model{}, err
	}

	return model, nil
}

// UpdateModel implements the models.ModelService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelService) UpdateModel(
	ctx context.Context,
	actor access.Actor,
	id string,
	data models.ModelData,
) (models.Model, error) {
	if err := validateID(id); err != nil {
		return models.Model{}, err
	}

	if err := validateModelData(data); err != nil {
		return models.Model{}, err
	}

	model := modelFromData(id, data)

	if err := s.store.UpdateModel(ctx, model); err != nil {
		return models.Model{}, err
	}

	if err := s.fireModelEvent(ctx, actor, models.ModelUpdated, model); err != nil {
		return models.Model{}, err
	}

	return model, nil
}

// DeleteModel implements the models.ModelService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelService) DeleteModel(
	ctx context.Context,
	actor access.Actor,
	id string,
) error {
	if err := validateID(id); err != nil {
		return err
	}

	if err := s.store.DeleteModel(ctx, id); err != nil {
		return err
	}

	if err := s.fireModelEvent(ctx, actor, models.ModelDeleted, models.Model{ID: id}); err != nil {
		return err
	}

	return nil
}

// GetModel implements the models.ModelService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelService) GetModel(
	ctx context.Context,
	id string,
) (models.Model, error) {
	if err := validateID(id); err != nil {
		return models.Model{}, err
	}

	model, err := s.store.GetModel(ctx, id)
	if err != nil {
		return models.Model{}, err
	}

	return model, nil
}

// GetModelsByIDs implements the models.ModelService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelService) GetModelsByIDs(
	ctx context.Context,
	ids []string,
) ([]models.Model, error) {
	if len(ids) == 0 {
		return []models.Model{}, nil
	}

	for _, id := range ids {
		if err := validateID(id); err != nil {
			return nil, err
		}
	}

	found, err := s.store.GetModelsByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return found, nil
}

// fireModelEvent fires a model event.
func (s *ModelService) fireModelEvent(
	ctx context.Context,
	actor access.Actor,
	eventType models.EventType,
	model models.Model,
) error {
	if s.listener == nil {
		return nil
	}

	event := models.ModelEvent{
		EventHeader: models.EventHeader{
			Type:      eventType,
			Actor:     actor,
			Timestamp: s.now(),
		},
		Model: model,
	}

	if err := s.listener.HandleModelEvent(ctx, event); err != nil {
		return errorz.NewInternalError("%s event handler failed: %v", eventType, err)
	}

	return nil
}
