package service

import (
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/services/models"
)

func validateID(id string) error {
	if id == "" {
		return errorz.NewValidationError("id is required")
	}

	return nil
}

func validateCode(code string) error {
	if code == "" {
		return errorz.NewValidationError("code is required")
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return errorz.NewValidationError("name is required")
	}

	return nil
}

func validateModelData(data models.ModelData) error {
	if err := validateCode(data.Code); err != nil {
		return err
	}

	if err := validateName(data.Name); err != nil {
		return err
	}

	return nil
}
