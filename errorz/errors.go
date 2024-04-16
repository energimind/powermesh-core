package errorz

import (
	"errors"
	"fmt"
)

// BadRequestError is the error returned when a request is invalid.
type BadRequestError struct {
	Message string
}

// NewBadRequestError returns a new BadRequestError.
func NewBadRequestError(format string, args ...any) BadRequestError {
	return BadRequestError{
		Message: fmt.Sprintf(format, args...),
	}
}

// IsBadRequestError returns true if the error is a BadRequestError.
func IsBadRequestError(err error) bool {
	var badRequestError BadRequestError

	return errors.As(err, &badRequestError)
}

// Error returns the error message.
func (e BadRequestError) Error() string {
	return e.Message
}

// NotFoundError is the error returned when an object is not found.
type NotFoundError struct {
	Message string
}

// NewNotFoundError returns a new NotFoundError.
func NewNotFoundError(format string, args ...any) NotFoundError {
	return NotFoundError{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns the error message.
func (e NotFoundError) Error() string {
	return e.Message
}

// IsNotFoundError returns true if the error is a NotFoundError.
func IsNotFoundError(err error) bool {
	var notFoundError NotFoundError

	return errors.As(err, &notFoundError)
}

// AccessDeniedError is the error returned when a user is not allowed to perform an action.
type AccessDeniedError struct {
	Message string
}

// NewAccessDeniedError returns a new AccessDeniedError.
func NewAccessDeniedError(format string, args ...any) AccessDeniedError {
	return AccessDeniedError{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns the error message.
func (e AccessDeniedError) Error() string {
	return e.Message
}

// IsAccessDeniedError returns true if the error is an AccessDeniedError.
func IsAccessDeniedError(err error) bool {
	var accessDeniedError AccessDeniedError

	return errors.As(err, &accessDeniedError)
}

// UnauthorizedError is the error returned when a user is not authorized to perform an action.
type UnauthorizedError struct {
	Message string
}

// NewUnauthorizedError returns a new UnauthorizedError.
func NewUnauthorizedError(format string, args ...any) UnauthorizedError {
	return UnauthorizedError{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns the error message.
func (e UnauthorizedError) Error() string {
	return e.Message
}

// IsUnauthorizedError returns true if the error is an UnauthorizedError.
func IsUnauthorizedError(err error) bool {
	var unauthorizedError UnauthorizedError

	return errors.As(err, &unauthorizedError)
}

// ValidationError is the error returned when an object is invalid.
type ValidationError struct {
	Message string
}

// NewValidationError returns a new ValidationError.
func NewValidationError(format string, args ...any) ValidationError {
	return ValidationError{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns the error message.
func (e ValidationError) Error() string {
	return e.Message
}

// IsValidationError returns true if the error is a ValidationError.
func IsValidationError(err error) bool {
	var validationError ValidationError

	return errors.As(err, &validationError)
}

// StoreError is the error returned when an error occurs while storing an object.
type StoreError struct {
	Message string
}

// NewStoreError returns a new StoreError.
func NewStoreError(format string, args ...any) StoreError {
	return StoreError{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns the error message.
func (e StoreError) Error() string {
	return e.Message
}

// IsStoreError returns true if the error is a StoreError.
func IsStoreError(err error) bool {
	var storeError StoreError

	return errors.As(err, &storeError)
}

// GatewayError is the error returned when an error occurs while communicating with
// an external service.
type GatewayError struct {
	Message string
}

// NewGatewayError returns a new GatewayError.
func NewGatewayError(format string, args ...any) GatewayError {
	return GatewayError{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns the error message.
func (e GatewayError) Error() string {
	return e.Message
}

// IsGatewayError returns true if the error is a GatewayError.
func IsGatewayError(err error) bool {
	var gatewayError GatewayError

	return errors.As(err, &gatewayError)
}

// InternalError is the error returned when an internal error occurs.
type InternalError struct {
	Message string
}

// NewInternalError returns a new InternalError.
func NewInternalError(format string, args ...any) InternalError {
	return InternalError{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns the error message.
func (e InternalError) Error() string {
	return e.Message
}

// IsInternalError returns true if the error is an InternalError.
func IsInternalError(err error) bool {
	var internalError InternalError

	return errors.As(err, &internalError)
}
