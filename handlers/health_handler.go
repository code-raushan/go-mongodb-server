package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HealthResponse struct {
	Health string `json:"health"`
}

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r * http.Request){
		if r.Method != http.MethodGet {
			log.Fatal("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		fmt.Println("/health route hit")
		w.Header().Set("Content-Type", "application/json")
		healthResponse := HealthResponse {
			Health: "OK",
		}
		
		json.NewEncoder(w).Encode(healthResponse)
	}
}