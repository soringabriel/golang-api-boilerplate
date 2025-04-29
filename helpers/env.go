package helpers

import (
	"os"
)

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}

func SetEnvVariable(key string, value string) {
	os.Setenv(key, value)
}
