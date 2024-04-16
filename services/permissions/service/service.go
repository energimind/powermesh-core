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
	CreateRoleBinding(ctx context.Context, id string, data permissions.RoleBindingData) (permissions.RoleBinding, error)
	UpdateRoleBinding(ctx context.Context, id string, data permissions.RoleBindingData) (permissions.RoleBinding, error)
	DeleteRoleBinding(ctx context.Context, id string) error
	GetRoleBinding(ctx context.Context, query permissions.RoleBindingQuery) (permissions.RoleBinding, error)
	GetAccessibleObjects(ctx context.Context, query permissions.AccessibleObjectsQuery) ([]string, error)
}

// listener defines the external permissions event listener.
type listener interface {
	HandlePermissionEvent(ctx context.Context, event permissions.Event) error
}

// PermissionService implements the permissions service.
//
// It implements the permissions.RoleBindingService interface.
//
// We do not wrap the errors returned by the repository because they are already
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

	roleBinding, err := s.store.CreateRoleBinding(ctx, s.idGen.GenerateID(), data)
	if err != nil {
		return permissions.RoleBinding{}, err
	}

	event := permissions.Event{
		Type:        permissions.RoleBindingCreated,
		Actor:       actor,
		Timestamp:   s.now(),
		RoleBinding: roleBinding,
	}

	if err := s.listener.HandlePermissionEvent(ctx, event); err != nil {
		return permissions.RoleBinding{}, errorz.NewInternalError("permission event handler failed: %v", err)
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

	roleBinding, err := s.store.UpdateRoleBinding(ctx, id, data)
	if err != nil {
		return permissions.RoleBinding{}, err
	}

	event := permissions.Event{
		Type:        permissions.RoleBindingUpdated,
		Actor:       actor,
		Timestamp:   s.now(),
		RoleBinding: roleBinding,
	}

	if err := s.listener.HandlePermissionEvent(ctx, event); err != nil {
		return permissions.RoleBinding{}, errorz.NewInternalError("permission event handler failed: %v", err)
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
		return errorz.NewInternalError("permission event handler failed: %v", err)
	}

	return nil
}

// GetRoleBinding implements the permissions.RoleBindingService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) GetRoleBinding(
	ctx context.Context,
	actor access.Actor,
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

// GetAccessibleObjects implements the permissions.RoleBindingService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *PermissionService) GetAccessibleObjects(
	ctx context.Context,
	actor access.Actor,
	query permissions.AccessibleObjectsQuery,
) ([]string, error) {
	if err := validateAccessibleObjectsQuery(query); err != nil {
		return nil, err
	}

	objects, err := s.store.GetAccessibleObjects(ctx, query)
	if err != nil {
		return nil, err
	}

	return objects, nil
}
