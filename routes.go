package main

func (s *server) routes() {
	s.router.HandleFunc("/hello", s.HandleHello())
	s.router.HandleFunc("/headers", s.HandleHeaders())
}
