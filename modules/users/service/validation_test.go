package service

import (
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/users"
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

func Test_validateExternalID(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateExternalID("1"))
	require.Error(t, validateExternalID(""))
}

func Test_validateUsername(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateUsername("user"))
	require.Error(t, validateUsername(""))
}

func Test_validateEmail(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateEmail("user@somewhere.com"))
	require.Error(t, validateEmail(""))
	require.Error(t, validateEmail("invalid"))
}

func Test_validateUserData(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		data    users.UserData
		wantErr bool
	}{
		"valid": {
			data: users.UserData{
				Username: "user",
				Email:    "user@somewhere.com",
			},
			wantErr: false,
		},
		"invalid-username": {
			data: users.UserData{
				Username: "",
				Email:    "user@somewhere.com",
			},
			wantErr: true,
		},
		"invalid-email": {
			data: users.UserData{
				Username: "user",
				Email:    "",
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateUserData(test.data)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
