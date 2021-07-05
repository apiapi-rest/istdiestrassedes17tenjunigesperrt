package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) routes() {
	s.router.HandlerFunc("GET", "/", s.handleRoot())
	s.router.HandlerFunc("GET", "/availability", s.handleAvailability())

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
