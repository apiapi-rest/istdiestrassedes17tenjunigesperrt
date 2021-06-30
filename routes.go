package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleRoot())
	s.router.HandleFunc("/availability", s.handleAvailability())
	s.router.HandleFunc("/availability/", s.handleAvailability())

}
