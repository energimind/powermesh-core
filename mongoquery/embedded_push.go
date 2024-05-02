package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/bson"
)

// EmbeddedPush creates a new EmbeddedPushQuery.
func EmbeddedPush[D, T any](coll collection, field string, mapper mapper[T, D]) EmbeddedPushQuery[D, T] {
	return EmbeddedPushQuery[D, T]{
		coll:   coll,
		field:  field,
		mapper: mapper,
	}
}

// EmbeddedPushQuery pushes a single embedded document to the collection item.
type EmbeddedPushQuery[D, T any] struct {
	coll   collection
	field  string
	mapper mapper[T, D]
	key    string
}

// Key sets the key to use for the query.
// It returns the query itself.
func (q EmbeddedPushQuery[D, T]) Key(key string) EmbeddedPushQuery[D, T] {
	q.key = key

	return q
}

// Exec executes the query.
// It pushes the embedded document to the collection item.
// It returns an error if the operation failed.
func (q EmbeddedPushQuery[D, T]) Exec(ctx context.Context, id any, value T) error {
	qValue := q.mapper(value)
	qFilter := buildFilter(q.key, id)
	qUpdate := bson.M{
		"$push": bson.M{
			q.field: qValue,
		},
	}

	res, err := q.coll.UpdateOne(ctx, qFilter, qUpdate)
	if err != nil {
		return errorz.NewStoreError("failed to push %s: %v", singular(q.coll.Name()), err)
	}

	if res.MatchedCount == 0 {
		return errorz.NewNotFoundError("%s %v not found", singular(q.coll.Name()), id)
	}

	return nil
}
