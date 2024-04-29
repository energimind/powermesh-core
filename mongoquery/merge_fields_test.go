package mongoquery

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestMergeFields(t *testing.T) {
	t.Parallel()

	testFields := bson.M{"name": "John", "age": 30}

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "MergeFields",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 1}, nil
			},
		}

		require.NoError(t, MergeFields(coll).Exec(context.Background(), testID, testFields))
	})

	t.Run("not-found", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "MergeFields",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 0}, nil
			},
		}

		require.ErrorContains(t,
			MergeFields(coll).Exec(context.Background(), testID, testFields),
			"person 1 not found")
	})

	t.Run("update-error", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "MergeFields",
			updateOne: func() (*mongo.UpdateResult, error) {
				return nil, forcedError{}
			},
		}

		require.ErrorContains(t,
			MergeFields(coll).Exec(context.Background(), testID, testFields),
			"forced error")
	})
}
