package http

import "github.com/gorilla/mux"

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/userGet", s.Arif.ArifHandler).Methods("GET")
	r.HandleFunc("/userInsert", s.Arif.ArifHandler).Methods("POST")
	r.HandleFunc("/userUpdate", s.Arif.ArifHandler).Methods("PUT")
	r.HandleFunc("/userDelete", s.Arif.ArifHandler).Methods("DELETE")
	return r
}
