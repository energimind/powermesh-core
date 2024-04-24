package mongoquery

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFindMany(t *testing.T) {
	t.Parallel()

	findResult := []any{testDBPerson}
	filter := Filter{}.GT("age", 20)

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			find: func() (*mongo.Cursor, error) {
				documents, err := mongo.NewCursorFromDocuments(findResult, nil, nil)

				return documents, err
			},
		}

		rsp, err := FindMany(coll, fromDBPerson).Exec(context.Background(), filter)

		require.NoError(t, err)
		require.NotNil(t, rsp)
	})

	t.Run("find-error", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			find: func() (*mongo.Cursor, error) {
				return nil, forcedError{}
			},
		}

		rsp, err := FindMany(coll, fromDBPerson).Exec(context.Background(), filter)

		require.ErrorContains(t, err, "forced error")
		require.ErrorContains(t, err, "failed to find many")
		require.Nil(t, rsp)
	})

	t.Run("drain-error", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			find: func() (*mongo.Cursor, error) {
				documents, err := mongo.NewCursorFromDocuments(findResult, nil, nil)

				return documents, err
			},
		}

		findMany := FindMany(coll, fromDBPerson)

		findMany.drain = func(ctx context.Context, cursor cursor) ([]dbPerson, error) {
			return nil, forcedError{}
		}

		rsp, err := findMany.Exec(context.Background(), filter)

		require.ErrorContains(t, err, "failed to drain many")
		require.Nil(t, rsp)
	})
}
