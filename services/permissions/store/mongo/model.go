package mongo

import (
	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/services/permissions"
)

// storeRoleBinding represents a role binding in the MongoDB store.
type storeRoleBinding struct {
	ID           string                   `bson:"id"`
	OwnerID      string                   `bson:"ownerId"`
	UserID       string                   `bson:"userId"`
	ResourceID   string                   `bson:"resourceId"`
	ResourceType permissions.ResourceType `bson:"resourceType"`
	Role         access.Role              `bson:"role"`
}
