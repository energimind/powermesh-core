package mongoquery

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_drainCursor(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		c := &testCursor{
			all: func(v any) error {
				*v.(*[]string) = []string{"a", "b"}
				return nil
			},
			close: func() error {
				return nil
			},
		}

		s, err := drainCursor[string](context.Background(), c)
		require.NoError(t, err)

		require.Equal(t, []string{"a", "b"}, s)
	})

	t.Run("allError", func(t *testing.T) {
		c := &testCursor{
			all: func(v any) error {
				return errors.New("forcedError")
			},
			close: func() error {
				return nil
			},
		}

		_, err := drainCursor[string](context.Background(), c)
		require.ErrorContains(t, err, "forcedError")
	})

	t.Run("closeError", func(t *testing.T) {
		c := &testCursor{
			all: func(v any) error {
				return nil
			},
			close: func() error {
				return errors.New("forcedError")
			},
		}

		_, err := drainCursor[string](context.Background(), c)
		require.ErrorContains(t, err, "forcedError")
	})
}

type testCursor struct {
	all   func(v any) error
	close func() error
}

// ensure testCursor implements cursor interface
var _ cursor = (*testCursor)(nil)

func (m *testCursor) All(_ context.Context, v any) error {
	return m.all(v)
}

func (m *testCursor) Close(_ context.Context) error {
	return m.close()
}
