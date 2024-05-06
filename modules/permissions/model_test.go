package permissions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceType_String(t *testing.T) {
	t.Parallel()

	for _, r := range AllResourceTypes {
		t.Run(r.String(), func(t *testing.T) {
			require.NotEmpty(t, r.String())
			require.False(t, strings.HasPrefix(r.String(), "ResourceType("))
		})
	}

	t.Run("unknown", func(t *testing.T) {
		r := ResourceType(100)

		require.Equal(t, "ResourceType(100)", r.String())
	})
}

func TestIsSupportedResourceType(t *testing.T) {
	t.Parallel()

	t.Run("supported", func(t *testing.T) {
		for _, r := range AllResourceTypes {
			t.Run(r.String(), func(t *testing.T) {
				require.True(t, IsSupportedResourceType(r))
			})
		}
	})

	t.Run("unknown", func(t *testing.T) {
		r := ResourceType(100)

		require.False(t, IsSupportedResourceType(r))
	})
}
