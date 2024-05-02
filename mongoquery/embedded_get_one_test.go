package mongoquery

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_EmbeddedGetOne(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			findOne: func() *mongo.SingleResult {
				return mongo.NewSingleResultFromDocument(testDBPerson, nil, nil)
			},
		}

		rsp, err := EmbeddedGetOne(coll, "address", "id", extractFirstAddress).Key("id").
			Exec(context.Background(), testID, testAddressID)

		require.NoError(t, err)
		require.Equal(t, testDomainAddress, rsp)
	})

	t.Run("not-found", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			findOne: func() *mongo.SingleResult {
				return mongo.NewSingleResultFromDocument(bson.M{}, mongo.ErrNoDocuments, nil)
			},
		}

		rsp, err := EmbeddedGetOne(coll, "address", "id", extractFirstAddress).
			Exec(context.Background(), testID, testAddressID)

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

		rsp, err := EmbeddedGetOne(coll, "address", "id", extractFirstAddress).
			Exec(context.Background(), testID, testAddressID)

		require.IsType(t, errorz.StoreError{}, err)
		require.ErrorContains(t, err, "forced error")
		require.Zero(t, rsp)
	})
}
