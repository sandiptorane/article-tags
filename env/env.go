package env

import (
	"os"
)

const (
	Environment = "ENVIRONMENT"
	Local       = "local"
)

func Get(key string) string {
	return os.Getenv(key)
}
