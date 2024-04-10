package service

import (
	"testing"

	"github.com/energimind/powermesh-core/services/users"
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

func Test_validateUsername(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		username string
		wantErr  bool
	}{
		"valid": {
			username: "user",
			wantErr:  false,
		},
		"empty": {
			username: "",
			wantErr:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateUsername(test.username)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateEmail(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		email   string
		wantErr bool
	}{
		"valid": {
			email:   "user@somewhere.com",
			wantErr: false,
		},
		"empty": {
			email:   "",
			wantErr: true,
		},
		"invalid": {
			email:   "invalid",
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateEmail(test.email)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
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
