package service

import (
	"github.com/energimind/powermesh-core/modules/permissions"
)

func roleBindingFromData(id string, data permissions.RoleBindingData) permissions.RoleBinding {
	return permissions.RoleBinding{
		ID:           id,
		OwnerID:      data.OwnerID,
		UserID:       data.UserID,
		ResourceID:   data.ResourceID,
		ResourceType: data.ResourceType,
		Role:         data.Role,
	}
}
