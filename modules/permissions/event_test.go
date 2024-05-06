package permissions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractRoleBindingEvent(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		rbe, ok := ExtractRoleBindingEvent(RoleBindingEvent{EventHeader: EventHeader{Type: RoleBindingCreated}})

		require.True(t, ok)
		require.NotZero(t, rbe)
		require.True(t, rbe.IsRoleBindingEvent())
	})

	t.Run("failure", func(t *testing.T) {
		rbe, ok := ExtractRoleBindingEvent(nil)

		require.False(t, ok)
		require.Zero(t, rbe)
	})
}
