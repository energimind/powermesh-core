package mongo

import (
	"context"

	"github.com/energimind/powermesh-core/modules/permissions"
	q "github.com/energimind/powermesh-core/mongoquery"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collPermissions   = "permissions"
	fieldOwnerID      = "ownerId"
	fieldUserID       = "userId"
	fieldResourceID   = "resourceId"
	fieldResourceType = "resourceType"
)

// PermissionStore is a MongoDB implementation of the permissions store.
//
// We do not wrap the errors returned by mongoquery utilities because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type PermissionStore struct {
	permissions *mongo.Collection
}

// NewPermissionStore creates a new MongoDB permissions store.
func NewPermissionStore(db *mongo.Database) *PermissionStore {
	return &PermissionStore{
		permissions: db.Collection(collPermissions),
	}
}

// CreateRoleBinding implements the permissions store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionStore) CreateRoleBinding(ctx context.Context, roleBinding permissions.RoleBinding) error {
	return q.CreateOne(s.permissions, toStoreRoleBinding).Exec(ctx, roleBinding)
}

// UpdateRoleBinding implements the permissions store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionStore) UpdateRoleBinding(ctx context.Context, roleBinding permissions.RoleBinding) error {
	return q.UpdateOne(s.permissions, toStoreRoleBinding).Exec(ctx, roleBinding.ID, roleBinding)
}

// DeleteRoleBinding implements the permissions store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionStore) DeleteRoleBinding(ctx context.Context, id string) error {
	return q.DeleteOne(s.permissions).Exec(ctx, id)
}

// DeleteRoleBindingsByResource implements the permissions store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionStore) DeleteRoleBindingsByResource(
	ctx context.Context,
	resourceID string,
	resourceType permissions.ResourceType,
) error {
	filter := q.Filter{}.
		EQ(fieldResourceID, resourceID).
		EQ(fieldResourceType, resourceType)

	_, err := q.DeleteMany(s.permissions).Exec(ctx, filter)

	return err
}

// GetRoleBinding implements the permissions store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionStore) GetRoleBinding(
	ctx context.Context,
	query permissions.RoleBindingQuery,
) (permissions.RoleBinding, error) {
	filter := q.Filter{}.
		EQ(fieldUserID, query.UserID).
		EQ(fieldResourceID, query.ResourceID).
		EQ(fieldResourceType, query.ResourceType)

	return q.GetOne(s.permissions, fromStoreRoleBinding).Exec(ctx, filter)
}

// GetRoleBindingsByOwner implements the permissions store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionStore) GetRoleBindingsByOwner(ctx context.Context, ownerID string) (
	[]permissions.RoleBinding,
	error,
) {
	return q.FindMany(s.permissions, fromStoreRoleBinding).Exec(ctx, q.Filter{}.EQ(fieldOwnerID, ownerID))
}

// GetAccessibleResources implements the permissions store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionStore) GetAccessibleResources(
	ctx context.Context,
	query permissions.AccessibleResourcesQuery,
) ([]string, error) {
	filter := q.Filter{}.
		EQ(fieldUserID, query.UserID).
		EQ(fieldResourceType, query.ResourceType)

	return q.FindMany(s.permissions, projectResourceID).WithProjection(fieldResourceID).Exec(ctx, filter)
}
