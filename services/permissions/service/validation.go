package service

import (
	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/permissions"
)

func requireString(value, name string) error {
	if value == "" {
		return errorz.NewValidationError(name + " is required")
	}

	return nil
}

func validateID(id string) error {
	return requireString(id, "id")
}

func validateOwnerID(ownerID string) error {
	return requireString(ownerID, "owner id")
}

func validateUserID(userID string) error {
	return requireString(userID, "user id")
}

func validateResourceID(resourceID string) error {
	return requireString(resourceID, "resource id")
}

func validateResourceType(resourceType permissions.ResourceType) error {
	if !permissions.IsSupportedResourceType(resourceType) {
		return errorz.NewValidationError("resourceType %s is not supported", resourceType)
	}

	return nil
}

func validateRole(role access.Role) error {
	if !access.IsSupportedRole(role) {
		return errorz.NewValidationError("role %s is not supported", role)
	}

	return nil
}

func validateRoleBindingData(data permissions.RoleBindingData) error {
	if err := validateOwnerID(data.OwnerID); err != nil {
		return err
	}

	if err := validateUserID(data.UserID); err != nil {
		return err
	}

	if err := validateResourceID(data.ResourceID); err != nil {
		return err
	}

	if err := validateResourceType(data.ResourceType); err != nil {
		return err
	}

	if err := validateRole(data.Role); err != nil {
		return err
	}

	return nil
}

func validateRoleBindingQuery(query permissions.RoleBindingQuery) error {
	if err := validateUserID(query.UserID); err != nil {
		return err
	}

	if err := validateResourceID(query.ResourceID); err != nil {
		return err
	}

	if err := validateResourceType(query.ResourceType); err != nil {
		return err
	}

	return nil
}

func validateAccessibleResourcesQuery(query permissions.AccessibleResourcesQuery) error {
	if err := validateUserID(query.UserID); err != nil {
		return err
	}

	if err := validateResourceType(query.ResourceType); err != nil {
		return err
	}

	return nil
}
