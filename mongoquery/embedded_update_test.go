package mongoquery

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_EmbeddedUpdate(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedUpdate",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil
			},
		}

		err := EmbeddedUpdate(coll, "address", "id", toDBAddress).Key("id").
			Exec(context.Background(), testID, testAddressID, testDomainAddress)

		require.NoError(t, err)
	})

	t.Run("not-found", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedUpdate",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 0}, nil
			},
		}

		err := EmbeddedUpdate(coll, "address", "id", toDBAddress).
			Exec(context.Background(), testID, testAddressID, testDomainAddress)

		require.IsType(t, errorz.NotFoundError{}, err)
	})

	t.Run("update-error", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedUpdate",
			updateOne: func() (*mongo.UpdateResult, error) {
				return nil, forcedError{}
			},
		}

		err := EmbeddedUpdate(coll, "address", "id", toDBAddress).
			Exec(context.Background(), testID, testAddressID, testDomainAddress)

		require.IsType(t, errorz.StoreError{}, err)
		require.ErrorContains(t, err, "forced error")
	})
}
