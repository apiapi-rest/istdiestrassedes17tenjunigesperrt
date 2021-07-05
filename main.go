package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	router *httprouter.Router
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

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), server.router); err != nil {
		return errors.Wrap(err, "ListenAndServe")
	}
	return nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() *server {
	s := &server{router: httprouter.New()}
	s.routes()
	return s
}
