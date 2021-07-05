package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleRoot()).Methods("GET")
	s.router.HandleFunc("/availability", s.handleAvailability()).Methods("GET")

}
