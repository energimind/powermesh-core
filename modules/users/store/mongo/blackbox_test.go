package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/energimind/go-kit/testutil/mongodb"
	"github.com/energimind/powermesh-core/modules/users"
	"github.com/energimind/powermesh-core/modules/users/store/mongo"
)

var mongoEnv mongodb.MongoEnvironment

// TestMain sets up the MongoDB test environment for all blackbox
// tests in the repository_test package.
func TestMain(m *testing.M) {
	cleanUp, err := mongoEnv.Start()
	defer cleanUp()

	if err != nil {
		panic(err)
	}

	m.Run()
}

func testUser() users.User {
	return users.User{
		ID:         "1",
		ExternalID: "ex1",
		Username:   "username1",
		Email:      "user1@somewhere.com",
	}
}

func testUser2() users.User {
	return users.User{
		ID:         "2",
		ExternalID: "ex2",
		Username:   "username2",
		Email:      "user2@somewhere.com",
	}
}

func withStore(t *testing.T, f func(*testing.T, context.Context, *mongo.UserStore)) {
	t.Helper()

	db, closer := mongoEnv.NewInstance()
	defer closer()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	store := mongo.NewUserStore(db)

	f(t, ctx, store)
}
