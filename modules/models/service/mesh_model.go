package service

import "github.com/energimind/powermesh-core/modules/models"

func meshFromData(modelID string, data models.MeshData) models.Mesh {
	return models.Mesh{
		ModelID: modelID,
		Code:    data.Code,
	}
}

func nodeFromData(id string, data models.NodeData) models.Node {
	return models.Node{
		ID:    id,
		Kind:  data.Kind,
		Code:  data.Code,
		Props: data.Props,
	}
}

func relationFromData(id string, data models.RelationData) models.Relation {
	return models.Relation{
		ID:    id,
		Kind:  data.Kind,
		From:  data.From,
		To:    data.To,
		Props: data.Props,
	}
}
