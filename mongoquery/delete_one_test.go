package mongoquery

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteOne(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			deleteOne: func() (*mongo.DeleteResult, error) {
				return &mongo.DeleteResult{DeletedCount: 1}, nil
			},
		}

		require.NoError(t, DeleteOne(coll).Exec(context.Background(), testID))
	})

	t.Run("not-found", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			deleteOne: func() (*mongo.DeleteResult, error) {
				return &mongo.DeleteResult{DeletedCount: 0}, nil
			},
		}

		require.ErrorContains(t,
			DeleteOne(coll).Exec(context.Background(), testID),
			"not found")
	})

	t.Run("delete-error", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			deleteOne: func() (*mongo.DeleteResult, error) {
				return nil, forcedError{}
			},
		}

		require.ErrorContains(t,
			DeleteOne(coll).Exec(context.Background(), testID),
			"forced error")
	})
}
