package mongo_test

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/permissions"
	"github.com/energimind/powermesh-core/modules/permissions/store/mongo"
	"github.com/stretchr/testify/require"
)

func TestPermissionStore_CreateRoleBinding(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.PermissionStore) {
		roleBinding := testRoleBinding()

		require.NoError(t, store.CreateRoleBinding(ctx, roleBinding))

		createdRoleBinding, err := store.GetRoleBinding(ctx, roleBinding.ID)

		require.NoError(t, err)
		require.Equal(t, roleBinding, createdRoleBinding)
	})
}

func TestPermissionStore_UpdateRoleBinding(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.PermissionStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.UpdateRoleBinding(ctx, testRoleBinding()))
		})

		t.Run("success", func(t *testing.T) {
			roleBinding := testRoleBinding()

			require.NoError(t, store.CreateRoleBinding(ctx, roleBinding))

			roleBinding.Role = access.RoleGuest

			require.NoError(t, store.UpdateRoleBinding(ctx, roleBinding))

			updatedRoleBinding, err := store.GetRoleBinding(ctx, roleBinding.ID)

			require.NoError(t, err)
			require.Equal(t, roleBinding, updatedRoleBinding)
		})
	})
}

func TestPermissionStore_DeleteRoleBinding(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.PermissionStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.DeleteRoleBinding(ctx, "missing"))
		})

		t.Run("success", func(t *testing.T) {
			roleBinding := testRoleBinding()

			require.NoError(t, store.CreateRoleBinding(ctx, roleBinding))

			require.NoError(t, store.DeleteRoleBinding(ctx, roleBinding.ID))

			_, err := store.GetRoleBinding(ctx, roleBinding.ID)
			require.Error(t, err)
		})
	})
}

func TestPermissionStore_GetRoleBinding(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.PermissionStore) {
		t.Run("not-found", func(t *testing.T) {
			_, err := store.GetRoleBinding(ctx, "missing")

			require.IsType(t, errorz.NotFoundError{}, err)
		})

		t.Run("success", func(t *testing.T) {
			roleBinding := testRoleBinding()

			require.NoError(t, store.CreateRoleBinding(ctx, roleBinding))

			createdRoleBinding, err := store.GetRoleBinding(ctx, roleBinding.ID)

			require.NoError(t, err)
			require.Equal(t, roleBinding, createdRoleBinding)
		})
	})
}

func TestPermissionStore_GetRoleBindingsByOwner(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.PermissionStore) {
		t.Run("not-found", func(t *testing.T) {
			roleBindings, err := store.GetRoleBindingsByOwner(ctx, "missing")

			require.NoError(t, err)
			require.Empty(t, roleBindings)
		})

		t.Run("success", func(t *testing.T) {
			roleBinding := testRoleBinding()

			require.NoError(t, store.CreateRoleBinding(ctx, roleBinding))
			require.NoError(t, store.CreateRoleBinding(ctx, testRoleBinding2()))

			roleBindings, err := store.GetRoleBindingsByOwner(ctx, roleBinding.OwnerID)

			require.NoError(t, err)
			require.Len(t, roleBindings, 1)
			require.Equal(t, roleBinding, roleBindings[0])
		})
	})
}

func TestPermissionStore_GetAccessibleResources(t *testing.T) {
	t.Parallel()

	withStore(t, func(t *testing.T, ctx context.Context, store *mongo.PermissionStore) {
		t.Run("not-found", func(t *testing.T) {
			ids, err := store.GetAccessibleResources(ctx, permissions.AccessibleResourcesQuery{})

			require.NoError(t, err)
			require.Empty(t, ids)
		})

		t.Run("success", func(t *testing.T) {
			roleBinding := testRoleBinding()

			require.NoError(t, store.CreateRoleBinding(ctx, roleBinding))
			require.NoError(t, store.CreateRoleBinding(ctx, testRoleBinding2()))

			ids, err := store.GetAccessibleResources(ctx, permissions.AccessibleResourcesQuery{
				UserID:       roleBinding.UserID,
				ResourceType: roleBinding.ResourceType,
			})

			require.NoError(t, err)
			require.Len(t, ids, 1)
			require.Equal(t, roleBinding.ResourceID, ids[0])
		})
	})
}
