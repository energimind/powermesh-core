package permissions

import (
	"time"

	"github.com/energimind/powermesh-core/access"
)

// EventType is the type of event that occurred.
type EventType string

// Event types.
const (
	RoleBindingCreated EventType = "role-binding.created"
	RoleBindingUpdated EventType = "role-binding.updated"
	RoleBindingDeleted EventType = "role-binding.deleted"
)

// Event models an event that occurs in the permissions service.
type Event struct {
	Type        EventType
	Actor       access.Actor
	Timestamp   time.Time
	RoleBinding RoleBinding
}
