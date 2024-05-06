package models

import (
	"context"

	"github.com/energimind/powermesh-core/access"
)

// ModelService defines a model service.
type ModelService interface {
	CreateModel(ctx context.Context, actor access.Actor, data ModelData) (Model, error)
	UpdateModel(ctx context.Context, actor access.Actor, id string, data ModelData) (Model, error)
	DeleteModel(ctx context.Context, actor access.Actor, id string) error
	GetModel(ctx context.Context, id string) (Model, error)
	GetModelsByIDs(ctx context.Context, ids []string) ([]Model, error)
}

// ModelData defines the model data. It is used to create or update a model.
type ModelData struct {
	Code string
	Name string
}

// MeshService defines a mesh service.
type MeshService interface {
	meshOperations
	nodeOperations
	relationOperations
}

// meshOperations defines the operations on meshes.
type meshOperations interface {
	CreateMesh(ctx context.Context, actor access.Actor, modelID string, data MeshData) (Mesh, error)
	UpdateMesh(ctx context.Context, actor access.Actor, modelID string, data MeshData) (Mesh, error)
	MergeMesh(ctx context.Context, actor access.Actor, modelID string, data MeshData) error
	DeleteMesh(ctx context.Context, actor access.Actor, modelID string) error
	GetMesh(ctx context.Context, modelID string) (Mesh, error)
}

// nodeOperations defines the operations on nodes.
type nodeOperations interface {
	CreateNode(ctx context.Context, actor access.Actor, modelID string, data NodeData) (Node, error)
	UpdateNode(ctx context.Context, actor access.Actor, modelID, nodeID string, data NodeData) (Node, error)
	DeleteNode(ctx context.Context, actor access.Actor, modelID, nodeID string) error
	GetNode(ctx context.Context, modelID, nodeID string) (Node, error)
}

// relationOperations defines the operations on relations.
type relationOperations interface {
	CreateRelation(ctx context.Context, actor access.Actor, modelID string, data RelationData) (Relation, error)
	UpdateRelation(ctx context.Context, actor access.Actor, modelID, relationID string, data RelationData) (Relation, error)
	DeleteRelation(ctx context.Context, actor access.Actor, modelID, relationID string) error
	GetRelation(ctx context.Context, modelID, relationID string) (Relation, error)
}

// MeshData defines the mesh data. It is used to create or update a mesh.
type MeshData struct {
	Code string // mesh code, copy from model
}

// NodeData defines the node data. It is used to create or update a node.
type NodeData struct {
	Kind  string  // node kind/type
	Code  string  // node code (optional)
	Name  string  // node name (optional)
	Props PropBag // custom node properties
}

// RelationData defines the relation data. It is used to create or update a relation.
type RelationData struct {
	Kind  string  // relation kind/type
	From  string  // public ID of the start node
	To    string  // public ID of the end node
	Props PropBag // custom relation properties
}
