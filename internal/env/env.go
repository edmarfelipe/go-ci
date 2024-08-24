package env

import (
	"os"
)

type Env struct {
	ServiceAddr string
}

// Load loads the environment variables from the system
func Load() (*Env, error) {
	return &Env{
		ServiceAddr: getEnv("SERVICE_PORT", ":8080"),
	}, nil
}

// getEnv returns the value of an environment variable, or a fallback value if it's not set.
func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
