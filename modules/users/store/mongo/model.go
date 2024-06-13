package mongo

// storeUser models a user in the MongoDB store.
type storeUser struct {
	ID          string `bson:"id"`
	ExternalID  string `bson:"externalId"`
	Username    string `bson:"username"`
	DisplayName string `bson:"displayName"`
	Email       string `bson:"email"`
}
