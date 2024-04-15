package permissions

import (
	"context"

	"github.com/energimind/powermesh-core/access"
)

// RoleBindingService defines the role binding service.
type RoleBindingService interface {
	CreateRoleBinding(ctx context.Context, actor access.Actor, data RoleBindingData) (RoleBinding, error)
	UpdateRoleBinding(ctx context.Context, actor access.Actor, id string, data RoleBindingData) (RoleBinding, error)
	DeleteRoleBinding(ctx context.Context, actor access.Actor, id string) error
	GetRoleBinding(ctx context.Context, actor access.Actor, query RoleBindingQuery) (RoleBinding, error)
	GetAccessibleObjects(ctx context.Context, actor access.Actor, query AccessibleObjectsQuery) ([]string, error)
}

// RoleBindingData defines the role binding data.
// It is used to create or update a role binding.
type RoleBindingData struct {
	UserID     string
	ObjectID   string
	ObjectType ObjectType
	Role       access.Role
}

// RoleBindingQuery defines the role binding query.
// It is used to get a role binding for a user and an object.
type RoleBindingQuery struct {
	UserID     string
	ObjectID   string
	ObjectType ObjectType
}

// AccessibleObjectsQuery defines the accessible objects query.
// It is used to get all accessible objects of the given type for a user.
type AccessibleObjectsQuery struct {
	UserID     string
	ObjectType ObjectType
}
