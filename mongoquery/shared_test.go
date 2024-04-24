package mongoquery

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	testID           = "1"
	testDomainPerson = person{
		ID:   testID,
		Name: "John",
		Age:  30,
	}
	testDBPerson = dbPerson{
		ID:   testID,
		Name: "John",
		Age:  30,
	}
)

type person struct {
	ID   string
	Name string
	Age  int
}

type dbPerson struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func toDBPerson(p person) dbPerson {
	return dbPerson(p)
}

func fromDBPerson(p dbPerson) person {
	return person(p)
}

type forcedError struct{}

func (e forcedError) Error() string {
	return "forced error"
}

type mockCollection struct {
	t         *testing.T
	insertOne func() (*mongo.InsertOneResult, error)
	updateOne func() (*mongo.UpdateResult, error)
	deleteOne func() (*mongo.DeleteResult, error)
	findOne   func() *mongo.SingleResult
	find      func() (*mongo.Cursor, error)
}

// Ensure mockCollection implements the collection interface.
var _ collection = (*mockCollection)(nil)

func (c *mockCollection) InsertOne(_ context.Context, document interface{}, _ ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c.t.Helper()

	require.NotNil(c.t, document)
	require.Equal(c.t, testDBPerson, document)

	if c.insertOne == nil {
		return nil, errors.New("insertOne not implemented")
	}

	return c.insertOne()
}

func (c *mockCollection) UpdateOne(_ context.Context, filter interface{}, update interface{}, _ ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	c.t.Helper()

	require.NotNil(c.t, filter)
	require.NotNil(c.t, update)

	fm := filter.(bson.M)
	um := update.(bson.M)

	require.Equal(c.t, testID, fm["id"])
	require.Equal(c.t, bson.M{"$set": testDBPerson}, um)

	if c.updateOne == nil {
		return nil, errors.New("updateOne not implemented")
	}

	return c.updateOne()
}

func (c *mockCollection) DeleteOne(_ context.Context, filter interface{}, _ ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c.t.Helper()

	require.NotNil(c.t, filter)

	fm := filter.(bson.M)

	require.Equal(c.t, testID, fm["id"])

	if c.deleteOne == nil {
		return nil, errors.New("deleteOne not implemented")
	}

	return c.deleteOne()
}

func (c *mockCollection) FindOne(_ context.Context, filter interface{}, _ ...*options.FindOneOptions) *mongo.SingleResult {
	c.t.Helper()

	require.NotNil(c.t, filter)

	fm := filter.(bson.M)

	require.Equal(c.t, testID, fm["id"])

	if c.findOne == nil {
		return nil
	}

	return c.findOne()
}

func (c *mockCollection) Find(_ context.Context, filter interface{}, _ ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	c.t.Helper()

	require.NotNil(c.t, filter)

	if c.find == nil {
		return nil, errors.New("find not implemented")
	}

	fm := filter.(bson.M)

	require.Equal(c.t, bson.M{"age": bson.M{"$gt": 20}}, fm)

	return c.find()
}

func (c *mockCollection) Name() string {
	return "persons"
}
