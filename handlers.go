package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/availability"
)

func (s *server) handleRoot() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("HandleRoot")
		fmt.Fprintf(w, "Hello World!")
	}
}

func (s *server) handleAvailability() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("HandleAvailability")
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
}
