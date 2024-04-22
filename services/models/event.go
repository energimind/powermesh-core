package models

import (
	"time"

	"github.com/energimind/powermesh-core/access"
)

// ModelEventType is the type of event that occurred.
type ModelEventType string

// Model event types.
const (
	ModelCreated ModelEventType = "model.created"
	ModelUpdated ModelEventType = "model.updated"
	ModelDeleted ModelEventType = "model.deleted"
)

// ModelEvent models an event that occurs in the model service.
type ModelEvent struct {
	Type      ModelEventType
	Actor     access.Actor
	Model     Model
	Timestamp time.Time
}
