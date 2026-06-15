package main

import (
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{svc: svc}
}

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.GetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) GetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}
