package service

import (
	"context"
	"time"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/permissions"
)

// idGenerator defines the external ID generator.
type idGenerator interface {
	GenerateID() string
}

// store defines the external permissions store.
type store interface {
	CreateRoleBinding(ctx context.Context, binding permissions.RoleBinding) error
	UpdateRoleBinding(ctx context.Context, binding permissions.RoleBinding) error
	DeleteRoleBinding(ctx context.Context, id string) error
	DeleteRoleBindingsByResource(ctx context.Context, resourceID string, resourceType permissions.ResourceType) error
	GetRoleBinding(ctx context.Context, query permissions.RoleBindingQuery) (permissions.RoleBinding, error)
	GetRoleBindingsByOwner(ctx context.Context, ownerID string) ([]permissions.RoleBinding, error)
	GetAccessibleResources(ctx context.Context, query permissions.AccessibleResourcesQuery) ([]string, error)
}

// listener defines the external permissions event listener.
type listener interface {
	HandlePermissionEvent(ctx context.Context, event permissions.Event) error
}

// PermissionService implements the permissions service.
//
// It implements the permissions.PermissionService interface.
//
// We do not wrap the errors returned by the store because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type PermissionService struct {
	idGen    idGenerator
	store    store
	listener listener
	now      func() time.Time
}

// Ensure PermissionService implements the permissions.PermissionService interface.
var _ permissions.PermissionService = (*PermissionService)(nil)

// NewPermissionService creates a new permissions service.
func NewPermissionService(store store, idGen idGenerator, opts ...Option) *PermissionService {
	svc := &PermissionService{
		idGen: idGen,
		store: store,
		now:   time.Now,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

// CreateRoleBinding implements the permissions.PermissionService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) CreateRoleBinding(
	ctx context.Context,
	actor access.Actor,
	data permissions.RoleBindingData,
) (permissions.RoleBinding, error) {
	if err := validateRoleBindingData(data); err != nil {
		return permissions.RoleBinding{}, err
	}

	roleBinding := roleBindingFromData(s.idGen.GenerateID(), data)

	if err := s.store.CreateRoleBinding(ctx, roleBinding); err != nil {
		return permissions.RoleBinding{}, err
	}

	if err := s.fireRoleBindingEvent(
		ctx,
		actor,
		permissions.RoleBindingCreated,
		roleBinding,
	); err != nil {
		return permissions.RoleBinding{}, err
	}

	return roleBinding, nil
}

// UpdateRoleBinding implements the permissions.PermissionService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) UpdateRoleBinding(
	ctx context.Context,
	actor access.Actor,
	id string,
	data permissions.RoleBindingData,
) (permissions.RoleBinding, error) {
	if err := validateID(id); err != nil {
		return permissions.RoleBinding{}, err
	}

	if err := validateRoleBindingData(data); err != nil {
		return permissions.RoleBinding{}, err
	}

	roleBinding := roleBindingFromData(id, data)

	if err := s.store.UpdateRoleBinding(ctx, roleBinding); err != nil {
		return permissions.RoleBinding{}, err
	}

	if err := s.fireRoleBindingEvent(
		ctx,
		actor,
		permissions.RoleBindingUpdated,
		roleBinding,
	); err != nil {
		return permissions.RoleBinding{}, err
	}

	return roleBinding, nil
}

// DeleteRoleBinding implements the permissions.PermissionService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) DeleteRoleBinding(
	ctx context.Context,
	actor access.Actor,
	id string,
) error {
	if err := validateID(id); err != nil {
		return err
	}

	if err := s.store.DeleteRoleBinding(ctx, id); err != nil {
		return err
	}

	if err := s.fireRoleBindingEvent(
		ctx,
		actor,
		permissions.RoleBindingDeleted,
		permissions.RoleBinding{ID: id},
	); err != nil {
		return err
	}

	return nil
}

// DeleteRoleBindingsByResource implements the permissions.PermissionService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) DeleteRoleBindingsByResource(
	ctx context.Context,
	actor access.Actor,
	resourceID string,
	resourceType permissions.ResourceType,
) error {
	if err := validateResourceID(resourceID); err != nil {
		return err
	}

	if err := s.store.DeleteRoleBindingsByResource(ctx, resourceID, resourceType); err != nil {
		return err
	}

	if err := s.fireRoleBindingEvent(
		ctx,
		actor,
		permissions.RoleBindingDeleted,
		permissions.RoleBinding{ResourceID: resourceID, ResourceType: resourceType},
	); err != nil {
		return err
	}

	return nil
}

// GetRoleBinding implements the permissions.PermissionService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) GetRoleBinding(
	ctx context.Context,
	query permissions.RoleBindingQuery,
) (permissions.RoleBinding, error) {
	if err := validateRoleBindingQuery(query); err != nil {
		return permissions.RoleBinding{}, err
	}

	roleBinding, err := s.store.GetRoleBinding(ctx, query)
	if err != nil {
		return permissions.RoleBinding{}, err
	}

	return roleBinding, nil
}

// GetRoleBindingsByOwner implements the permissions.PermissionService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) GetRoleBindingsByOwner(
	ctx context.Context,
	ownerID string,
) ([]permissions.RoleBinding, error) {
	if err := validateID(ownerID); err != nil {
		return nil, err
	}

	roleBindings, err := s.store.GetRoleBindingsByOwner(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	return roleBindings, nil
}

// GetAccessibleResources implements the permissions.PermissionService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) GetAccessibleResources(
	ctx context.Context,
	query permissions.AccessibleResourcesQuery,
) ([]string, error) {
	if err := validateAccessibleResourcesQuery(query); err != nil {
		return nil, err
	}

	resources, err := s.store.GetAccessibleResources(ctx, query)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

// fireRoleBindingEvent fires a role-binding event.
func (s *PermissionService) fireRoleBindingEvent(
	ctx context.Context,
	actor access.Actor,
	eventType permissions.EventType,
	roleBinding permissions.RoleBinding,
) error {
	event := permissions.RoleBindingEvent{
		EventHeader: permissions.EventHeader{
			Type:      eventType,
			Actor:     actor,
			Timestamp: s.now(),
		},
		RoleBinding: roleBinding,
	}

	if err := s.listener.HandlePermissionEvent(ctx, event); err != nil {
		return errorz.NewInternalError("%s event handler failed: %v", eventType, err)
	}

	return nil
}
