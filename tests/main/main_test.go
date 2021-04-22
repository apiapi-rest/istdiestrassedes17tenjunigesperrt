package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/availability"
)

func TestHttp(t *testing.T) {
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	// Start HTTP server.
	log.Printf("listening on port %s", port)

	url := fmt.Sprintf("http://localhost:%s", port)

	log.Print(url)

	req, err := http.Get(url + "/availability")

	if err != nil {
		t.Error(err)
	}

	if req.StatusCode != http.StatusOK {
		t.Fatalf("Expected %q, got %q.", http.StatusText(http.StatusOK), http.StatusText(req.StatusCode))

	}
	var resp availability.Response

	if err := json.NewDecoder(req.Body).Decode(&resp); err != nil {
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
