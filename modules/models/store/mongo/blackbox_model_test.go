package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/energimind/powermesh-core/modules/models"
	"github.com/energimind/powermesh-core/modules/models/store/mongo"
)

func testModel() models.Model {
	return models.Model{
		ID:          "1",
		Code:        "code1",
		Name:        "model1",
		Description: "description1",
	}
}

func testModel2() models.Model {
	return models.Model{
		ID:          "2",
		Code:        "code2",
		Name:        "model2",
		Description: "description2",
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
