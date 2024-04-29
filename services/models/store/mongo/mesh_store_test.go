package mongo_test

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/models/store/mongo"
	"github.com/stretchr/testify/require"
)

func TestMeshStore_CreateMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		require.NoError(t, store.CreateMesh(ctx, testMesh()))
	})
}

func TestMeshStore_UpdateMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.UpdateMesh(ctx, testMesh()))
		})

		t.Run("success", func(t *testing.T) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			mesh.Code = "new-code"

			require.NoError(t, store.UpdateMesh(ctx, mesh))

			updatedMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.Equal(t, mesh.Code, updatedMesh.Code)
		})
	})
}

func TestMeshStore_DeleteMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.DeleteMesh(ctx, "missing"))
		})

		t.Run("success", func(t *testing.T) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			require.NoError(t, store.DeleteMesh(ctx, mesh.ModelID))

			_, err := store.GetMesh(ctx, mesh.ModelID)

			require.IsType(t, errorz.NotFoundError{}, err)
		})
	})
}

func TestMeshStore_GetMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		t.Run("not-found", func(t *testing.T) {
			_, err := store.GetMesh(ctx, "missing")

			require.IsType(t, errorz.NotFoundError{}, err)
		})

		t.Run("success", func(t *testing.T) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			foundMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.Equal(t, mesh, foundMesh)
		})
	})
}
