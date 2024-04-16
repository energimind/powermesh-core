package errorz

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	t.Parallel()

	tester := func(err, exp error) {
		require.IsType(t, exp, err)
		require.Equal(t, "test:42", err.Error())
	}

	tester(NewBadRequestError("test:%d", 42), BadRequestError{})
	tester(NewNotFoundError("test:%d", 42), NotFoundError{})
	tester(NewAccessDeniedError("test:%d", 42), AccessDeniedError{})
	tester(NewUnauthorizedError("test:%d", 42), UnauthorizedError{})
	tester(NewValidationError("test:%d", 42), ValidationError{})
	tester(NewStoreError("test:%d", 42), StoreError{})
	tester(NewGatewayError("test:%d", 42), GatewayError{})
	tester(NewInternalError("test:%d", 42), InternalError{})
}

func TestErrorIs(t *testing.T) {
	t.Parallel()

	tester := func(f func(err error) bool, exp error) {
		require.True(t, f(exp))
		require.False(t, f(errors.New("other-error")))
	}

	tester(IsBadRequestError, BadRequestError{})
	tester(IsNotFoundError, NotFoundError{})
	tester(IsAccessDeniedError, AccessDeniedError{})
	tester(IsUnauthorizedError, UnauthorizedError{})
	tester(IsValidationError, ValidationError{})
	tester(IsStoreError, StoreError{})
	tester(IsGatewayError, GatewayError{})
	tester(IsInternalError, InternalError{})
}
