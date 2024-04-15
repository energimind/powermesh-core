package permissions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestObjectType_String(t *testing.T) {
	t.Parallel()

	for _, r := range AllObjectTypes {
		t.Run(r.String(), func(t *testing.T) {
			require.NotEmpty(t, r.String())
			require.False(t, strings.HasPrefix(r.String(), "ObjectType("))
		})
	}

	t.Run("unknown", func(t *testing.T) {
		r := ObjectType(100)

		require.Equal(t, "ObjectType(100)", r.String())
	})
}

func TestIsSupportedObjectType(t *testing.T) {
	t.Parallel()

	t.Run("supported", func(t *testing.T) {
		for _, r := range AllObjectTypes {
			t.Run(r.String(), func(t *testing.T) {
				require.True(t, IsSupportedObjectType(r))
			})
		}
	})

	t.Run("unknown", func(t *testing.T) {
		r := ObjectType(100)

		require.False(t, IsSupportedObjectType(r))
	})
}
