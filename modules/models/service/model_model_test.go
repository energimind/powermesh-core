package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_modelFromData(t *testing.T) {
	require.Equal(t,
		validModel,
		modelFromData(validModelID, validModelData),
	)
}
