package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/availability"
)

func Mux() http.Handler {
	m := http.NewServeMux()

	m.HandleFunc("/", handleRoot)
	m.HandleFunc("/availability", handleAvailability)
	m.HandleFunc("/availability/", handleAvailability)

	return m
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleAvailability(w http.ResponseWriter, r *http.Request) {
	data, status := availability.AvailabilityResponse()

	json, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")
	w.WriteHeader(status)

	fmt.Fprint(w, string(json))
}
