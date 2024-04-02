package accessctx

import (
	"context"
	"testing"

	"github.com/energimind/powermesh-core/access"
	"github.com/stretchr/testify/require"
)

func TestActor(t *testing.T) {
	t.Parallel()

	t.Run("actorInjected-returningInjected", func(t *testing.T) {
		t.Parallel()

		actor := access.NewActor("userID", access.RoleAdmin)
		ctx := context.Background()

		ctx = WithActor(ctx, actor)

		require.Equal(t, actor, Actor(ctx))
	})

	t.Run("noActorInjected-returningEmpty", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		require.Empty(t, Actor(ctx))
	})
}
