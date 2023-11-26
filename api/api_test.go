// Package api_test
package api

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTest(tb testing.TB) func(tb testing.TB) {
	a := Server{}
	server := a.NewServer(":8080")
	server.Start()

	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

func TestHelloWorld(t *testing.T) {
	//setup := setupTest(t)
	//defer setup(t)

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	// API GET
	HelloWorld(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Expected error to be nil got %v", err)
	}
	if string(data) != "Hello World" {
		t.Errorf("Expected 'Hello World' got %v", string(data))
	}
}
