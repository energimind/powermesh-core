package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/energimind/powermesh-core/modules/models"
	"github.com/energimind/powermesh-core/modules/models/store/mongo"
)

func testMesh() models.Mesh {
	n, r := testNode(), testRelation()

	return models.Mesh{
		ModelID:   "1",
		Code:      "code1",
		Nodes:     map[string]models.Node{n.ID: n},
		Relations: map[string]models.Relation{r.ID: r},
	}
}

func testNode() models.Node {
	return models.Node{
		ID:   "1",
		Kind: "kind1",
		Code: "code1",
		Props: models.PropBag{
			"n-key1": models.PropSection{
				"n-subkey1": "value1",
			},
		},
	}
}

func testRelation() models.Relation {
	return models.Relation{
		ID:   "1",
		Kind: "kind1",
		From: "1",
		To:   "2",
		Props: models.PropBag{
			"p-key1": models.PropSection{
				"p-subkey1": "value1",
			},
		},
	}
}

func withMeshStore(t *testing.T, f func(*testing.T, context.Context, *mongo.MeshStore)) {
	t.Helper()

	db, closer := mongoEnv.NewInstance()
	defer closer()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	store := mongo.NewMeshStore(db)

	f(t, ctx, store)
}
