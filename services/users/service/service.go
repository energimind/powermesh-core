package service

import (
	"context"
	"time"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/users"
)

// idGenerator defines the external ID generator.
type idGenerator interface {
	GenerateID() string
}

// store defines the external user store.
type store interface {
	CreateUser(ctx context.Context, id string, data users.UserData) (users.User, error)
	UpdateUser(ctx context.Context, id string, data users.UserData) (users.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUserByUsername(ctx context.Context, username string) (users.User, error)
}

// listener defines the external user event listener.
type listener interface {
	HandleUserEvent(ctx context.Context, event users.Event) error
}

// UserService implements the user service.
//
// It implements the users.UserService interface.
//
// We do not wrap the errors returned by the repository because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type UserService struct {
	idGen    idGenerator
	store    store
	listener listener
	now      func() time.Time
}

// NewUserService creates a new user service.
func NewUserService(idGen idGenerator, store store, listener listener) *UserService {
	return &UserService{
		idGen:    idGen,
		store:    store,
		listener: listener,
		now:      time.Now,
	}
}

// CreateUser implements the users.UserService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserService) CreateUser(
	ctx context.Context,
	actor access.Actor,
	data users.UserData,
) (users.User, error) {
	if !canAct(actor) {
		return users.User{}, errorz.NewAccessDeniedError("actor %s cannot create a user", actor.String())
	}

	if err := validateUserData(data); err != nil {
		return users.User{}, err
	}

	user, err := s.store.CreateUser(ctx, s.idGen.GenerateID(), data)
	if err != nil {
		return users.User{}, err
	}

	if err := s.listener.HandleUserEvent(ctx, users.Event{
		Type:      users.UserCreated,
		Actor:     actor,
		User:      user,
		Timestamp: s.now(),
	}); err != nil {
		return users.User{}, errorz.NewInternalError("user event handler failed: %v", err)
	}

	return user, nil
}

// UpdateUser implements the users.UserService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserService) UpdateUser(
	ctx context.Context,
	actor access.Actor,
	id string,
	data users.UserData,
) (users.User, error) {
	if !canAct(actor) {
		return users.User{}, errorz.NewAccessDeniedError("actor %s cannot update a user", actor.String())
	}

	if err := validateID(id); err != nil {
		return users.User{}, err
	}

	if err := validateUserData(data); err != nil {
		return users.User{}, err
	}

	user, err := s.store.UpdateUser(ctx, id, data)
	if err != nil {
		return users.User{}, err
	}

	if err := s.listener.HandleUserEvent(ctx, users.Event{
		Type:      users.UserUpdated,
		Actor:     actor,
		User:      user,
		Timestamp: s.now(),
	}); err != nil {
		return users.User{}, errorz.NewInternalError("user event handler failed: %v", err)
	}

	return user, nil
}

// DeleteUser implements the users.UserService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserService) DeleteUser(
	ctx context.Context,
	actor access.Actor,
	id string,
) error {
	if !canAct(actor) {
		return errorz.NewAccessDeniedError("actor %s cannot delete a user", actor.String())
	}

	if err := validateID(id); err != nil {
		return err
	}

	if err := s.store.DeleteUser(ctx, id); err != nil {
		return err
	}

	if err := s.listener.HandleUserEvent(ctx, users.Event{
		Type:      users.UserDeleted,
		Actor:     actor,
		User:      users.User{ID: id},
		Timestamp: s.now(),
	}); err != nil {
		return errorz.NewInternalError("user event handler failed: %v", err)
	}

	return nil
}

// GetUserByUsername implements the users.UserService interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserService) GetUserByUsername(
	ctx context.Context,
	actor access.Actor,
	username string,
) (users.User, error) {
	if !canAct(actor) {
		return users.User{}, errorz.NewAccessDeniedError("actor %s cannot get a user by username", actor.String())
	}

	if err := validateUsername(username); err != nil {
		return users.User{}, err
	}

	user, err := s.store.GetUserByUsername(ctx, username)
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}
