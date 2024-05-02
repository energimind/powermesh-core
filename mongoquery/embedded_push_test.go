package mongoquery

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_EmbeddedPush(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedPush",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 1}, nil
			},
		}

		err := EmbeddedPush(coll, "address", toDBAddress).Key("id").
			Exec(context.Background(), testID, testDomainAddress)

		require.NoError(t, err)
	})

	t.Run("not-found", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedPush",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 0}, nil
			},
		}

		err := EmbeddedPush(coll, "address", toDBAddress).
			Exec(context.Background(), testID, testDomainAddress)

		require.IsType(t, errorz.NotFoundError{}, err)
		require.ErrorContains(t, err, "person 1 not found")
	})

	t.Run("update-error", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "EmbeddedPush",
			updateOne: func() (*mongo.UpdateResult, error) {
				return nil, forcedError{}
			},
		}

		err := EmbeddedPush(coll, "address", toDBAddress).
			Exec(context.Background(), testID, testDomainAddress)

		require.IsType(t, errorz.StoreError{}, err)
		require.ErrorContains(t, err, "forced error")
	})
}
