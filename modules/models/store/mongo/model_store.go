package mongo

import (
	"context"

	"github.com/energimind/powermesh-core/modules/models"
	q "github.com/energimind/powermesh-core/mongoquery"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collModels = "models"
)

// ModelStore is a MongoDB implementation of the model store.
//
// We do not wrap the errors returned by mongoquery utilities because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type ModelStore struct {
	models *mongo.Collection
}

// NewModelStore creates a new MongoDB model store.
func NewModelStore(db *mongo.Database) *ModelStore {
	return &ModelStore{
		models: db.Collection(collModels),
	}
}

// CreateModel implements the model store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelStore) CreateModel(ctx context.Context, model models.Model) error {
	return q.CreateOne(s.models, toStoreModel).Exec(ctx, model)
}

// UpdateModel implements the model store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelStore) UpdateModel(ctx context.Context, model models.Model) error {
	return q.UpdateOne(s.models, toStoreModel).Exec(ctx, model.ID, model)
}

// DeleteModel implements the model store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelStore) DeleteModel(ctx context.Context, id string) error {
	return q.DeleteOne(s.models).Exec(ctx, id)
}

// GetModel implements the model store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelStore) GetModel(ctx context.Context, id string) (models.Model, error) {
	return q.GetOne(s.models, fromStoreModel).Exec(ctx, id)
}

// GetModelsByIDs implements the model store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *ModelStore) GetModelsByIDs(ctx context.Context, ids []string) ([]models.Model, error) {
	return q.FindMany(s.models, fromStoreModel).Exec(ctx, q.Filter{}.IN(fieldID, ids))
}
