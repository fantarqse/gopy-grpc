package http

import (
	"net/http"
)

func (s *Server) PingPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
