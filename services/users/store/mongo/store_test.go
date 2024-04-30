package mongo_test

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/users/store/mongo"
	"github.com/stretchr/testify/require"
)

func TestUserStore_CreateUser(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.UserStore) {
		require.NoError(t, store.CreateUser(ctx, testUser()))
	})
}

func TestUserStore_UpdateUser(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.UserStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.UpdateUser(ctx, testUser()))
		})

		t.Run("success", func(t *testing.T) {
			user := testUser()

			require.NoError(t, store.CreateUser(ctx, user))

			user.Username = "new-username"

			require.NoError(t, store.UpdateUser(ctx, user))

			updatedUser, err := store.GetUser(ctx, user.ID)

			require.NoError(t, err)
			require.Equal(t, user, updatedUser)
		})
	})
}

func TestUserStore_DeleteUser(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.UserStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.DeleteUser(ctx, "missing"))
		})

		t.Run("success", func(t *testing.T) {
			user := testUser()

			require.NoError(t, store.CreateUser(ctx, user))

			require.NoError(t, store.DeleteUser(ctx, user.ID))

			_, err := store.GetUser(ctx, user.ID)

			require.Error(t, err)
		})
	})
}

func TestUserStore_GetUser(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.UserStore) {
		t.Run("not-found", func(t *testing.T) {
			_, err := store.GetUser(ctx, "missing")

			require.IsType(t, errorz.NotFoundError{}, err)
		})

		t.Run("success", func(t *testing.T) {
			user := testUser()

			require.NoError(t, store.CreateUser(ctx, user))

			createdUser, err := store.GetUser(ctx, user.ID)

			require.NoError(t, err)
			require.Equal(t, user, createdUser)
		})
	})
}

func TestUserStore_GetUsersByIDs(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.UserStore) {
		t.Run("not-found", func(t *testing.T) {
			users, err := store.GetUsersByIDs(ctx, []string{"non-existent"})

			require.NoError(t, err)
			require.Empty(t, users)
		})

		t.Run("success", func(t *testing.T) {
			user := testUser()

			require.NoError(t, store.CreateUser(ctx, user))
			require.NoError(t, store.CreateUser(ctx, testUser2()))

			users, err := store.GetUsersByIDs(ctx, []string{user.ID})

			require.NoError(t, err)
			require.Len(t, users, 1)
			require.Equal(t, user, users[0])
		})
	})
}

func TestUserStore_GetUserByUsername(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.UserStore) {
		t.Run("not-found", func(t *testing.T) {
			_, err := store.GetUserByUsername(ctx, "missing")

			require.IsType(t, errorz.NotFoundError{}, err)
		})

		t.Run("success", func(t *testing.T) {
			user := testUser()

			require.NoError(t, store.CreateUser(ctx, user))

			createdUser, err := store.GetUserByUsername(ctx, user.Username)

			require.NoError(t, err)
			require.Equal(t, user, createdUser)
		})
	})
}
