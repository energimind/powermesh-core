package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
)

// CreateOne creates a new CreateOneQuery.
func CreateOne[D, T any](coll collection, mapper mapper[T, D]) CreateOneQuery[D, T] {
	return CreateOneQuery[D, T]{
		coll:   coll,
		mapper: mapper,
	}
}

// CreateOneQuery creates a new document in the collection.
type CreateOneQuery[D, T any] struct {
	coll   collection
	mapper mapper[T, D]
}

// Exec executes the query.
// It inserts the document into the collection.
// It returns an error if the operation failed.
func (q CreateOneQuery[D, T]) Exec(ctx context.Context, value T) error {
	qValue := q.mapper(value)

	if _, err := q.coll.InsertOne(ctx, qValue); err != nil {
		return errorz.NewStoreError("failed to create %s: %v", singular(q.coll.Name()), err)
	}

	return nil
}
