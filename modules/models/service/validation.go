package service

import "github.com/energimind/powermesh-core/errorz"

func requireString(value, name string) error {
	if value == "" {
		return errorz.NewValidationError("%s is required", name)
	}

	return nil
}
