package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractModelEvent(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		me, ok := ExtractModelEvent(ModelEvent{EventHeader: EventHeader{Type: ModelCreated}})

		require.True(t, ok)
		require.NotZero(t, me)
		require.True(t, me.IsModelEvent())
		require.False(t, me.IsMeshEvent())
	})

	t.Run("failure", func(t *testing.T) {
		me, ok := ExtractModelEvent(nil)

		require.False(t, ok)
		require.Zero(t, me)
	})
}

func TestExtractMeshEvent(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		me, ok := ExtractMeshEvent(MeshEvent{EventHeader: EventHeader{Type: MeshCreated}})

		require.True(t, ok)
		require.NotZero(t, me)
		require.True(t, me.IsMeshEvent())
		require.False(t, me.IsModelEvent())
	})

	t.Run("failure", func(t *testing.T) {
		me, ok := ExtractMeshEvent(nil)

		require.False(t, ok)
		require.Zero(t, me)
	})
}
