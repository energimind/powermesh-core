package service

import (
	"testing"

	"github.com/energimind/powermesh-core/services/models"
	"github.com/stretchr/testify/require"
)

func Test_validateID(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateID("1"))
	require.Error(t, validateID(""))
}

func Test_validateCode(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateCode("code"))
	require.Error(t, validateCode(""))
}

func Test_validateName(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateName("name"))
	require.Error(t, validateName(""))
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
