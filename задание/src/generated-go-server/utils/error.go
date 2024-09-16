package openapi

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrTypeAssertionError = errors.New("unable to assert type")

type ParsingError struct {
	Param string
	Err   error
}

func (e *ParsingError) Unwrap() error {
	return e.Err
}

func (e *ParsingError) Error() string {
	if e.Param == "" {
		return e.Err.Error()
	}
	return e.Param + ": " + e.Err.Error()
}

type RequiredError struct {
	Field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("required field '%s' is zero value.", e.Field)
}

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return e.Field
}

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error, result *ImplResponse)

func DefaultErrorHandler(w http.ResponseWriter, _ *http.Request, err error, result *ImplResponse) {
	var parsingErr *ParsingError
	if ok := errors.As(err, &parsingErr); ok {
		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusBadRequest), w)
		return
	}

	var requiredErr *RequiredError
	if ok := errors.As(err, &requiredErr); ok {
		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusUnprocessableEntity), w)
		return
	}

	var validErr *ValidationError
	if ok := errors.As(err, &validErr); ok {
		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusBadRequest), w)
		return
	}

	if result.Code == http.StatusBadRequest {
		_ = EncodeJSONResponse(result.Body, func(i int) *int { return &i }(http.StatusBadRequest), w)
		return
	}

	if result.Code == http.StatusNotFound {
		_ = EncodeJSONResponse(result.Body, func(i int) *int { return &i }(http.StatusNotFound), w)
		return
	}

	_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusInternalServerError), w)
	return
}

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

var (
	ErrSQLQuery       = errors.New("sql query error")
	ErrEmptyResultSet = errors.New("empty result set")
)

var (
	ErrNoUser             = errors.New("no user found")
	ErrUserNoRightsTender = errors.New("user does not have permission for this tender")
	ErrUserNoRightsBid    = errors.New("user does not have permission for this bid")
)

var (
	ErrNoOrganization    = errors.New("no organization found")
	ErrOrgNoRightsTender = errors.New("organization does not belong to tender")
)
