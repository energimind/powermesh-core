package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/bson"
)

// EmbeddedPull creates a new EmbeddedPullQuery.
func EmbeddedPull(coll collection, field, subDocKey string) EmbeddedPullQuery {
	return EmbeddedPullQuery{
		coll:      coll,
		field:     field,
		subDocKey: subDocKey,
	}
}

// EmbeddedPullQuery deletes a single embedded document from the collection item.
type EmbeddedPullQuery struct {
	coll      collection
	field     string
	subDocKey string
	key       string
}

// Key sets the key to use for the query.
// It returns the query itself.
func (q EmbeddedPullQuery) Key(key string) EmbeddedPullQuery {
	q.key = key

	return q
}

// Exec executes the query.
// It deletes the embedded document from the collection item.
// It returns an error if the operation failed.
func (q EmbeddedPullQuery) Exec(ctx context.Context, id any, subDocID any) error {
	qFilter := buildFilter(q.key, id)
	qUpdate := bson.M{
		"$pull": bson.M{
			q.field: bson.M{
				q.subDocKey: subDocID,
			},
		},
	}

	res, err := q.coll.UpdateOne(ctx, qFilter, qUpdate)
	if err != nil {
		return errorz.NewStoreError("failed to pull %s: %v", singular(q.coll.Name()), err)
	}

	if res.MatchedCount == 0 {
		return errorz.NewNotFoundError("%s %v not found", singular(q.coll.Name()), id)
	}

	if res.ModifiedCount == 0 {
		return errorz.NewNotFoundError("field %s[%s] not found in %s %v",
			q.field, subDocID, singular(q.coll.Name()), id)
	}

	return nil
}
