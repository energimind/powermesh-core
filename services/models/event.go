package models

import (
	"time"

	"github.com/energimind/powermesh-core/access"
)

// EventType is the type of event that occurred.
type EventType string

// Event types.
const (
	ModelCreated EventType = "model.created"
	ModelUpdated EventType = "model.updated"
	ModelDeleted EventType = "model.deleted"
)

// Event models an event that occurs in the model service.
type Event struct {
	Type      EventType
	Actor     access.Actor
	Model     Model
	Timestamp time.Time
}
