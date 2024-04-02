package access

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrix_GetPermissions(t *testing.T) {
	t.Parallel()

	m := NewMatrix()

	t.Run("none", func(t *testing.T) {
		require.Equal(t, PermissionsNone, m.GetPermissions(RoleNone))
	})

	t.Run("others", func(t *testing.T) {
		const skipNone = 1

		for _, r := range AllRoles[skipNone:] {
			t.Run(r.String(), func(t *testing.T) {
				require.NotEqual(t, PermissionsNone, m.GetPermissions(r))
			})
		}
	})
}

func TestMatrix_HasPermissions(t *testing.T) {
	t.Parallel()

	m := NewMatrix()

	require.True(t, m.HasPermission(RoleAdmin, PermissionRead))
}
