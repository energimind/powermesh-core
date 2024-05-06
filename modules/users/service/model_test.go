package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_userFromData(t *testing.T) {
	require.Equal(t,
		validUser,
		userFromData(validUserID, validUserData),
	)
}
