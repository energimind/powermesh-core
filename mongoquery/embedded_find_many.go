package mongoquery

import (
	"context"
	"errors"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/mongo"
)

// EmbeddedFindMany creates a new EmbeddedFindManyQuery.
func EmbeddedFindMany[D, T any](coll collection, field string, mapper mapper[D, T]) EmbeddedFindManyQuery[D, T] {
	return EmbeddedFindManyQuery[D, T]{
		coll:   coll,
		field:  field,
		mapper: mapper,
	}
}

// EmbeddedFindManyQuery retrieves multiple embedded documents from the collection item.
type EmbeddedFindManyQuery[D, T any] struct {
	coll      collection
	field     string
	subDocKey string
	mapper    mapper[D, T]
	key       string
}

// Key sets the key to use for the query.
// It returns the query itself.
func (q EmbeddedFindManyQuery[D, T]) Key(key string) EmbeddedFindManyQuery[D, T] {
	q.key = key

	return q
}

// Exec executes the query.
// It retrieves the documents from the collection item.
// It returns an error if the operation failed.
func (q EmbeddedFindManyQuery[D, T]) Exec(ctx context.Context, id any) (T, error) { //nolint:ireturn
	qFilter := buildFilter(q.key, id)

	var qValue D

	if err := q.coll.FindOne(ctx, qFilter).Decode(&qValue); err != nil {
		var zero T

		if errors.Is(err, mongo.ErrNoDocuments) {
			return zero, errorz.NewNotFoundError("field %s[%s] not found in %s %v",
				q.field, q.subDocKey, singular(q.coll.Name()), id)
		}

		return zero, errorz.NewStoreError("failed to get %s: %v", singular(q.coll.Name()), err)
	}

	return q.mapper(qValue), nil
}
