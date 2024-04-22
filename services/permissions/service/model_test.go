package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_roleBindingFromData(t *testing.T) {
	require.Equal(t,
		validRoleBinding,
		roleBindingFromData(validRoleBindingID, validRoleBindingData),
	)
}
