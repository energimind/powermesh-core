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

// MeshEventType is the type of event that occurred.
type MeshEventType string

// Mesh event types.
const (
	MeshCreated         MeshEventType = "mesh.created"
	MeshUpdated         MeshEventType = "mesh.updated"
	MeshDeleted         MeshEventType = "mesh.deleted"
	MeshContentsCreated MeshEventType = "mesh-contents.created"
	MeshContentsUpdated MeshEventType = "mesh-contents.updated"
	MeshContentsDeleted MeshEventType = "mesh-contents.deleted"
)

// MeshEvent models an event that occurs in the mesh service.
type MeshEvent struct {
	Type      MeshEventType
	Actor     access.Actor
	Updates   Mesh
	Deletes   Mesh
	Timestamp time.Time
}
