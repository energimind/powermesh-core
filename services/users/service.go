package users

import (
	"context"

	"github.com/energimind/powermesh-core/access"
)

// UserService defines the user service.
type UserService interface {
	CreateUser(ctx context.Context, actor access.Actor, data UserData) (User, error)
	UpdateUser(ctx context.Context, actor access.Actor, id string, data UserData) (User, error)
	DeleteUser(ctx context.Context, actor access.Actor, id string) error
	GetUserByUsername(ctx context.Context, actor access.Actor, username string) (User, error)
}

// UserData defines the user data. It is used to create or update a user.
type UserData struct {
	Username string
	Email    string
}
