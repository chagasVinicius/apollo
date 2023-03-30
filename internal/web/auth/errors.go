package auth

import (
	"errors"
	"fmt"
)

// Error is used to pass an error during the request through the
// application with auth specific context.
type Error struct {
	msg string
}

// NewAuthError creates an AuthError for the provided message.
func NewAuthError(format string, args ...any) error {
	return &Error{
		msg: fmt.Sprintf(format, args...),
	}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (ae *Error) Error() string {
	return ae.msg
}

// IsAuthError checks if an error of type AuthError exists.
func IsAuthError(err error) bool {
	var ae *Error
	return errors.As(err, &ae)
}
