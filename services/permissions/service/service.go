package service

import (
	"context"
	"time"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/permissions"
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
// It implements the permissions.RoleBindingService interface.
//
// We do not wrap the errors returned by the store because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type PermissionService struct {
	idGen    idGenerator
	store    store
	listener listener
	now      func() time.Time
}

// NewPermissionService creates a new permissions service.
func NewPermissionService(idGen idGenerator, store store, listener listener) *PermissionService {
	return &PermissionService{
		idGen:    idGen,
		store:    store,
		listener: listener,
		now:      time.Now,
	}
}

// CreateRoleBinding implements the permissions.RoleBindingService interface.
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

	event := permissions.Event{
		Type:        permissions.RoleBindingCreated,
		Actor:       actor,
		Timestamp:   s.now(),
		RoleBinding: roleBinding,
	}

	if err := s.listener.HandlePermissionEvent(ctx, event); err != nil {
		return permissions.RoleBinding{}, errorz.NewInternalError("role-binding.created event handler failed: %v", err)
	}

	return roleBinding, nil
}

// UpdateRoleBinding implements the permissions.RoleBindingService interface.
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

	event := permissions.Event{
		Type:        permissions.RoleBindingUpdated,
		Actor:       actor,
		Timestamp:   s.now(),
		RoleBinding: roleBinding,
	}

	if err := s.listener.HandlePermissionEvent(ctx, event); err != nil {
		return permissions.RoleBinding{}, errorz.NewInternalError("role-binding.updated event handler failed: %v", err)
	}

	return roleBinding, nil
}

// DeleteRoleBinding implements the permissions.RoleBindingService interface.
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

	err := s.store.DeleteRoleBinding(ctx, id)
	if err != nil {
		return err
	}

	event := permissions.Event{
		Type:        permissions.RoleBindingDeleted,
		Actor:       actor,
		Timestamp:   s.now(),
		RoleBinding: permissions.RoleBinding{ID: id},
	}

	if err := s.listener.HandlePermissionEvent(ctx, event); err != nil {
		return errorz.NewInternalError("role-binding.deleted event handler failed: %v", err)
	}

	return nil
}

// GetRoleBinding implements the permissions.RoleBindingService interface.
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

// GetRoleBindingsByOwner implements the permissions.RoleBindingService interface.
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

// GetAccessibleResources implements the permissions.RoleBindingService interface.
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
