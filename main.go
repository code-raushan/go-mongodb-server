package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/code-raushan/go-mongodb-server/config"
	"github.com/code-raushan/go-mongodb-server/handlers"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	uri, available := os.LookupEnv("MONGO_URI")
	if !available {
		fmt.Println("uri is not available")
		return
	}

	config.ConnectDB(uri)

	http.HandleFunc("/health", handlers.HealthHandler())

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Error in the http server %v", err)
	}
}
