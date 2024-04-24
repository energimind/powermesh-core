package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
)

// FindMany creates a new FindManyQuery.
func FindMany[D, T any](coll collection, mapper mapper[D, T]) FindManyQuery[D, T] {
	return FindManyQuery[D, T]{
		coll:   coll,
		mapper: mapper,
		drain:  drainCursor[D],
	}
}

// FindManyQuery retrieves multiple documents from the collection.
type FindManyQuery[D, T any] struct {
	coll   collection
	mapper mapper[D, T]
	drain  func(ctx context.Context, cursor cursor) ([]D, error)
}

// Exec executes the query.
// It retrieves the documents from the collection.
// It returns an error if the operation failed.
func (q FindManyQuery[D, T]) Exec(ctx context.Context, filter Filter) ([]T, error) {
	cur, err := q.coll.Find(ctx, filter.toBSON())
	if err != nil {
		return nil, errorz.NewStoreError("failed to find many from %s: %v", q.coll.Name(), err)
	}

	qValues, err := q.drain(ctx, cur)
	if err != nil {
		return nil, errorz.NewStoreError("failed to drain many from %s: %v", q.coll.Name(), err)
	}

	values := make([]T, len(qValues))

	for i := range qValues {
		values[i] = q.mapper(qValues[i])
	}

	return values, nil
}
