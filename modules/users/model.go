package users

// User models a user.
type User struct {
	ID          string
	ExternalID  string
	Username    string
	DisplayName string
	Email       string
}
