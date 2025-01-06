
package handlers

import (
	"Assignment2_TomirisTapen/server"
	"encoding/json"
	"net/http"
)

func DataHandler(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var input map[string]string
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			s.Mu.Lock()
			for key, value := range input {
				s.Data[key] = value
			}
			s.Requests++
			s.Mu.Unlock()
			w.WriteHeader(http.StatusCreated)
		case http.MethodGet:
			s.Mu.Lock()
			defer s.Mu.Unlock()
			s.Requests++
			json.NewEncoder(w).Encode(s.Data)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func DeleteDataHandler(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[len("/data/"):]
		s.Mu.Lock()
		delete(s.Data, key)
		s.Requests++
		s.Mu.Unlock()
		w.WriteHeader(http.StatusOK)
	}
}
