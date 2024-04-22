package models

// Model defines a model.
type Model struct {
	ID   string
	Code string
	Name string
}

// Mesh represents a model mesh.
type Mesh struct {
	ModelID   string              // public ID of the model
	Code      string              // mesh code, copy from model
	Nodes     map[string]Node     // public node ID -> Node
	Relations map[string]Relation // public relation ID -> Relation
}

// Node represents a node in the mesh.
type Node struct {
	ID    string  // public ID
	Kind  string  // node kind/type
	Code  string  // node code (optional)
	Props PropBag // custom node properties
}

// Relation represents a relation in the mesh.
type Relation struct {
	ID    string  // public ID
	Kind  string  // relation kind/type
	From  string  // public ID of the start node
	To    string  // public ID of the end node
	Props PropBag // custom relation properties
}

// PropBag represents a set of property sets.
// The bag is divided into named sections.
type PropBag map[string]PropSection

// PropSection represents a set of named properties.
type PropSection map[string]any
