package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

type server struct {
	router *http.ServeMux
	// db, etc.
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	log.Print("starting server...")

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)

	server := newServer()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      server.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		return errors.Wrap(err, "ListenAndServe")
	}
	return nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() *server {
	s := &server{router: http.NewServeMux()}
	s.routes()
	return s
}
