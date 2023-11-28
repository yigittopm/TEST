// Package api_test
package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHelloWorld(t *testing.T) {
	router := SetupServer()

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	// API GET
	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	expected := []byte(`{"message":"Hello World!"}`)

	if err != nil {
		t.Errorf("Expected error to be nil got %v", err)
	}
	if string(data) != string(expected) {
		t.Errorf("Expected %s got %s", expected, data)
	}
}
