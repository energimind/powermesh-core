package service

import (
	"net/mail"

	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/users"
)

func validateID(id string) error {
	if id == "" {
		return errorz.NewValidationError("id is required")
	}

	return nil
}

func validateUsername(username string) error {
	if username == "" {
		return errorz.NewValidationError("username is required")
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return errorz.NewValidationError("email is required")
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
