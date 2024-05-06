package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_meshFromData(t *testing.T) {
	require.Equal(t,
		validMesh,
		meshFromData(validModelID, validMeshData),
	)
}

func Test_nodeFromData(t *testing.T) {
	require.Equal(t,
		validNode,
		nodeFromData(validNodeID, validNodeData),
	)
}

func Test_relationFromData(t *testing.T) {
	require.Equal(t,
		validRelation,
		relationFromData(validRelationID, validRelationData),
	)
}
