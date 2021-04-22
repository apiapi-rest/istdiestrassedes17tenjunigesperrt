package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/availability"
	"apiapi.rest/istdiestrassedes17tenjunigesperrt/handler"
)

func TestMain(m *testing.M) {
	// fmt.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func setupAPI(t *testing.T) (string, func()) {
	t.Helper()

	ts := httptest.NewServer(handler.Mux())

	return ts.URL, func() {
		ts.Close()

	}
}

func TestHttp(t *testing.T) {
	url, cleanup := setupAPI(t)
	defer cleanup()

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
