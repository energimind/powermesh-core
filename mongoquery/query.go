package mongoquery

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collection is an interface for a MongoDB collection.
// It is used to abstract the MongoDB driver.
type collection interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	CountDocuments(ctx context.Context, filter interface{},
		opts ...*options.CountOptions) (int64, error)
	Name() string
}

// mapper is a function that maps a value of type t to a value of type D.
type mapper[F, T any] func(F) T

// defaultKey is the default key used to identify documents in a collection.
const defaultKey = "id"

// resolveKey returns the key to use for the query.
// It returns the default key if the input key is empty.
func resolveKey(key string) string {
	if key == "" {
		return defaultKey
	}

	return key
}
