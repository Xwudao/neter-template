// Package errs defines business errors that can cross the biz/repository
// boundary without exposing storage or transport details.
package errs

import "fmt"

// Kind describes how a caller can safely present a business error.
type Kind string

const (
	Invalid     Kind = "invalid"
	NotFound    Kind = "not_found"
	Conflict    Kind = "conflict"
	Unavailable Kind = "unavailable"
)

// Error carries a stable, user-facing message and an optional underlying
// cause. The cause remains available to logs and errors.Is/errors.As.
type Error struct {
	Kind    Kind
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.Err == nil {
		return e.Message
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *Error) Unwrap() error { return e.Err }

func New(kind Kind, message string, err error) *Error {
	return &Error{Kind: kind, Message: message, Err: err}
}

func NewInvalid(message string, err error) *Error { return New(Invalid, message, err) }

func NewNotFound(message string, err error) *Error { return New(NotFound, message, err) }

func NewConflict(message string, err error) *Error { return New(Conflict, message, err) }

func NewUnavailable(message string, err error) *Error { return New(Unavailable, message, err) }
