package mongoquery

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestUpdateOne(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "UpdateOne",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 1}, nil
			},
		}

		require.NoError(t, UpdateOne(coll, toDBPerson).Key("id").Exec(context.Background(), testID, testDomainPerson))
	})

	t.Run("not-found", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "UpdateOne",
			updateOne: func() (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{MatchedCount: 0}, nil
			},
		}

		require.ErrorContains(t,
			UpdateOne(coll, toDBPerson).Exec(context.Background(), testID, testDomainPerson),
			"person 1 not found")
	})

	t.Run("update-error", func(t *testing.T) {
		coll := &mockCollection{
			t:      t,
			caller: "UpdateOne",
			updateOne: func() (*mongo.UpdateResult, error) {
				return nil, forcedError{}
			},
		}

		require.ErrorContains(t,
			UpdateOne(coll, toDBPerson).Exec(context.Background(), testID, testDomainPerson),
			"forced error")
	})
}
