package service

import (
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/stretchr/testify/require"
)

func Test_canAct(t *testing.T) {
	t.Parallel()

	canActMap := map[access.Role]bool{
		access.RoleAdmin:   true,
		access.RoleCreator: true,
	}

	for _, role := range access.AllRoles {
		if _, preset := canActMap[role]; !preset {
			canActMap[role] = false
		}
	}

	for role, wantCan := range canActMap {
		t.Run(role.String(), func(t *testing.T) {
			can := canAct(access.Actor{Role: role})

			if wantCan {
				require.True(t, can)
			} else {
				require.False(t, can)
			}
		})
	}
}
