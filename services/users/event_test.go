package users

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractUserEvent(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		ue, ok := ExtractUserEvent(UserEvent{EventHeader: EventHeader{Type: UserCreated}})

		require.True(t, ok)
		require.NotZero(t, ue)
		require.True(t, ue.IsUserEvent())
	})

	t.Run("failure", func(t *testing.T) {
		ue, ok := ExtractUserEvent(nil)

		require.False(t, ok)
		require.Zero(t, ue)
	})
}
