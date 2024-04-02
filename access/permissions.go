package access

import "strings"

// Permissions is a bit mask of permissions.
type Permissions Permission

// PermissionsNone represents no permissions.
const PermissionsNone = Permissions(0)

// PermissionsAll is a combination of all permissions.
const PermissionsAll = Permissions(
	PermissionCreate | PermissionRead | PermissionWrite | PermissionDelete)

// Has returns true if the permission contains the given permission.
func (p Permissions) Has(perm Permission) bool {
	return Permission(p)&perm == perm
}

// Add adds the given permission(s) to the permission.
func (p Permissions) Add(perms ...Permission) Permissions {
	for _, perm := range perms {
		p |= Permissions(perm)
	}

	return p
}

// Remove removes the given permission(s) from the permission.
func (p Permissions) Remove(perms ...Permission) Permissions {
	for _, perm := range perms {
		p &^= Permissions(perm)
	}

	return p
}

// String returns a string representation of the permission.
func (p Permissions) String() string {
	if p == PermissionsNone {
		return "none"
	}

	if p == PermissionsAll {
		return "all"
	}

	var b strings.Builder

	for _, perm := range AllPermissions {
		if !p.Has(perm) {
			continue
		}

		if b.Len() > 0 {
			b.WriteString(",")
		}

		b.WriteString(perm.String())
	}

	return b.String()
}
