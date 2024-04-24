package mongoquery

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateOne(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			insertOne: func() (*mongo.InsertOneResult, error) {
				return nil, nil
			},
		}

		require.NoError(t, CreateOne(coll, toDBPerson).Exec(context.Background(), testDomainPerson))
	})

	t.Run("insert-error", func(t *testing.T) {
		coll := &mockCollection{
			t: t,
			insertOne: func() (*mongo.InsertOneResult, error) {
				return nil, forcedError{}
			},
		}

		require.ErrorContains(t,
			CreateOne(coll, toDBPerson).Exec(context.Background(), testDomainPerson),
			"forced error")
	})
}
