package mongo

import (
	"context"

	"github.com/energimind/powermesh-core/modules/users"
	q "github.com/energimind/powermesh-core/mongoquery"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collUsers     = "users"
	fieldID       = "id"
	fieldUsername = "username"
)

// UserStore is a MongoDB implementation of the users store.
//
// We do not wrap the errors returned by mongoquery utilities because they are already
// packed as domain errors. Therefore, we disable the wrapcheck linter for these calls.
type UserStore struct {
	users *mongo.Collection
}

// NewUserStore creates a new MongoDB users store.
func NewUserStore(db *mongo.Database) *UserStore {
	return &UserStore{
		users: db.Collection(collUsers),
	}
}

// CreateUser implements the users store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserStore) CreateUser(ctx context.Context, user users.User) error {
	return q.CreateOne(s.users, toStoreUser).Exec(ctx, user)
}

// UpdateUser implements the users store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserStore) UpdateUser(ctx context.Context, user users.User) error {
	return q.UpdateOne(s.users, toStoreUser).Exec(ctx, user.ID, user)
}

// DeleteUser implements the users store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserStore) DeleteUser(ctx context.Context, id string) error {
	return q.DeleteOne(s.users).Exec(ctx, id)
}

// GetUser implements the users store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserStore) GetUser(ctx context.Context, id string) (users.User, error) {
	return q.GetOne(s.users, fromStoreUser).Exec(ctx, id)
}

// GetUsersByIDs implements the users store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserStore) GetUsersByIDs(ctx context.Context, ids []string) ([]users.User, error) {
	return q.FindMany(s.users, fromStoreUser).Exec(ctx, q.Filter{}.IN(fieldID, ids))
}

// GetUserByUsername implements the users store interface.
//
//nolint:wrapcheck // see comment in the header
func (s *UserStore) GetUserByUsername(ctx context.Context, username string) (users.User, error) {
	return q.GetOne(s.users, fromStoreUser).Exec(ctx, q.Filter{}.EQ(fieldUsername, username))
}
