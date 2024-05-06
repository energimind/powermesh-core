package mongo

import "github.com/energimind/powermesh-core/modules/permissions"

func toStoreRoleBinding(rb permissions.RoleBinding) storeRoleBinding {
	return storeRoleBinding{
		ID:           rb.ID,
		OwnerID:      rb.OwnerID,
		UserID:       rb.UserID,
		ResourceID:   rb.ResourceID,
		ResourceType: rb.ResourceType,
		Role:         rb.Role,
	}
}

func fromStoreRoleBinding(rb storeRoleBinding) permissions.RoleBinding {
	return permissions.RoleBinding{
		ID:           rb.ID,
		OwnerID:      rb.OwnerID,
		UserID:       rb.UserID,
		ResourceID:   rb.ResourceID,
		ResourceType: rb.ResourceType,
		Role:         rb.Role,
	}
}

func projectResourceID(rb storeRoleBinding) string {
	return rb.ResourceID
}
