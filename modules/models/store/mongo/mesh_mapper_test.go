package mongo

import (
	"testing"

	"github.com/energimind/powermesh-core/modules/models"
	"github.com/stretchr/testify/require"
)

func Test_meshMappers(t *testing.T) {
	t.Parallel()

	require.Equal(t, validStoreMesh, toStoreMesh(validModelMesh))
	require.Equal(t, validModelMesh, fromStoreMesh(validStoreMesh))
}

func Test_extractNodes(t *testing.T) {
	t.Parallel()

	exp := []models.Node{validModelMesh.Nodes["node-id"]}

	require.Equal(t, exp, extractNodes(validStoreMesh))
}

func Test_extractFirstNode(t *testing.T) {
	t.Parallel()

	require.Equal(t, validModelMesh.Nodes["node-id"], extractFirstNode(validStoreMesh))
}

func Test_extractRelations(t *testing.T) {
	t.Parallel()

	exp := []models.Relation{validModelMesh.Relations["relation-id"]}

	require.Equal(t, exp, extractRelations(validStoreMesh))
}

func Test_extractFirstRelation(t *testing.T) {
	t.Parallel()

	require.Equal(t, validModelMesh.Relations["relation-id"], extractFirstRelation(validStoreMesh))
}

func Test_mergeMeshUpdate(t *testing.T) {
	t.Parallel()

	update := mergeMeshUpdate(validModelMesh)

	require.Equal(t, map[string]any{
		fieldCode:      validModelMesh.Code,
		fieldNodes:     validStoreMesh.Nodes,
		fieldRelations: validStoreMesh.Relations,
	}, update)
}
