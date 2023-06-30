package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.WriteHeader(http.StatusOK)
    w.Write([]byte(`ok`))
}

func main() {
    s := &server{}
    http.Handle("/health", s)
    log.Fatal(http.ListenAndServe(":8080", nil))
}