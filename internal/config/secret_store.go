package config

import "os"

var secret_store = make(map[string]string)

func GetSecret(key string) string {
	if secret_store[key] != "" {
		return secret_store[key]
	}

	value := os.Getenv(key)
	secret_store[key] = value
	return value
}
