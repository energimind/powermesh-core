package service

import (
	"testing"

	"github.com/energimind/powermesh-core/services/models"
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

func Test_validateCode(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		code    string
		wantErr bool
	}{
		"valid": {
			code:    "code",
			wantErr: false,
		},
		"empty": {
			code:    "",
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateCode(test.code)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateName(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		name    string
		wantErr bool
	}{
		"valid": {
			name:    "name",
			wantErr: false,
		},
		"empty": {
			name:    "",
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateName(test.name)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateModelData(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		data    models.ModelData
		wantErr bool
	}{
		"valid": {
			data: models.ModelData{
				Code: "code",
				Name: "name",
			},
			wantErr: false,
		},
		"invalid-code": {
			data: models.ModelData{
				Code: "",
				Name: "name",
			},
			wantErr: true,
		},
		"invalid-name": {
			data: models.ModelData{
				Code: "code",
				Name: "",
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateModelData(test.data)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
