package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/energimind/go-kit/testutil/mongodb"
	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/services/permissions"
	"github.com/energimind/powermesh-core/services/permissions/store/mongo"
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

func testRoleBinding() permissions.RoleBinding {
	return permissions.RoleBinding{
		ID:           "1",
		OwnerID:      "user0",
		UserID:       "user1",
		ResourceID:   "res1",
		ResourceType: permissions.ResourceTypeModel,
		Role:         access.RoleAdmin,
	}
}

func testRoleBinding2() permissions.RoleBinding {
	return permissions.RoleBinding{
		ID:           "3",
		OwnerID:      "user1",
		UserID:       "user2",
		ResourceID:   "res2",
		ResourceType: permissions.ResourceTypeModel,
		Role:         access.RoleAdmin,
	}
}

func withStore(t *testing.T, f func(*testing.T, context.Context, *mongo.PermissionStore)) {
	t.Helper()

	db, closer := mongoEnv.NewInstance()
	defer closer()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	store := mongo.NewPermissionStore(db)

	f(t, ctx, store)
}
