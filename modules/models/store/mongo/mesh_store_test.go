package mongo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/models/store/mongo"
	"github.com/stretchr/testify/require"
)

func TestMeshStore_CreateMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		mesh := testMesh()

		require.NoError(t, store.CreateMesh(ctx, mesh))

		foundMesh, err := store.GetMesh(ctx, mesh.ModelID)

		require.NoError(t, err)
		require.Equal(t, mesh, foundMesh)
	})
}

func TestMeshStore_UpdateMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.UpdateMesh(ctx, testMesh()))
		})

		t.Run("success", func(t *testing.T) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			newNode := testNode()
			newNode.Code = "new-node"

			mesh.Code = "new-code"
			mesh.Nodes[newNode.ID] = newNode

			require.NoError(t, store.UpdateMesh(ctx, mesh))

			updatedMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.Equal(t, mesh, updatedMesh)
		})
	})
}

func TestMeshStore_MergeMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.MergeMesh(ctx, testMesh()))
		})

		t.Run("success", func(t *testing.T) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			mesh.Code = "new-code"

			require.NoError(t, store.MergeMesh(ctx, mesh))

			updatedMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.Equal(t, mesh, updatedMesh)
		})
	})
}

func TestMeshStore_DeleteMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		t.Run("not-found", func(t *testing.T) {
			require.IsType(t, errorz.NotFoundError{}, store.DeleteMesh(ctx, "missing"))
		})

		t.Run("success", func(t *testing.T) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			require.NoError(t, store.DeleteMesh(ctx, mesh.ModelID))

			_, err := store.GetMesh(ctx, mesh.ModelID)

			require.IsType(t, errorz.NotFoundError{}, err)
		})
	})
}

func TestMeshStore_GetMesh(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		t.Run("not-found", func(t *testing.T) {
			_, err := store.GetMesh(ctx, "missing")

			require.IsType(t, errorz.NotFoundError{}, err)
		})

		t.Run("success", func(t *testing.T) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			foundMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.Equal(t, mesh, foundMesh)
		})
	})
}

func TestMeshStore_CreateNode(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		mesh := testMesh()

		// create without nodes
		mesh.Nodes = nil

		require.NoError(t, store.CreateMesh(ctx, mesh))

		node := testNode()

		require.NoError(t, store.CreateNode(ctx, mesh.ModelID, node))

		updatedMesh, err := store.GetMesh(ctx, mesh.ModelID)

		require.NoError(t, err)
		require.Contains(t, updatedMesh.Nodes, node.ID)
	})
}

func TestMeshStore_UpdateNode(t *testing.T) {
	t.Parallel()

	t.Run("not-found", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			// create without nodes
			mesh.Nodes = nil

			require.NoError(t, store.CreateMesh(ctx, mesh))

			require.IsType(t, errorz.NotFoundError{}, store.UpdateNode(ctx, mesh.ModelID, testNode()))
		})
	})

	t.Run("success", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			// create without nodes
			mesh.Nodes = nil

			require.NoError(t, store.CreateMesh(ctx, mesh))

			node := testNode()

			require.NoError(t, store.CreateNode(ctx, mesh.ModelID, node))

			node.Code = "new-code"

			require.NoError(t, store.UpdateNode(ctx, mesh.ModelID, node))

			updatedMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.Equal(t, node, updatedMesh.Nodes[node.ID])
		})
	})
}

func TestMeshStore_DeleteNode(t *testing.T) {
	t.Parallel()

	t.Run("not-found", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			// create without nodes
			mesh.Nodes = nil

			require.NoError(t, store.CreateMesh(ctx, mesh))

			require.IsType(t, errorz.NotFoundError{}, store.DeleteNode(ctx, mesh.ModelID, "missing"))
		})
	})

	t.Run("success", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			node := mesh.Nodes["1"]

			require.NotZero(t, node)

			updatedMesh1, _ := store.GetMesh(ctx, mesh.ModelID)
			fmt.Printf("%+v\n", updatedMesh1)

			require.NoError(t, store.DeleteNode(ctx, mesh.ModelID, node.ID))

			updatedMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.NotContains(t, updatedMesh.Nodes, node.ID)
		})
	})
}

func TestMeshStore_GetNode(t *testing.T) {
	t.Parallel()

	t.Run("not-found", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			_, err := store.GetNode(ctx, mesh.ModelID, "missing")

			require.IsType(t, errorz.NotFoundError{}, err)
		})
	})

	t.Run("success", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			node := mesh.Nodes["1"]

			require.NotZero(t, node)

			foundNode, err := store.GetNode(ctx, mesh.ModelID, node.ID)

			require.NoError(t, err)
			require.Equal(t, node, foundNode)
		})
	})
}

func TestMeshStore_CreateRelation(t *testing.T) {
	t.Parallel()

	withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
		mesh := testMesh()

		// create without relations
		mesh.Relations = nil

		require.NoError(t, store.CreateMesh(ctx, mesh))

		relation := testRelation()

		require.NoError(t, store.CreateRelation(ctx, mesh.ModelID, relation))

		foundMesh, err := store.GetMesh(ctx, mesh.ModelID)

		require.NoError(t, err)
		require.Contains(t, foundMesh.Relations, relation.ID)
	})
}

func TestMeshStore_UpdateRelation(t *testing.T) {
	t.Parallel()

	t.Run("not-found", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			// create without relations
			mesh.Relations = nil

			require.NoError(t, store.CreateMesh(ctx, mesh))

			require.IsType(t, errorz.NotFoundError{}, store.UpdateRelation(ctx, mesh.ModelID, testRelation()))
		})
	})

	t.Run("success", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			// create without relations
			mesh.Relations = nil

			require.NoError(t, store.CreateMesh(ctx, mesh))

			relation := testRelation()

			require.NoError(t, store.CreateRelation(ctx, mesh.ModelID, relation))

			relation.To = "new-to"

			require.NoError(t, store.UpdateRelation(ctx, mesh.ModelID, relation))

			updatedMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.Equal(t, relation, updatedMesh.Relations[relation.ID])
		})
	})
}

func TestMeshStore_DeleteRelation(t *testing.T) {
	t.Parallel()

	t.Run("not-found", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			// create without relations
			mesh.Relations = nil

			require.NoError(t, store.CreateMesh(ctx, mesh))

			require.IsType(t, errorz.NotFoundError{}, store.DeleteRelation(ctx, mesh.ModelID, "missing"))
		})
	})

	t.Run("success", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			relation := mesh.Relations["1"]

			require.NotZero(t, relation)

			require.NoError(t, store.DeleteRelation(ctx, mesh.ModelID, relation.ID))

			updatedMesh, err := store.GetMesh(ctx, mesh.ModelID)

			require.NoError(t, err)
			require.NotContains(t, updatedMesh.Relations, relation.ID)
		})
	})
}

func TestMeshStore_GetRelation(t *testing.T) {
	t.Parallel()

	t.Run("not-found", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			_, err := store.GetRelation(ctx, mesh.ModelID, "missing")

			require.IsType(t, errorz.NotFoundError{}, err)
		})
	})

	t.Run("success", func(t *testing.T) {
		withMeshStore(t, func(t *testing.T, ctx context.Context, store *mongo.MeshStore) {
			mesh := testMesh()

			require.NoError(t, store.CreateMesh(ctx, mesh))

			relation := mesh.Relations["1"]

			require.NotZero(t, relation)

			foundRelation, err := store.GetRelation(ctx, mesh.ModelID, relation.ID)

			require.NoError(t, err)
			require.Equal(t, relation, foundRelation)
		})
	})
}
