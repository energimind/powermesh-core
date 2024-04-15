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
		"valid": {
			id:      "1",
			wantErr: false,
		},
		"empty": {
			id:      "",
			wantErr: true,
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

func Test_validateObjectID(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		objectID string
		wantErr  bool
	}{
		"valid": {
			objectID: "1",
			wantErr:  false,
		},
		"empty": {
			objectID: "",
			wantErr:  true,
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
		"valid": {
			objectType: permissions.ObjectTypeModel,
			wantErr:    false,
		},
		"empty": {
			objectType: permissions.ObjectType(100),
			wantErr:    true,
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
		"valid": {
			role:    access.RoleAdmin,
			wantErr: false,
		},
		"empty": {
			role:    access.Role(100),
			wantErr: true,
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
		"valid": {
			data: permissions.RoleBindingData{
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.RoleAdmin,
			},
			wantErr: false,
		},
		"invalidUserID": {
			data: permissions.RoleBindingData{
				UserID:     "",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidObjectID": {
			data: permissions.RoleBindingData{
				UserID:     "1",
				ObjectID:   "",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidObjectType": {
			data: permissions.RoleBindingData{
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectType(100),
				Role:       access.RoleAdmin,
			},
			wantErr: true,
		},
		"invalidRole": {
			data: permissions.RoleBindingData{
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
				Role:       access.Role(100),
			},
			wantErr: true,
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
		"valid": {
			query: permissions.RoleBindingQuery{
				UserID:     "1",
				ObjectID:   "1",
				ObjectType: permissions.ObjectTypeModel,
			},
			wantErr: false,
		},
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
		"valid": {
			query: permissions.AccessibleObjectsQuery{
				UserID:     "1",
				ObjectType: permissions.ObjectTypeModel,
			},
			wantErr: false,
		},
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
