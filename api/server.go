package api

import (
	"fmt"
	"log"
	"os"
	"sample/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

// Run is...
func Run() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting .env. Reason: %v", err)
	} else {
		fmt.Println("ENV values loaded")
	}


	server.Initalize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if environment variable is not set
	}

	serverRunning := server.Run(":" + port)

	if serverRunning != nil {
		log.Fatalln("Server failed to start")
	}
}
