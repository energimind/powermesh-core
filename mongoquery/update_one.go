package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateOne creates a new UpdateOneQuery.
func UpdateOne[D, T any](coll collection, mapper mapper[T, D]) UpdateOneQuery[D, T] {
	return UpdateOneQuery[D, T]{
		coll:   coll,
		mapper: mapper,
	}
}

// UpdateOneQuery updates a single document in the collection.
type UpdateOneQuery[D, T any] struct {
	coll   collection
	mapper mapper[T, D]
}

// Exec executes the query.
// It updates the document in the collection.
// It returns an error if the operation failed.
func (q UpdateOneQuery[D, T]) Exec(ctx context.Context, id any, value T) error {
	qValue := q.mapper(value)
	qFilter := bson.M{"id": id}
	qUpdate := bson.M{"$set": qValue}

	res, err := q.coll.UpdateOne(ctx, qFilter, qUpdate)
	if err != nil {
		return errorz.NewStoreError("failed to update %s: %v", singular(q.coll.Name()), err)
	}

	if res.MatchedCount == 0 {
		return errorz.NewNotFoundError("%s %v not found", singular(q.coll.Name()), id)
	}

	return nil
}
