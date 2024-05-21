package main

import (
	"fmt"
	"log"
	"os"

	"github.com/code-raushan/go-mongodb-server/config"

	"github.com/joho/godotenv"
)

func main() {

	if err:=godotenv.Load();  err != nil {
		log.Fatalf("Error loading .env file")
	}

	uri, available := os.LookupEnv("MONGO_URI")
	if !available {
		fmt.Println("uri is not available")
		return
	}

	config.ConnectDB(uri)
}