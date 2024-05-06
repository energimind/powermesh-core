package mongo_test

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/models/store/mongo"
	"github.com/stretchr/testify/require"
)

func TestModelStore_CreateModel(t *testing.T) {
	t.Parallel()

	withModelStore(t, func(t *testing.T, ctx context.Context, store *mongo.ModelStore) {
		require.NoError(t, store.CreateModel(ctx, testModel()))
	})
}

func TestModelStore_UpdateModel(t *testing.T) {
	t.Parallel()

	withModelStore(t, func(t *testing.T, ctx context.Context, store *mongo.ModelStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.UpdateModel(ctx, testModel()))
		})

		t.Run("success", func(t *testing.T) {
			model := testModel()

			require.NoError(t, store.CreateModel(ctx, model))

			model.Name = "new-name"

			require.NoError(t, store.UpdateModel(ctx, model))

			updatedModel, err := store.GetModel(ctx, model.ID)

			require.NoError(t, err)
			require.Equal(t, model, updatedModel)
		})
	})
}

func TestModelStore_DeleteModel(t *testing.T) {
	t.Parallel()

	withModelStore(t, func(t *testing.T, ctx context.Context, store *mongo.ModelStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.DeleteModel(ctx, "missing"))
		})

		t.Run("success", func(t *testing.T) {
			model := testModel()

			require.NoError(t, store.CreateModel(ctx, model))

			require.NoError(t, store.DeleteModel(ctx, model.ID))

			_, err := store.GetModel(ctx, model.ID)

			require.IsType(t, errorz.NotFoundError{}, err)
		})
	})
}

func TestModelStore_GetModel(t *testing.T) {
	t.Parallel()

	withModelStore(t, func(t *testing.T, ctx context.Context, store *mongo.ModelStore) {
		t.Run("not-found", func(t *testing.T) {
			_, err := store.GetModel(ctx, "missing")

			require.IsType(t, errorz.NotFoundError{}, err)
		})

		t.Run("success", func(t *testing.T) {
			model := testModel()

			require.NoError(t, store.CreateModel(ctx, model))

			foundModel, err := store.GetModel(ctx, model.ID)

			require.NoError(t, err)
			require.Equal(t, model, foundModel)
		})
	})
}

func TestModelStore_GetModelsByIDs(t *testing.T) {
	t.Parallel()

	t.Run("not-found", func(t *testing.T) {
		withModelStore(t, func(t *testing.T, ctx context.Context, store *mongo.ModelStore) {
			models, err := store.GetModelsByIDs(ctx, []string{"non-existent"})

			require.NoError(t, err)
			require.Len(t, models, 0)
		})
	})

	t.Run("success", func(t *testing.T) {
		withModelStore(t, func(t *testing.T, ctx context.Context, store *mongo.ModelStore) {
			model1 := testModel()
			model2 := testModel2()

			require.NoError(t, store.CreateModel(ctx, model1))
			require.NoError(t, store.CreateModel(ctx, model2))

			models, err := store.GetModelsByIDs(ctx, []string{model1.ID, model2.ID})

			require.NoError(t, err)
			require.Len(t, models, 2)
			require.Contains(t, models, model1)
			require.Contains(t, models, model2)
		})
	})
}
