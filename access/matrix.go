package access

// Matrix is an access matrix.
// It maps roles to permissions.
type Matrix map[Role]Permissions

// NewMatrix creates a new access matrix.
func NewMatrix() Matrix {
	matrix := make(Matrix)

	matrix[RoleNone] = PermissionsNone
	matrix[RoleAdmin] = PermissionsAll
	matrix[RoleCreator] = Permissions(PermissionCreate | PermissionRead | PermissionWrite)
	matrix[RoleEditor] = Permissions(PermissionRead | PermissionWrite)
	matrix[RoleGuest] = Permissions(PermissionRead)

	return matrix
}

// GetPermissions returns the permissions for the given role.
func (m Matrix) GetPermissions(role Role) Permissions {
	return m[role]
}

// HasPermission checks if the given role has the given permission.
func (m Matrix) HasPermission(role Role, perm Permission) bool {
	return m[role].Has(perm)
}
