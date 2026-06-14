package foohandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetFooResponseRecorder(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/foo", nil)

	if err != nil {
		t.Error(err)
	}

	handleGetFoo(rr, req)

	// Check the response status code
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Result().StatusCode)
	}

	defer rr.Result().Body.Close()

	expected := "Foo!"
	b, err := io.ReadAll(rr.Result().Body)

	if err != nil {
		t.Error(err)
	}

	// Check the response body
	if string(b) != expected {
		t.Errorf("Expected response body %q, got %q", expected, string(b))
	}
}

func TestHandleGetFoo(t *testing.T) {
	// Create a test server using the handleGetFoo function
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	expected := "Foo!"
	b, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	// Check the response body
	if string(b) != expected {
		t.Errorf("Expected response body %q, got %q", expected, string(b))
	}
}
