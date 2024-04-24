package mongo

import (
	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/services/permissions"
)

var (
	validModelRoleBinding = permissions.RoleBinding{
		ID:           "1",
		OwnerID:      "user0",
		UserID:       "user1",
		ResourceID:   "res1",
		ResourceType: permissions.ResourceTypeModel,
		Role:         access.RoleAdmin,
	}
	validStoreRoleBinding = storeRoleBinding{
		ID:           validModelRoleBinding.ID,
		OwnerID:      validModelRoleBinding.OwnerID,
		UserID:       validModelRoleBinding.UserID,
		ResourceID:   validModelRoleBinding.ResourceID,
		ResourceType: validModelRoleBinding.ResourceType,
		Role:         validModelRoleBinding.Role,
	}
)
