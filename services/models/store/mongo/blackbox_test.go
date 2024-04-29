package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/energimind/go-kit/testutil/mongodb"
	"github.com/energimind/powermesh-core/services/models"
	"github.com/energimind/powermesh-core/services/models/store/mongo"
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

func testModel() models.Model {
	return models.Model{
		ID:   "1",
		Code: "code1",
		Name: "model1",
	}
}

func testModel2() models.Model {
	return models.Model{
		ID:   "2",
		Code: "code2",
		Name: "model2",
	}
}

func withModelStore(t *testing.T, f func(*testing.T, context.Context, *mongo.ModelStore)) {
	t.Helper()

	db, closer := mongoEnv.NewInstance()
	defer closer()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	store := mongo.NewModelStore(db)

	f(t, ctx, store)
}
