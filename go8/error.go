package main

import (
	"net/http"
)

type apiError struct {
	Err    string
	Status int
}

// Error implements the error interface for apiError.
func (e *apiError) Error() string {
	return e.Err
}

var internalServerError = &apiError{
	Err:    "internal server error",
	Status: http.StatusInternalServerError,
}

var methodNotAllowedError = &apiError{
	Err:    "method not allowed",
	Status: http.StatusMethodNotAllowed,
}

var unauthorizedError = &apiError{
	Err:    "invalid user",
	Status: http.StatusUnauthorized,
}
