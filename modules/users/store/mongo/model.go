package mongo

// storeUser models a user in the MongoDB store.
type storeUser struct {
	ID       string `bson:"id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
}
