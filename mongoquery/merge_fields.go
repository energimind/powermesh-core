package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
	"go.mongodb.org/mongo-driver/bson"
)

// MergeFields creates a new query to update one or more fields in a document.
func MergeFields(coll collection) MergeFieldsQuery {
	return MergeFieldsQuery{
		coll: coll,
	}
}

// MergeFieldsQuery is a query to update one or more fields in a document.
type MergeFieldsQuery struct {
	coll collection
}

// Exec executes the query.
// It updates the fields in the document.
// It returns an error if the operation failed.
func (q MergeFieldsQuery) Exec(ctx context.Context, id any, fields map[string]any) error {
	qFilter := bson.M{"id": id}
	qUpdate := bson.M{"$set": bson.M(fields)}

	res, err := q.coll.UpdateOne(ctx, qFilter, qUpdate)
	if err != nil {
		return errorz.NewStoreError("failed to update %s: %v", singular(q.coll.Name()), err)
	}

	if res.MatchedCount == 0 {
		return errorz.NewNotFoundError("%s %v not found", singular(q.coll.Name()), id)
	}

	return nil
}
