package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/handler"
)

func main() {
	log.Print("starting server...")

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)

	s := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", port),
		Handler:      handler.Mux(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Wtf() {

}
