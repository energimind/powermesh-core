package access

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewActor(t *testing.T) {
	t.Parallel()

	require.Equal(t, Actor{
		UserID: "user",
		Role:   RoleAdmin,
	}, NewActor("user", RoleAdmin))
}

func TestActor_String(t *testing.T) {
	t.Parallel()

	a := NewActor("user", RoleAdmin)

	require.Equal(t, "user (admin)", a.String())
}
