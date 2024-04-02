package access

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermission_String(t *testing.T) {
	t.Parallel()

	for _, p := range AllPermissions {
		t.Run(p.String(), func(t *testing.T) {
			require.NotEmpty(t, p.String())
			require.False(t, p.String() == "Permission(")
		})
	}

	t.Run("unknown", func(t *testing.T) {
		p := Permission(100)

		require.Equal(t, "Permission(100)", p.String())
	})
}
