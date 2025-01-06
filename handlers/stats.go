
package handlers

import (
	"Assignment2_TomirisTapen/server"
	"encoding/json"
	"net/http"
)

func StatsHandler(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Mu.Lock()
		defer s.Mu.Unlock()
		s.Requests++
		stats := map[string]int{
			"requests": s.Requests,
			"db_size":  len(s.Data),
		}
		json.NewEncoder(w).Encode(stats)
	}
}
