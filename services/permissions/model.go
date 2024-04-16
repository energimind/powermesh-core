package permissions

import (
	"strconv"

	"github.com/energimind/powermesh-core/access"
)

// RoleBinding represents a role binding.
// Role bindings are used to bind a role to a user and an object.
type RoleBinding struct {
	ID         string
	OwnerID    string
	UserID     string
	ObjectID   string
	ObjectType ObjectType
	Role       access.Role
}

// ObjectType represents the type of object.
type ObjectType int

// ObjectTypeModel is the model object type.
const ObjectTypeModel ObjectType = iota

// AllObjectTypes is a list of all object types. Used for testing purposes to validate that all
// enum values are covered.
//
//nolint:gochecknoglobals
var AllObjectTypes = []ObjectType{
	ObjectTypeModel,
}

// String returns the string representation of the object type.
func (o ObjectType) String() string {
	if o == ObjectTypeModel {
		return "model"
	}

	return "ObjectType(" + strconv.Itoa(int(o)) + ")"
}

// IsSupportedObjectType checks if the given object type is supported.
func IsSupportedObjectType(objectType ObjectType) bool {
	for _, ot := range AllObjectTypes {
		if ot == objectType {
			return true
		}
	}

	return false
}
