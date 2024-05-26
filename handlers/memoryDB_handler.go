package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/code-raushan/go-mongodb-server/services"
	"github.com/code-raushan/go-mongodb-server/types"
)

func InMemoryDBHandler(s *services.InMemoryDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			key := r.URL.Query()["key"][0]

			response, err := s.Get(key)

			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

		case http.MethodPost:
			body := r.Body
			var params *types.MemoryDataStore

			if err := json.NewDecoder(body).Decode(&params); err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			}

			response, err := s.Post(params)
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(response)
			
		default:
			http.Error(w, "Invalid request method", http.StatusBadRequest)
		}
	}
}
