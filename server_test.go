package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestHandleAbout(t *testing.T) {
	is := is.New(t)
	srv := server{}
	srv.routes()
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.Code, http.StatusOK)

	req = httptest.NewRequest("GET", "/headers", nil)
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.Code, http.StatusOK)
}
