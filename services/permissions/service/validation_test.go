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

func Test_validateObjectID(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		objectID string
		wantErr  bool
	}{
		"empty": {
			objectID: "",
			wantErr:  true,
		},
		"valid": {
			objectID: "1",
			wantErr:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateObjectID(test.objectID)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateObjectType(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		objectType permissions.ObjectType
		wantErr    bool
	}{
		"empty": {
			objectType: permissions.ObjectType(100),
			wantErr:    true,
		},
		"valid": {
			objectType: permissions.ObjectTypeModel,
			wantErr:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateObjectType(test.objectType)

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
				OwnerID:    "",
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidUserID": {
			data: permissions.RoleBindingData{
				OwnerID:    validOwnerID,
				UserID:     "",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidObjectID": {
			data: permissions.RoleBindingData{
				OwnerID:    validOwnerID,
				UserID:     "1",
				ObjectID:   "",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidObjectType": {
			data: permissions.RoleBindingData{
				OwnerID:    validOwnerID,
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectType(100),
				Role:       access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidRole": {
			data: permissions.RoleBindingData{
				OwnerID:    validOwnerID,
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.Role(100),
			},
			wantErr: true,
		},
		"valid": {
			data: permissions.RoleBindingData{
				OwnerID:    validOwnerID,
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.RoleAdmin,
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
				UserID:     "",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
			},
			wantErr: true,
		},
		"invalidObjectID": {
			query: permissions.RoleBindingQuery{
				UserID:     "1",
				ObjectID:   "",
				ObjectType: permissions.ObjectTypeModel,
			},
			wantErr: true,
		},
		"invalidObjectType": {
			query: permissions.RoleBindingQuery{
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectType(100),
			},
			wantErr: true,
		},
		"valid": {
			query: permissions.RoleBindingQuery{
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
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

func Test_validateAccessibleObjectsQuery(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		query   permissions.AccessibleObjectsQuery
		wantErr bool
	}{
		"invalidUserID": {
			query: permissions.AccessibleObjectsQuery{
				UserID:     "",
				ObjectType: permissions.ObjectTypeModel,
			},
			wantErr: true,
		},
		"invalidObjectType": {
			query: permissions.AccessibleObjectsQuery{
				UserID:     "1",
				ObjectType: permissions.ObjectType(100),
			},
			wantErr: true,
		},
		"valid": {
			query: permissions.AccessibleObjectsQuery{
				UserID:     "1",
				ObjectType: permissions.ObjectTypeModel,
			},
			wantErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateAccessibleObjectsQuery(test.query)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
