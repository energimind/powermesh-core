package access

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRole_Has(t *testing.T) {
	t.Parallel()

	for _, r := range AllRoles {
		t.Run(r.String(), func(t *testing.T) {
			require.True(t, r.Has(r))
		})
	}
}

func TestRole_String(t *testing.T) {
	t.Parallel()

	for _, r := range AllRoles {
		t.Run(r.String(), func(t *testing.T) {
			require.NotEmpty(t, r.String())
			require.False(t, strings.HasPrefix(r.String(), "Role("))
		})
	}

	t.Run("unknown", func(t *testing.T) {
		r := Role(100)

		require.Equal(t, "Role(100)", r.String())
	})
}

func TestIsSupportedRole(t *testing.T) {
	t.Parallel()

	t.Run("supported", func(t *testing.T) {
		for _, r := range AllRoles {
			t.Run(r.String(), func(t *testing.T) {
				require.True(t, IsSupportedRole(r))
			})
		}
	})

	t.Run("unknown", func(t *testing.T) {
		r := Role(100)

		require.False(t, IsSupportedRole(r))
	})
}
