package access

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermissions_Has(t *testing.T) {
	t.Parallel()

	for _, p := range AllPermissions {
		t.Run(p.String(), func(t *testing.T) {
			require.True(t, Permissions(p).Has(p))
			require.True(t, PermissionsAll.Has(p))
		})
	}
}

func TestPermissions_Add(t *testing.T) {
	t.Parallel()

	t.Run("none", func(t *testing.T) {
		require.Equal(t,
			PermissionsAll,
			PermissionsNone.Add(AllPermissions...))
	})

	t.Run("all", func(t *testing.T) {
		require.Equal(t,
			PermissionsAll,
			PermissionsAll.Add(AllPermissions...))
	})

	t.Run("single", func(t *testing.T) {
		require.Equal(t,
			Permissions(PermissionCreate),
			PermissionsNone.Add(PermissionCreate))
	})

	t.Run("multiple", func(t *testing.T) {
		require.Equal(t,
			Permissions(PermissionCreate|PermissionRead),
			PermissionsNone.Add(PermissionCreate, PermissionRead))
	})
}

func TestPermissions_Remove(t *testing.T) {
	t.Parallel()

	t.Run("none", func(t *testing.T) {
		require.Equal(t,
			PermissionsNone,
			PermissionsNone.Remove(AllPermissions...))
	})

	t.Run("all", func(t *testing.T) {
		require.Equal(t,
			PermissionsNone,
			PermissionsAll.Remove(AllPermissions...))
	})

	t.Run("single", func(t *testing.T) {
		require.Equal(t,
			Permissions(PermissionRead),
			Permissions(PermissionCreate|PermissionRead).Remove(PermissionCreate))
	})

	t.Run("multiple", func(t *testing.T) {
		require.Equal(t,
			Permissions(PermissionWrite),
			Permissions(PermissionCreate|PermissionRead|PermissionWrite).Remove(PermissionCreate, PermissionRead))
	})
}

func TestPermissions_String(t *testing.T) {
	t.Parallel()

	t.Run("none", func(t *testing.T) {
		require.Equal(t,
			"none",
			PermissionsNone.String())
	})

	t.Run("all", func(t *testing.T) {
		require.Equal(t,
			"all",
			PermissionsAll.String())
	})

	t.Run("single", func(t *testing.T) {
		require.Equal(t,
			"create",
			Permissions(PermissionCreate).String())
	})

	t.Run("multiple", func(t *testing.T) {
		require.Equal(t,
			"create,read",
			Permissions(PermissionCreate|PermissionRead).String())
	})
}
