package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	} else {
		return fallback
	}
}

func GetInt(key string, fallback int) int {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return v
}
