package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/code-raushan/go-mongodb-server/services"
	"github.com/code-raushan/go-mongodb-server/utils"
)

func FetchRecordsHandler(s *services.MongoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Fatal("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		filters, err := utils.ExtractFilters(r.URL.Query())
		if err != nil {
			log.Fatalf("Error while filtering the response %v ", err)
		}

		res := s.FetchRecords(*filters)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
