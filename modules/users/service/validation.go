package service

import (
	"net/mail"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/users"
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

func validateUsername(username string) error {
	return requireString(username, "username")
}

func validateEmail(email string) error {
	if err := requireString(email, "email"); err != nil {
		return err
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return errorz.NewValidationError("invalid email address")
	}

	return nil
}

func validateUserData(data users.UserData) error {
	if err := validateUsername(data.Username); err != nil {
		return err
	}

	if err := validateEmail(data.Email); err != nil {
		return err
	}

	return nil
}
