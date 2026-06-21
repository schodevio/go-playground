package api

import (
	"encoding/json"
	"net/http"
	"sc/go13/storage"
	"sc/go13/util"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/users/{id}", s.handleGetUserByID)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(1)      // For simplicity, we are using a hardcoded ID here
	_ = util.Round2Dec(123.456) // Example usage of the utility function

	json.NewEncoder(w).Encode(user)
}
