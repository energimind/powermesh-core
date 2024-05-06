package permissions

import (
	"strconv"

	"github.com/energimind/powermesh-core/access"
)

// RoleBinding represents a role binding.
// Role bindings are used to bind a role to a user and a resource.
type RoleBinding struct {
	ID           string
	OwnerID      string
	UserID       string
	ResourceID   string
	ResourceType ResourceType
	Role         access.Role
}

// ResourceType represents the type of resource.
type ResourceType int

// ResourceTypeModel is the model resource type.
const ResourceTypeModel ResourceType = iota

// AllResourceTypes is a list of all resource types. Used for testing purposes to validate that all
// enum values are covered.
//
//nolint:gochecknoglobals
var AllResourceTypes = []ResourceType{
	ResourceTypeModel,
}

// String returns the string representation of the resource type.
func (o ResourceType) String() string {
	if o == ResourceTypeModel {
		return "model"
	}

	return "ResourceType(" + strconv.Itoa(int(o)) + ")"
}

// IsSupportedResourceType checks if the given resource type is supported.
func IsSupportedResourceType(resourceType ResourceType) bool {
	for _, rt := range AllResourceTypes {
		if rt == resourceType {
			return true
		}
	}

	return false
}
