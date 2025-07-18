package cfg

import (
	"os"

	l "github.com/Neurasita/rest-api/pkg/logger"
)

var (
	APP_HOST   string
	APP_PORT   string
	APP_ID     string
	APP_SECRET string

	DATABASE_URL string
)

func init() {
	APP_HOST = getEnv("APP_HOST", "0.0.0.0")
	APP_PORT = mustGetEnv("APP_PORT")
	APP_ID = mustGetEnv("APP_ID")
	APP_SECRET = mustGetEnv("APP_SECRET")

	DATABASE_URL = mustGetEnv("DATABASE_URL")
}

func getEnv(key, fb string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fb
}

func mustGetEnv(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	l.Error.Fatalf("unable to get %s environment variable", key)
	panic("environment variable was not set")
}
