package http

import (
	"net/http"
	"tugas-arif/pkg/grace"

	"github.com/rs/cors"
)

// ArifHandler ...
type ArifHandler interface {
	// Masukkan fungsi handler di sini
	ArifHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server *http.Server
	Arif   ArifHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
