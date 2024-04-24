package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
)

// DeleteOne creates a new DeleteOneQuery.
func DeleteOne(coll collection) DeleteOneQuery {
	return DeleteOneQuery{
		coll: coll,
	}
}

// DeleteOneQuery deletes a single document from the collection.
type DeleteOneQuery struct {
	coll collection
}

// Exec executes the query.
// It deletes the document from the collection.
// It accepts an ID or a filter as input.
// It returns an error if the operation failed.
func (q DeleteOneQuery) Exec(ctx context.Context, idOrFilter any) error {
	qFilter := buildFilter(idOrFilter)

	res, err := q.coll.DeleteOne(ctx, qFilter)
	if err != nil {
		return errorz.NewStoreError("failed to delete %s: %v", singular(q.coll.Name()), err)
	}

	if res.DeletedCount == 0 {
		return errorz.NewNotFoundError("%s %v not found", singular(q.coll.Name()), idOrFilter)
	}

	return nil
}
