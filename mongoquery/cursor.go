package mongoquery

import "context"

// cursor is an interface for a MongoDB cursor.
type cursor interface {
	All(ctx context.Context, v any) error
	Close(ctx context.Context) error
}

// drainCursor drains the cursor and returns the results.
// It also closes the cursor.
//
//nolint:wrapcheck
func drainCursor[T any](ctx context.Context, cursor cursor) ([]T, error) {
	const preallocate = 4

	results := make([]T, 0, preallocate)

	if err := cursor.All(ctx, &results); err != nil {
		_ = cursor.Close(ctx)

		return nil, err
	}

	if err := cursor.Close(ctx); err != nil {
		return nil, err
	}

	return results, nil
}
