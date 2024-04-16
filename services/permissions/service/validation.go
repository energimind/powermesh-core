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

func validateObjectID(objectID string) error {
	if objectID == "" {
		return errorz.NewValidationError("objectID is required")
	}

	return nil
}

func validateObjectType(objectType permissions.ObjectType) error {
	if !permissions.IsSupportedObjectType(objectType) {
		return errorz.NewValidationError("objectType %s is not supported", objectType)
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

	if err := validateObjectID(data.ObjectID); err != nil {
		return err
	}

	if err := validateObjectType(data.ObjectType); err != nil {
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

	if err := validateObjectID(query.ObjectID); err != nil {
		return err
	}

	if err := validateObjectType(query.ObjectType); err != nil {
		return err
	}

	return nil
}

func validateAccessibleObjectsQuery(query permissions.AccessibleObjectsQuery) error {
	if err := validateUserID(query.UserID); err != nil {
		return err
	}

	if err := validateObjectType(query.ObjectType); err != nil {
		return err
	}

	return nil
}
