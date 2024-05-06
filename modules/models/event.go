package models

import (
	"time"

	"github.com/energimind/powermesh-core/access"
)

// EventType is the type of event that occurred.
type EventType string

// Event types.
const (
	ModelCreated        EventType = "model.created"
	ModelUpdated        EventType = "model.updated"
	ModelDeleted        EventType = "model.deleted"
	MeshCreated         EventType = "mesh.created"
	MeshUpdated         EventType = "mesh.updated"
	MeshDeleted         EventType = "mesh.deleted"
	MeshContentsCreated EventType = "mesh-contents.created"
	MeshContentsUpdated EventType = "mesh-contents.updated"
	MeshContentsDeleted EventType = "mesh-contents.deleted"
)

// Event models an event that occurs in the models service.
type Event interface {
	IsModelEvent() bool
	IsMeshEvent() bool
}

// EventHeader models the header of an event.
type EventHeader struct {
	Type      EventType
	Actor     access.Actor
	Timestamp time.Time
}

// ModelEvent models an event that occurs in the models service related to a model.
type ModelEvent struct {
	EventHeader
	Model Model
}

// IsModelEvent implements the Event interface.
func (ModelEvent) IsModelEvent() bool {
	return true
}

// IsMeshEvent implements the Event interface.
func (ModelEvent) IsMeshEvent() bool {
	return false
}

// MeshEvent models an event that occurs in the models service related to a mesh.
type MeshEvent struct {
	EventHeader
	Updates Mesh
	Deletes Mesh
}

// IsModelEvent implements the Event interface.
func (MeshEvent) IsModelEvent() bool {
	return false
}

// IsMeshEvent implements the Event interface.
func (MeshEvent) IsMeshEvent() bool {
	return true
}

// ExtractModelEvent extracts a model event from an event.
func ExtractModelEvent(e Event) (ModelEvent, bool) {
	if me, ok := e.(ModelEvent); ok {
		return me, true
	}

	return ModelEvent{}, false
}

// ExtractMeshEvent extracts a mesh event from an event.
func ExtractMeshEvent(e Event) (MeshEvent, bool) {
	if me, ok := e.(MeshEvent); ok {
		return me, true
	}

	return MeshEvent{}, false
}
