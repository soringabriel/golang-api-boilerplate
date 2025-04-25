package main

import (
	"log"
	"net/http"

	"api/app/endpoints/user_endpoints"
	"api/databases"
	"api/helpers"
	"api/logger"

	"github.com/joho/godotenv"
)

func main() {
	// Get env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup logger
	logger.SetupLogger()

	// Setup databases
	databases.SetupMongodbDatabase()

	// Setup endpoints
	http.HandleFunc(user_endpoints.CreateUserEndpoint.Path, user_endpoints.CreateUserEndpoint.ServeHTTP)
	http.HandleFunc(user_endpoints.ReadUserEndpoint.Path, user_endpoints.ReadUserEndpoint.ServeHTTP)
	http.HandleFunc(user_endpoints.UpdateUserEndpoint.Path, user_endpoints.UpdateUserEndpoint.ServeHTTP)
	http.HandleFunc(user_endpoints.DeleteUserEndpoint.Path, user_endpoints.DeleteUserEndpoint.ServeHTTP)

	// Start application
	logger.Instance.Info("Start server on " + helpers.EnvVariable("API_IP_PORT"))
	err = http.ListenAndServe(helpers.EnvVariable("API_IP_PORT"), nil)
	if err != nil {
		logger.Instance.Fatal("Failed to start server", err)
	}
}
