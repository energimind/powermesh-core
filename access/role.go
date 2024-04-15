package access

import "strconv"

// Role represents the role of a user.
// The role is used to determine the access level of the user.
type Role int

// Role enumeration.
const (
	RoleNone Role = iota
	RoleAdmin
	RoleCreator
	RoleEditor
	RoleGuest
)

// AllRoles is a list of all roles. Used for testing purposes to validate that all
// enum values are covered.
//
//nolint:gochecknoglobals
var AllRoles = []Role{
	RoleNone,
	RoleAdmin,
	RoleCreator,
	RoleEditor,
	RoleGuest,
}

// Has checks if the current role contains the given role.
// Use this method instead of == to allow for future upgrades to combined roles
// using bitwise operations.
func (r Role) Has(role Role) bool {
	return r&role == role
}

// String returns the string representation of the role.
func (r Role) String() string {
	switch r {
	case RoleNone:
		return "none"
	case RoleAdmin:
		return "admin"
	case RoleCreator:
		return "creator"
	case RoleEditor:
		return "editor"
	case RoleGuest:
		return "guest"
	}

	return "Role(" + strconv.Itoa(int(r)) + ")"
}

// IsSupportedRole checks if the given role is supported.
func IsSupportedRole(role Role) bool {
	for _, r := range AllRoles {
		if r == role {
			return true
		}
	}

	return false
}
