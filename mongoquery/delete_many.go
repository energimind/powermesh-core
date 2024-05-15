package mongoquery

import (
	"context"

	"github.com/energimind/powermesh-core/errorz"
)

// DeleteMany creates a new DeleteManyQuery.
func DeleteMany(coll collection) DeleteManyQuery {
	return DeleteManyQuery{
		coll: coll,
	}
}

// DeleteManyQuery deletes multiple documents from the collection.
type DeleteManyQuery struct {
	coll collection
}

// Exec executes the query.
// It deletes the documents from the collection.
// It returns the number of deleted documents and an error if the operation failed.
func (q DeleteManyQuery) Exec(ctx context.Context, filter Filter) (int64, error) {
	res, err := q.coll.DeleteMany(ctx, filter.toBSON())
	if err != nil {
		return 0, errorz.NewStoreError("failed to delete %s: %v", singular(q.coll.Name()), err)
	}

	return res.DeletedCount, nil
}
