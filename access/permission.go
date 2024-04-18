package access

import (
	"strconv"
)

// Permission represents a permission to act on a resource.
type Permission int

// Permission enumeration for access control.
const (
	PermissionCreate Permission = 1 << iota
	PermissionRead
	PermissionWrite
	PermissionDelete
)

// AllPermissions is a list of all permissions. Used for testing purposes to validate that all
// enum values are covered.
//
//nolint:gochecknoglobals
var AllPermissions = []Permission{
	PermissionCreate,
	PermissionRead,
	PermissionWrite,
	PermissionDelete,
}

// String returns a string representation of the permission.
func (p Permission) String() string {
	switch p {
	case PermissionCreate:
		return "create"
	case PermissionRead:
		return "read"
	case PermissionWrite:
		return "write"
	case PermissionDelete:
		return "delete"
	default:
		return "Permission(" + strconv.Itoa(int(p)) + ")"
	}
}
