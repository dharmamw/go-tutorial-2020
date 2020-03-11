package http

import "github.com/gorilla/mux"

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", s.User.UserHandler).Methods("GET")
	r.HandleFunc("/usersInsert", s.User.UserHandler).Methods("POST")
	r.HandleFunc("/usersUpdate", s.User.UserHandler).Methods("PUT")
	r.HandleFunc("/usersDelete", s.User.UserHandler).Methods("DELETE")

	return r
}


