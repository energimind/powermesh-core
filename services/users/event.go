package users

import (
	"time"

	"github.com/energimind/powermesh-core/access"
)

// EventType is the type of event that occurred.
type EventType string

// Event types.
const (
	UserCreated EventType = "user.created"
	UserUpdated EventType = "user.updated"
	UserDeleted EventType = "user.deleted"
)

// Event models an event that occurs in the user service.
type Event interface {
	IsUserEvent() bool
}

// EventHeader models the header of an event.
type EventHeader struct {
	Type      EventType
	Actor     access.Actor
	Timestamp time.Time
}

// UserEvent models an event that occurs in the user service related to a user.
type UserEvent struct {
	EventHeader
	User User
}

// IsUserEvent implements the Event interface.
func (UserEvent) IsUserEvent() bool {
	return true
}

// ExtractUserEvent extracts a user event from an event.
func ExtractUserEvent(e Event) (UserEvent, bool) {
	if ue, ok := e.(UserEvent); ok {
		return ue, true
	}

	return UserEvent{}, false
}
