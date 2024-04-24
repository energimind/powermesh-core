package mongo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mapper(t *testing.T) {
	t.Parallel()

	require.Equal(t, validStoreRoleBinding, toStoreRoleBinding(validModelRoleBinding))
	require.Equal(t, validModelRoleBinding, fromStoreRoleBinding(validStoreRoleBinding))
}
