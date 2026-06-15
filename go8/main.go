package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/user", makeHTTPHandler(handleGetUserID))
	http.ListenAndServe(":8080", nil)
}
