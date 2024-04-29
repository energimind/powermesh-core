package mongoquery

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFilter(t *testing.T) {
	t.Run("EQ", func(t *testing.T) {
		f := Filter{}.EQ("key", "value")

		require.Equal(t, f, Filter{"key": "value"})
	})

	t.Run("NE", func(t *testing.T) {
		f := Filter{}.NE("key", "value")

		require.Equal(t, Filter{"key": bson.M{"$ne": "value"}}, f)
	})

	t.Run("GT", func(t *testing.T) {
		f := Filter{}.GT("key", 42)

		require.Equal(t, Filter{"key": bson.M{"$gt": 42}}, f)
	})

	t.Run("GTE", func(t *testing.T) {
		f := Filter{}.GTE("key", 42)

		require.Equal(t, Filter{"key": bson.M{"$gte": 42}}, f)
	})

	t.Run("LT", func(t *testing.T) {
		f := Filter{}.LT("key", 42)

		require.Equal(t, Filter{"key": bson.M{"$lt": 42}}, f)
	})

	t.Run("LTE", func(t *testing.T) {
		f := Filter{}.LTE("key", 42)

		require.Equal(t, Filter{"key": bson.M{"$lte": 42}}, f)
	})

	t.Run("IN", func(t *testing.T) {
		f := Filter{}.IN("key", []any{1, 2, 3})

		require.Equal(t, Filter{"key": bson.M{"$in": []any{1, 2, 3}}}, f)
	})
}

func TestFilter_toBSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := Filter{"key": "value"}

		require.Equal(t, bson.M{"key": "value"}, f.toBSON())
	})
}

func Test_buildFilter(t *testing.T) {
	t.Parallel()

	t.Run("filter", func(t *testing.T) {
		f := buildFilter("id", Filter{"key1": "value1", "key2": "value2"})

		require.Equal(t, bson.M{"key1": "value1", "key2": "value2"}, f)
	})

	t.Run("bson", func(t *testing.T) {
		f := buildFilter("id", bson.M{"key1": "value1", "key2": "value2"})

		require.Equal(t, bson.M{"key1": "value1", "key2": "value2"}, f)
	})

	t.Run("integer", func(t *testing.T) {
		f := buildFilter("id", 42)

		require.Equal(t, bson.M{"id": 42}, f)
	})

	t.Run("other", func(t *testing.T) {
		f := buildFilter("id", "someId")

		require.Equal(t, bson.M{"id": "someId"}, f)
	})
}
