package http

import "github.com/gorilla/mux"

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", s.User.UserHandler).Methods("GET", "POST", "PUT", "DELETE")

	return r
}
