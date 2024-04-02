package accessctx

import (
	"context"

	"github.com/energimind/powermesh-core/access"
)

// internal type to prevent collisions with other packages.
type ctxKey struct{}

// WithActor returns a new context with the given actor.
func WithActor(ctx context.Context, actor access.Actor) context.Context {
	return context.WithValue(ctx, ctxKey{}, actor)
}

// Actor returns the actor from the given context.
// The empty actor is returned if the actor was not found in the context.
func Actor(ctx context.Context) access.Actor {
	act, ok := ctx.Value(ctxKey{}).(access.Actor)
	if !ok {
		return access.Actor{}
	}

	return act
}
