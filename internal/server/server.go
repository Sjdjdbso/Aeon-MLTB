package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	if port == "" {
		port = "8080"
	}
	return &Server{port: port}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleHealth)
	mux.HandleFunc("/health", s.handleHealth)
	return mux
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func (s *Server) Start() error {
	addr := ":" + s.port
	return http.ListenAndServe(addr, s.Routes())
}
