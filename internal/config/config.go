package config

import (
	"movie-service/internal/env"
	"os"
	"strconv"
)

type Config struct {
	Port int
	Env  string
}

func New() *Config {
	cfg := &Config{}
	cfg.Port = getFromEnv("PORT", 8080).(int)
	cfg.Env = getFromEnv("APP_ENV", "local").(string)
	return cfg
}
func getFromEnv(key string, defaultValue interface{}) interface{} {
	env.LoadEnv()
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	if v, err := strconv.Atoi(val); err == nil {
		return v
	}
	return val

}
