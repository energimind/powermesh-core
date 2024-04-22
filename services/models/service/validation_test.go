package service

import (
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/stretchr/testify/require"
)

func Test_requireString(t *testing.T) {
	t.Parallel()

	require.NoError(t, requireString("value", "name"))
	require.Error(t, requireString("", "name"))
	require.IsType(t, errorz.ValidationError{}, requireString("", "name"))
}
