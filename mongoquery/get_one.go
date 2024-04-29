package mongoquery

import (
	"context"
	"errors"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOne creates a new GetOneQuery.
func GetOne[D, T any](coll collection, mapper mapper[D, T]) GetOneQuery[D, T] {
	return GetOneQuery[D, T]{
		coll:   coll,
		mapper: mapper,
	}
}

// GetOneQuery retrieves a single document from the collection.
type GetOneQuery[D, T any] struct {
	coll   collection
	mapper mapper[D, T]
	key    string
}

// Key sets the key to use for the query.
// It returns the query itself.
func (q GetOneQuery[D, T]) Key(key string) GetOneQuery[D, T] {
	q.key = key

	return q
}

// Exec executes the query.
// It retrieves the document from the collection.
// It accepts an ID or a filter as input.
// It returns an error if the operation failed.
func (q GetOneQuery[D, T]) Exec(ctx context.Context, idOrFilter any) (T, error) { //nolint:ireturn
	qFilter := buildFilter(q.key, idOrFilter)

	var qValue D

	if err := q.coll.FindOne(ctx, qFilter).Decode(&qValue); err != nil {
		var zero T

		if errors.Is(err, mongo.ErrNoDocuments) {
			return zero, errorz.NewNotFoundError("%s %v not found", singular(q.coll.Name()), idOrFilter)
		}

		return zero, errorz.NewStoreError("failed to get %s: %v", singular(q.coll.Name()), err)
	}

	return q.mapper(qValue), nil
}
