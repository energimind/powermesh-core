package mongo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mapper(t *testing.T) {
	t.Parallel()

	require.Equal(t, validStoreUser, toStoreUser(validModelUser))
	require.Equal(t, validModelUser, fromStoreUser(validStoreUser))
}
