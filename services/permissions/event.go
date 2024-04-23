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
type Event interface {
	IsRoleBindingEvent() bool
}

// EventHeader models the header of an event.
type EventHeader struct {
	Type      EventType
	Actor     access.Actor
	Timestamp time.Time
}

// RoleBindingEvent models an event that occurs in the permissions service related to a role binding.
type RoleBindingEvent struct {
	EventHeader
	RoleBinding RoleBinding
}

// IsRoleBindingEvent implements the Event interface.
func (RoleBindingEvent) IsRoleBindingEvent() bool {
	return true
}

// ExtractRoleBindingEvent extracts a role binding event from an event.
func ExtractRoleBindingEvent(e Event) (RoleBindingEvent, bool) {
	if rbe, ok := e.(RoleBindingEvent); ok {
		return rbe, true
	}

	return RoleBindingEvent{}, false
}
