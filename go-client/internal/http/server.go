package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Serve(port string) error {
	s.mux.HandleFunc("/ping", s.PingPong)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.mux)
}
