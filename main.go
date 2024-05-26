package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/code-raushan/go-mongodb-server/config"
	"github.com/code-raushan/go-mongodb-server/handlers"
	"github.com/code-raushan/go-mongodb-server/repositories"
	"github.com/code-raushan/go-mongodb-server/services"

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

	dbClient := config.ConnectDB(uri)

	const (
		dbName   = "getircase-study"
		collName = "records"
	)

	m := repositories.NewMongoRepo(dbClient, dbName, collName)

	mongoService := services.NewMongoService(m)

	inMemoryService := services.NewInMemoryDB()

	http.HandleFunc("/fetch", handlers.FetchRecordsHandler(mongoService))

	http.HandleFunc("/health", handlers.HealthHandler())

	http.HandleFunc("/in-memory", handlers.InMemoryDBHandler(inMemoryService))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Error in the http server %v", err)
	}
}
