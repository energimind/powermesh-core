package mongoquery

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestEmbeddedPull(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedPull",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil
			},
		}

		err := EmbeddedPull(coll, "address", "id").Key("id").
			Exec(context.Background(), testID, testAddressID)

		require.NoError(t, err)
	})

	t.Run("parent-not-found", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedPull",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 0}, nil
			},
		}

		err := EmbeddedPull(coll, "address", "id").
			Exec(context.Background(), testID, testAddressID)

		require.IsType(t, errorz.NotFoundError{}, err)
		require.ErrorContains(t, err, "person 1 not found")
	})

	t.Run("embedded-not-found", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedPull",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 0}, nil
			},
		}

		err := EmbeddedPull(coll, "address", "id").
			Exec(context.Background(), testID, testAddressID)

		require.IsType(t, errorz.NotFoundError{}, err)
		require.ErrorContains(t, err, "field address[2] not found in person 1")
	})

	t.Run("update-error", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedPull",
			updateOne: func() (*mongo.UpdateResult, error) {
				return nil, forcedError{}
			},
		}

		err := EmbeddedPull(coll, "address", "id").
			Exec(context.Background(), testID, testAddressID)

		require.IsType(t, errorz.StoreError{}, err)
		require.ErrorContains(t, err, "forced error")
	})
}
