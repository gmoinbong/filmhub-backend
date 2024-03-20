package variables

import (
	"movie-service/internal/env"
	"os"
)

func getEnv(key string) string {
	env.LoadEnv()
	value := os.Getenv(key)
	if value == "" {
		Logger.Info("Environment variable %s not set", key)
	}
	return value
}
