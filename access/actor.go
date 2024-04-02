package access

// Actor is a user with a role.
type Actor struct {
	UserID string
	Role   Role
}

// NewActor creates a new actor.
func NewActor(userID string, role Role) Actor {
	return Actor{
		UserID: userID,
		Role:   role,
	}
}

// String returns the string representation of the actor.
func (a Actor) String() string {
	return a.UserID + " (" + a.Role.String() + ")"
}
