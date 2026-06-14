package foohandler

import (
	"net/http"
)

func handleGetFoo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// w.Header().Set("X-Request-ID", "abc123")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Foo!"))
}
