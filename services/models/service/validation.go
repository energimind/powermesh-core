package service

import "github.com/energimind/powermesh-core/errorz"

func requireString(value, name string) error {
	if value == "" {
		return errorz.NewValidationError(name + " is required")
	}

	return nil
}
