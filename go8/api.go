package main

import (
	"encoding/json"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if apiErr, ok := err.(*apiError); ok {
				writeJSON(w, apiErr.Status, apiErr)
				return
			}

			writeJSON(w, http.StatusInternalServerError, internalServerError)
		}
	}
}

func handleGetUserID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return methodNotAllowedError
	}

	user := User{ID: "123", Valid: true}
	if !user.Valid {
		return unauthorizedError
	}

	return writeJSON(w, http.StatusOK, user)
}
