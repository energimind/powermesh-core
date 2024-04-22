package service

import (
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/permissions"
	"github.com/stretchr/testify/require"
)

func Test_requireString(t *testing.T) {
	t.Parallel()

	require.NoError(t, requireString("value", "name"))
	require.Error(t, requireString("", "name"))
	require.IsType(t, errorz.ValidationError{}, requireString("", "name"))
}

func Test_validateID(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateID("1"))
	require.Error(t, validateID(""))
}

func Test_validateOwnerID(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateOwnerID("1"))
	require.Error(t, validateOwnerID(""))
}

func Test_validateResourceID(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateResourceID("1"))
	require.Error(t, validateResourceID(""))
}

func Test_validateResourceType(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateResourceType(permissions.ResourceTypeModel))
	require.Error(t, validateResourceType(permissions.ResourceType(100)))
}

func Test_validateRole(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateRole(access.RoleAdmin))
	require.Error(t, validateRole(access.Role(100)))
}

func Test_validateRoleBindingData(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		data    permissions.RoleBindingData
		wantErr bool
	}{
		"invalidOwnerID": {
			data: permissions.RoleBindingData{
				OwnerID:      "",
				UserID:       "1",
				ResourceID:   "1",
				ResourceType: permissions.ResourceTypeModel,
				Role:         access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidUserID": {
			data: permissions.RoleBindingData{
				OwnerID:      validOwnerID,
				UserID:       "",
				ResourceID:   "1",
				ResourceType: permissions.ResourceTypeModel,
				Role:         access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidResourceID": {
			data: permissions.RoleBindingData{
				OwnerID:      validOwnerID,
				UserID:       "1",
				ResourceID:   "",
				ResourceType: permissions.ResourceTypeModel,
				Role:         access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidResourceType": {
			data: permissions.RoleBindingData{
				OwnerID:      validOwnerID,
				UserID:       "1",
				ResourceID:   "1",
				ResourceType: permissions.ResourceType(100),
				Role:         access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidRole": {
			data: permissions.RoleBindingData{
				OwnerID:      validOwnerID,
				UserID:       "1",
				ResourceID:   "1",
				ResourceType: permissions.ResourceTypeModel,
				Role:         access.Role(100),
			},
			wantErr: true,
		},
		"valid": {
			data: permissions.RoleBindingData{
				OwnerID:      validOwnerID,
				UserID:       "1",
				ResourceID:   "1",
				ResourceType: permissions.ResourceTypeModel,
				Role:         access.RoleAdmin,
			},
			wantErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateRoleBindingData(test.data)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateRoleBindingQuery(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		query   permissions.RoleBindingQuery
		wantErr bool
	}{
		"invalidUserID": {
			query: permissions.RoleBindingQuery{
				UserID:       "",
				ResourceID:   "1",
				ResourceType: permissions.ResourceTypeModel,
			},
			wantErr: true,
		},
		"invalidResourceID": {
			query: permissions.RoleBindingQuery{
				UserID:       "1",
				ResourceID:   "",
				ResourceType: permissions.ResourceTypeModel,
			},
			wantErr: true,
		},
		"invalidResourceType": {
			query: permissions.RoleBindingQuery{
				UserID:       "1",
				ResourceID:   "1",
				ResourceType: permissions.ResourceType(100),
			},
			wantErr: true,
		},
		"valid": {
			query: permissions.RoleBindingQuery{
				UserID:       "1",
				ResourceID:   "1",
				ResourceType: permissions.ResourceTypeModel,
			},
			wantErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateRoleBindingQuery(test.query)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateAccessibleResourcesQuery(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		query   permissions.AccessibleResourcesQuery
		wantErr bool
	}{
		"invalidUserID": {
			query: permissions.AccessibleResourcesQuery{
				UserID:       "",
				ResourceType: permissions.ResourceTypeModel,
			},
			wantErr: true,
		},
		"invalidResourceType": {
			query: permissions.AccessibleResourcesQuery{
				UserID:       "1",
				ResourceType: permissions.ResourceType(100),
			},
			wantErr: true,
		},
		"valid": {
			query: permissions.AccessibleResourcesQuery{
				UserID:       "1",
				ResourceType: permissions.ResourceTypeModel,
			},
			wantErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateAccessibleResourcesQuery(test.query)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
