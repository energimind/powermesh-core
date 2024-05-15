package permissions

import (
	"context"

	"github.com/energimind/powermesh-core/access"
)

// PermissionService defines the role binding service.
type PermissionService interface {
	CreateRoleBinding(ctx context.Context, actor access.Actor, data RoleBindingData) (RoleBinding, error)
	UpdateRoleBinding(ctx context.Context, actor access.Actor, id string, data RoleBindingData) (RoleBinding, error)
	DeleteRoleBinding(ctx context.Context, actor access.Actor, id string) error
	DeleteRoleBindingsByResource(ctx context.Context, actor access.Actor, resourceID string, resourceType ResourceType) error
	GetRoleBinding(ctx context.Context, query RoleBindingQuery) (RoleBinding, error)
	GetRoleBindingsByOwner(ctx context.Context, ownerID string) ([]RoleBinding, error)
	GetAccessibleResources(ctx context.Context, query AccessibleResourcesQuery) ([]string, error)
}

// RoleBindingData defines the role binding data.
// It is used to create or update a role binding.
type RoleBindingData struct {
	OwnerID      string
	UserID       string
	ResourceID   string
	ResourceType ResourceType
	Role         access.Role
}

// RoleBindingQuery defines the role binding query.
// It is used to get a role binding for a user and a resource.
type RoleBindingQuery struct {
	UserID       string
	ResourceID   string
	ResourceType ResourceType
}

// AccessibleResourcesQuery defines the accessible resources query.
// It is used to get all accessible resources of the given type for a user.
type AccessibleResourcesQuery struct {
	UserID       string
	ResourceType ResourceType
}
