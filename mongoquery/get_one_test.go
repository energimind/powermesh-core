package mongoquery

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetOne(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			findOne: func() *mongo.SingleResult {
				return mongo.NewSingleResultFromDocument(testDBPerson, nil, nil)
			},
		}

		rsp, err := GetOne(coll, fromDBPerson).Exec(context.Background(), testID)

		require.NoError(t, err)
		require.Equal(t, testDomainPerson, rsp)
	})

	t.Run("not-found", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			findOne: func() *mongo.SingleResult {
				return mongo.NewSingleResultFromDocument(bson.M{}, mongo.ErrNoDocuments, nil)
			},
		}

		rsp, err := GetOne(coll, fromDBPerson).Exec(context.Background(), testID)

		require.IsType(t, errorz.NotFoundError{}, err)
		require.Zero(t, rsp)
	})

	t.Run("find-error", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			findOne: func() *mongo.SingleResult {
				return mongo.NewSingleResultFromDocument(bson.M{}, forcedError{}, nil)
			},
		}

		rsp, err := GetOne(coll, fromDBPerson).Exec(context.Background(), testID)

		require.Error(t, err)
		require.Zero(t, rsp)
	})
}
