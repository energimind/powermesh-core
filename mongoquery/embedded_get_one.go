package mongoquery

import (
	"context"
	"errors"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EmbeddedGetOne creates a new EmbeddedGetOneQuery.
func EmbeddedGetOne[D, T any](coll collection, field, subDocKey string, mapper mapper[D, T]) EmbeddedGetOneQuery[D, T] {
	return EmbeddedGetOneQuery[D, T]{
		coll:      coll,
		field:     field,
		subDocKey: subDocKey,
		mapper:    mapper,
	}
}

// EmbeddedGetOneQuery retrieves a single embedded document from the collection item.
type EmbeddedGetOneQuery[D, T any] struct {
	coll      collection
	field     string
	subDocKey string
	mapper    mapper[D, T]
	key       string
}

// Key sets the key to use for the query.
// It returns the query itself.
func (q EmbeddedGetOneQuery[D, T]) Key(key string) EmbeddedGetOneQuery[D, T] {
	q.key = key

	return q
}

// Exec executes the query.
// It retrieves the single document from the collection item.
// It returns an error if the operation failed.
func (q EmbeddedGetOneQuery[D, T]) Exec(ctx context.Context, id, subDocID any) (T, error) { //nolint:ireturn
	qFilter := buildFilter(q.key, id)
	qProjection := bson.M{q.field + ".$": 1} // return the matched sub-document only

	qFilter[q.field+"."+q.subDocKey] = subDocID

	opts := options.FindOne().SetProjection(qProjection)

	var qValue D

	if err := q.coll.FindOne(ctx, qFilter, opts).Decode(&qValue); err != nil {
		var zero T

		if errors.Is(err, mongo.ErrNoDocuments) {
			return zero, errorz.NewNotFoundError("field %s[%s] not found in %s %v",
				q.field, q.subDocKey, singular(q.coll.Name()), id)
		}

		return zero, errorz.NewStoreError("failed to get %s: %v", singular(q.coll.Name()), err)
	}

	return q.mapper(qValue), nil
}
