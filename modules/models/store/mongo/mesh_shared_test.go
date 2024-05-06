package mongo

import "github.com/energimind/powermesh-core/modules/models"

var (
	validModelMesh = models.Mesh{
		ModelID: "model-id",
		Code:    "model-code",
		Nodes: map[string]models.Node{
			"node-id": {
				ID: "node-id",
			},
		},
		Relations: map[string]models.Relation{
			"relation-id": {
				ID: "relation-id",
			},
		},
	}
	validStoreMesh = storeMesh{
		ModelID: validModelMesh.ModelID,
		Code:    validModelMesh.Code,
		Nodes: []storeNode{
			toStoreNode(validModelMesh.Nodes["node-id"]),
		},
		Relations: []storeRelation{
			toStoreRelation(validModelMesh.Relations["relation-id"]),
		},
	}
)
