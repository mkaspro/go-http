package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type server struct {
	router http.ServeMux
	// Add server dependencies here
}

func newServer() *server {
	s := &server{}
	s.routes()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) HandleSomething() http.HandlerFunc {
	// Do setup stuff here
	return func(w http.ResponseWriter, r *http.Request) {
		// Do stuff
	}

}

func (s *server) HandleAnotherthing(someArg string) http.HandlerFunc {
	// Do setup stuff here
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s", someArg)
	}

}

func (s *server) AdminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}

func (s *server) Respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *server) Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *server) HandleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n") // Do stuff
	}
}

func (s *server) HandleHeaders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	}
}
