package service

import (
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/services/permissions"
	"github.com/stretchr/testify/require"
)

func Test_validateID(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id      string
		wantErr bool
	}{
		"empty": {
			id:      "",
			wantErr: true,
		},
		"valid": {
			id:      "1",
			wantErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateID(test.id)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateOwnerID(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		ownerID string
		wantErr bool
	}{
		"empty": {
			ownerID: "",
			wantErr: true,
		},
		"valid": {
			ownerID: "1",
			wantErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateOwnerID(test.ownerID)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateResourceID(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		resourceID string
		wantErr    bool
	}{
		"empty": {
			resourceID: "",
			wantErr:    true,
		},
		"valid": {
			resourceID: "1",
			wantErr:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateResourceID(test.resourceID)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateResourceType(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		resourceType permissions.ResourceType
		wantErr      bool
	}{
		"empty": {
			resourceType: permissions.ResourceType(100),
			wantErr:      true,
		},
		"valid": {
			resourceType: permissions.ResourceTypeModel,
			wantErr:      false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateResourceType(test.resourceType)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateRole(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		role    access.Role
		wantErr bool
	}{
		"empty": {
			role:    access.Role(100),
			wantErr: true,
		},
		"valid": {
			role:    access.RoleAdmin,
			wantErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateRole(test.role)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
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
