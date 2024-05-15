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
	testAddressID    = "2"
	testDomainPerson = person{
		ID:   testID,
		Name: "John",
		Age:  30,
		Addresses: []address{
			testDomainAddress,
		},
	}
	testDBPerson = dbPerson{
		ID:   testDomainPerson.ID,
		Name: testDomainPerson.Name,
		Age:  testDomainPerson.Age,
		Addresses: []dbAddress{
			testDBAddress,
		},
	}
	testDomainAddress = address{
		ID:     testAddressID,
		Street: "Main St",
	}
	testDBAddress = dbAddress(testDomainAddress)
)

type person struct {
	ID        string
	Name      string
	Age       int
	Addresses []address
}

type dbPerson struct {
	ID        string      `bson:"id"`
	Name      string      `bson:"name"`
	Age       int         `bson:"age"`
	Addresses []dbAddress `bson:"addresses"`
}

type address struct {
	ID     string
	Street string
}

type dbAddress struct {
	ID     string `bson:"id"`
	Street string `bson:"street"`
}

func toDBPerson(p person) dbPerson {
	return dbPerson{
		ID:        p.ID,
		Name:      p.Name,
		Age:       p.Age,
		Addresses: toDBAddresses(p.Addresses),
	}
}

func fromDBPerson(p dbPerson) person {
	return person{
		ID:        p.ID,
		Name:      p.Name,
		Age:       p.Age,
		Addresses: fromDBAddresses(p.Addresses),
	}
}

func toDBAddress(a address) dbAddress {
	return dbAddress(a)
}

func fromDBAddress(a dbAddress) address {
	return address(a)
}

func toDBAddresses(as []address) []dbAddress {
	var dbs []dbAddress

	for _, a := range as {
		dbs = append(dbs, toDBAddress(a))
	}

	return dbs
}

func fromDBAddresses(as []dbAddress) []address {
	var ads []address

	for _, a := range as {
		ads = append(ads, fromDBAddress(a))
	}

	return ads
}

func extractFirstAddress(p person) address {
	return p.Addresses[0]
}

func projectName(p dbPerson) string {
	return p.Name
}

type forcedError struct{}

func (e forcedError) Error() string {
	return "forced error"
}

type mockCollection struct {
	t              *testing.T
	caller         string
	insertOne      func() (*mongo.InsertOneResult, error)
	updateOne      func() (*mongo.UpdateResult, error)
	deleteOne      func() (*mongo.DeleteResult, error)
	deleteMany     func() (*mongo.DeleteResult, error)
	findOne        func() *mongo.SingleResult
	find           func() (*mongo.Cursor, error)
	countDocuments func() (int64, error)
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

	switch c.caller {
	case "MergeFields":
		require.Equal(c.t, bson.M{"id": testID}, fm)
		require.Equal(c.t, bson.M{"$set": bson.M{"name": "John", "age": 30}}, um)
	case "UpdateOne":
		require.Equal(c.t, bson.M{"id": testID}, fm)
		require.Equal(c.t, bson.M{"$set": testDBPerson}, um)
	case "EmbeddedPull":
		require.Equal(c.t, bson.M{"id": testID}, fm)
		require.Equal(c.t, bson.M{"$pull": bson.M{"address": bson.M{"id": testAddressID}}}, um)
	case "EmbeddedPush":
		require.Equal(c.t, bson.M{"id": testID}, fm)
		require.Equal(c.t, bson.M{"$push": bson.M{"address": testDBAddress}}, um)
	case "EmbeddedUpdate":
		require.Equal(c.t, bson.M{"id": testID, "address.id": testAddressID}, fm)
		require.Equal(c.t, bson.M{"$set": bson.M{"address.$": testDBAddress}}, um)
	default:
		require.Fail(c.t, "unexpected caller: %s", c.caller)
	}

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

func (c *mockCollection) DeleteMany(_ context.Context, filter interface{}, _ ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c.t.Helper()

	require.NotNil(c.t, filter)

	fm := filter.(bson.M)

	require.Equal(c.t, bson.M{"age": bson.M{"$gt": 20}}, fm)

	if c.deleteMany == nil {
		return nil, errors.New("deleteMany not implemented")
	}

	return c.deleteMany()
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

func (c *mockCollection) CountDocuments(_ context.Context, filter interface{}, _ ...*options.CountOptions) (int64, error) {
	c.t.Helper()

	require.NotNil(c.t, filter)

	fm := filter.(bson.M)

	require.Equal(c.t, bson.M{"age": bson.M{"$gt": 20}}, fm)

	if c.countDocuments == nil {
		return 0, errors.New("countDocuments not implemented")
	}

	return c.countDocuments()
}

func (c *mockCollection) Name() string {
	return "persons"
}
