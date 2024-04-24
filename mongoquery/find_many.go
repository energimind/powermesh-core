package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	coll       collection
	mapper     mapper[D, T]
	projection bson.D
	drain      func(ctx context.Context, cursor cursor) ([]D, error)
}

// Exec executes the query.
// It retrieves the documents from the collection.
// It returns an error if the operation failed.
func (q FindManyQuery[D, T]) Exec(ctx context.Context, filter Filter) ([]T, error) {
	buildOptions := func() *options.FindOptions {
		if len(q.projection) == 0 {
			return nil
		}

		return options.Find().SetProjection(q.projection)
	}

	cur, err := q.coll.Find(ctx, filter.toBSON(), buildOptions())
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

// WithProjection sets the projections for the query.
func (q FindManyQuery[D, T]) WithProjection(fields ...string) FindManyQuery[D, T] {
	for _, field := range fields {
		q.projection = append(q.projection, bson.E{Key: field, Value: 1})
	}

	return q
}
