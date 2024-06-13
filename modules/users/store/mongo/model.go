package mongo

// storeUser models a user in the MongoDB store.
type storeUser struct {
	ID         string `bson:"id"`
	ExternalID string `bson:"externalId"`
	Username   string `bson:"username"`
	Email      string `bson:"email"`
}
