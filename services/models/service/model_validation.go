package service

import (
	"github.com/energimind/powermesh-core/services/models"
)

func validateID(id string) error {
	return requireString(id, "id")
}

func validateCode(code string) error {
	return requireString(code, "code")
}

func validateName(name string) error {
	return requireString(name, "name")
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
