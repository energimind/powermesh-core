package service

import (
	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/permissions"
)

func validateID(id string) error {
	if id == "" {
		return errorz.NewValidationError("id is required")
	}

	return nil
}

func validateOwnerID(ownerID string) error {
	if ownerID == "" {
		return errorz.NewValidationError("ownerID is required")
	}

	return nil
}

func validateUserID(userID string) error {
	if userID == "" {
		return errorz.NewValidationError("userID is required")
	}

	return nil
}

func validateResourceID(resourceID string) error {
	if resourceID == "" {
		return errorz.NewValidationError("resourceID is required")
	}

	return nil
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
