package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/availability"
)

func TestHandleRoot(t *testing.T) {
	srv := server{
		router: http.NewServeMux(),
	}
	srv.routes()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected Status Code %q, got %q.", http.StatusText(http.StatusOK), http.StatusText(w.Result().StatusCode))
	}

	expected := "Hello World!"
	if w.Body.String() != expected {
		t.Fatalf("Expected Body '%q', got '%q'.", expected, w.Body.String())
	}
}

func TestAvailability(t *testing.T) {
	srv := server{
		router: http.NewServeMux(),
	}
	srv.routes()
	req := httptest.NewRequest("GET", "/availability", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected Status Code %q, got %q.", http.StatusText(http.StatusOK), http.StatusText(w.Result().StatusCode))
	}

	var resp availability.Response

	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	req.Body.Close()

	if len(resp.Error) > 0 {
		t.Errorf("Expected no error. Got: %s", resp.Error)
	}

	if resp.Success != true {
		t.Errorf("Expected Response to be true. Got: %t", resp.Success)
	}

}
