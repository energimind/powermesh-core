package mongo

// storeModel models a model in the MongoDB store.
type storeModel struct {
	ID   string `bson:"id"`
	Code string `bson:"code"`
	Name string `bson:"name"`
}
