package mongo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_toStoreModel(t *testing.T) {
	t.Parallel()

	require.Equal(t, validStoreModel, toStoreModel(validModelModel))
	require.Equal(t, validModelModel, fromStoreModel(validStoreModel))
}
