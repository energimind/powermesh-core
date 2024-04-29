package mongo

import "github.com/energimind/powermesh-core/services/models"

// storeMesh models a mesh in the MongoDB store.
type storeMesh struct {
	ModelID   string          `bson:"modelId"`
	Code      string          `bson:"code"`
	Nodes     []storeNode     `bson:"nodes"`
	Relations []storeRelation `bson:"relations"`
}

// storeNode models a node in the MongoDB store.
type storeNode struct {
	ID    string         `bson:"id"`
	Kind  string         `bson:"kind"`
	Code  string         `bson:"code"`
	Props models.PropBag `bson:"props"`
}

// storeRelation models a relation in the MongoDB store.
type storeRelation struct {
	ID    string         `bson:"id"`
	Kind  string         `bson:"kind"`
	From  string         `bson:"from"`
	To    string         `bson:"to"`
	Props models.PropBag `bson:"props"`
}
