package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/bson"
)

// EmbeddedUpdate creates a new EmbeddedUpdateQuery.
func EmbeddedUpdate[D, T any](coll collection, field, subDocKey string, mapper mapper[T, D]) EmbeddedUpdateQuery[D, T] {
	return EmbeddedUpdateQuery[D, T]{
		coll:      coll,
		field:     field,
		subDocKey: subDocKey,
		mapper:    mapper,
	}
}

// EmbeddedUpdateQuery updates a single embedded document in the collection item.
type EmbeddedUpdateQuery[D, T any] struct {
	coll      collection
	field     string
	subDocKey string
	mapper    mapper[T, D]
	key       string
}

// Key sets the key to use for the query.
// It returns the query itself.
func (q EmbeddedUpdateQuery[D, T]) Key(key string) EmbeddedUpdateQuery[D, T] {
	q.key = key

	return q
}

// Exec executes the query.
// It updates the embedded document in the collection item.
// It returns an error if the operation failed.
func (q EmbeddedUpdateQuery[D, T]) Exec(ctx context.Context, id, subDocID any, fieldValue T) error {
	qValue := q.mapper(fieldValue)
	qFilter := buildFilter(q.key, id)

	qFilter[q.field+"."+q.subDocKey] = subDocID

	qUpdate := bson.M{
		"$set": bson.M{
			q.field + ".$": qValue,
		},
	}

	res, err := q.coll.UpdateOne(ctx, qFilter, qUpdate)
	if err != nil {
		return errorz.NewStoreError("failed to update %s: %v", singular(q.coll.Name()), err)
	}

	if res.MatchedCount == 0 {
		return errorz.NewNotFoundError("%s[%s] %v[%v] not found", singular(q.coll.Name()), q.field, id, subDocID)
	}

	return nil
}
