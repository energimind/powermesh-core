package mongoquery

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteMany(t *testing.T) {
	t.Parallel()

	filter := Filter{}.GT("age", 20)

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			deleteMany: func() (*mongo.DeleteResult, error) {
				return &mongo.DeleteResult{DeletedCount: 2}, nil
			},
		}

		count, err := DeleteMany(coll).Exec(context.Background(), filter)

		require.NoError(t, err)
		require.Equal(t, int64(2), count)
	})

	t.Run("not-found", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			deleteMany: func() (*mongo.DeleteResult, error) {
				return &mongo.DeleteResult{DeletedCount: 0}, nil
			},
		}

		count, err := DeleteMany(coll).Exec(context.Background(), filter)

		require.NoError(t, err)
		require.Equal(t, int64(0), count)
	})

	t.Run("delete-error", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			deleteMany: func() (*mongo.DeleteResult, error) {
				return nil, forcedError{}
			},
		}

		count, err := DeleteMany(coll).Exec(context.Background(), filter)

		require.ErrorContains(t, err, "forced error")
		require.Equal(t, int64(0), count)
	})
}
