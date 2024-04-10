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
type Event struct {
	Type      EventType
	Actor     access.Actor
	User      User
	Timestamp time.Time
}
