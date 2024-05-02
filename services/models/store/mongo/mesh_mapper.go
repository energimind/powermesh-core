package mongo

import (
	"github.com/energimind/powermesh-core/services/models"
)

func toStoreMesh(m models.Mesh) storeMesh {
	return storeMesh{
		ModelID:   m.ModelID,
		Code:      m.Code,
		Nodes:     toStoreNodes(m.Nodes),
		Relations: toStoreRelations(m.Relations),
	}
}

func fromStoreMesh(m storeMesh) models.Mesh {
	return models.Mesh{
		ModelID:   m.ModelID,
		Code:      m.Code,
		Nodes:     fromStoreNodes(m.Nodes),
		Relations: fromStoreRelations(m.Relations),
	}
}

func toStoreNodes(ns map[string]models.Node) []storeNode {
	nodes := make([]storeNode, len(ns))

	i := 0

	for _, n := range ns {
		nodes[i] = toStoreNode(n)
		i++
	}

	return nodes
}

func fromStoreNodes(ns []storeNode) map[string]models.Node {
	nodes := make(map[string]models.Node)

	for _, n := range ns {
		nodes[n.ID] = fromStoreNode(n)
	}

	return nodes
}

func toStoreNode(n models.Node) storeNode {
	return storeNode{
		ID:    n.ID,
		Kind:  n.Kind,
		Code:  n.Code,
		Props: n.Props,
	}
}

func fromStoreNode(n storeNode) models.Node {
	return models.Node{
		ID:    n.ID,
		Kind:  n.Kind,
		Code:  n.Code,
		Props: n.Props,
	}
}

func toStoreRelations(rs map[string]models.Relation) []storeRelation {
	relations := make([]storeRelation, len(rs))

	i := 0

	for _, r := range rs {
		relations[i] = toStoreRelation(r)
		i++
	}

	return relations
}

func fromStoreRelations(rs []storeRelation) map[string]models.Relation {
	relations := make(map[string]models.Relation)

	for _, r := range rs {
		relations[r.ID] = fromStoreRelation(r)
	}

	return relations
}

func toStoreRelation(r models.Relation) storeRelation {
	return storeRelation{
		ID:    r.ID,
		Kind:  r.Kind,
		From:  r.From,
		To:    r.To,
		Props: r.Props,
	}
}

func fromStoreRelation(r storeRelation) models.Relation {
	return models.Relation{
		ID:    r.ID,
		Kind:  r.Kind,
		From:  r.From,
		To:    r.To,
		Props: r.Props,
	}
}

func extractFirstNode(m storeMesh) models.Node {
	return fromStoreNode(m.Nodes[0])
}

func extractFirstRelation(m storeMesh) models.Relation {
	return fromStoreRelation(m.Relations[0])
}

// mergeMeshUpdate returns a MongoDB update document for a mesh.
// The update document contains only the fields that are set in the mesh.
func mergeMeshUpdate(m models.Mesh) map[string]any {
	update := map[string]any{}

	if m.Code != "" {
		update[fieldCode] = m.Code
	}

	if len(m.Nodes) > 0 {
		update[fieldNodes] = toStoreNodes(m.Nodes)
	}

	if len(m.Relations) > 0 {
		update[fieldRelations] = toStoreRelations(m.Relations)
	}

	return update
}
