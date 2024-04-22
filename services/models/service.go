package models

import (
	"context"

	"github.com/energimind/powermesh-core/access"
)

// ModelService defines a model service.
type ModelService interface {
	CreateModel(ctx context.Context, actor access.Actor, data ModelData) (Model, error)
	UpdateModel(ctx context.Context, actor access.Actor, id string, data ModelData) (Model, error)
	DeleteModel(ctx context.Context, actor access.Actor, id string) error
	GetModel(ctx context.Context, id string) (Model, error)
	GetModelsByIDs(ctx context.Context, ids []string) ([]Model, error)
}

// ModelData defines the model data. It is used to create or update a model.
type ModelData struct {
	Code string
	Name string
}
