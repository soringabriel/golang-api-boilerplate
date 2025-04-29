package tests

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"api/app/endpoints"
	"api/app/endpoints/user_endpoints"
	"api/databases"
	"api/helpers"
	"api/logger"

	"github.com/joho/godotenv"
)

var BASE_REQUEST_HEADERS = map[string]string{}

func FailTest(t *testing.T, message string, attempt int, maxRetries int) {
	if attempt >= maxRetries {
		t.Fatal(message)
	} else {
		t.Log(message)
	}
}

func TestMain(m *testing.M) {
	// Get env variables
	err := godotenv.Load("../.env-tests")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup logger
	logger.SetupLogger()

	// Start MongoDB containers
	uriMongoDB, cleanupMongoDB, err := helpers.SetupMongoDBDockerContainer()
	if err != nil {
		logger.Instance.Fatal("Failed to start MongoDB container", err)
	}
	helpers.SetEnvVariable("MONGODB_URL", uriMongoDB)
	defer cleanupMongoDB()

	// Setup databases
	databases.SetupMongodbDatabase()

	// Setup endpoints
	endpoints.SetupUniversalMiddlewares()
	http.HandleFunc(user_endpoints.CreateUserEndpoint.Path, user_endpoints.CreateUserEndpoint.ServeHTTP)
	http.HandleFunc(user_endpoints.ReadUserEndpoint.Path, user_endpoints.ReadUserEndpoint.ServeHTTP)
	http.HandleFunc(user_endpoints.ReadLimitedUserEndpoint.Path, user_endpoints.ReadLimitedUserEndpoint.ServeHTTP)
	http.HandleFunc(user_endpoints.UpdateUserEndpoint.Path, user_endpoints.UpdateUserEndpoint.ServeHTTP)
	http.HandleFunc(user_endpoints.DeleteUserEndpoint.Path, user_endpoints.DeleteUserEndpoint.ServeHTTP)

	// Start server
	go func() {
		logger.Instance.Info("Start server on " + helpers.GetEnvVariable("API_IP_PORT"))
		err = http.ListenAndServe(helpers.GetEnvVariable("API_IP_PORT"), nil)
		if err != nil {
			logger.Instance.Fatal("Failed to start server", err)
		}
	}()

	// Run the tests
	time.Sleep(time.Second)
	exitCode := m.Run()

	// Exit with the appropriate status code
	os.Exit(exitCode)
}
