package config

import "github.com/joho/godotenv"

func LoadEnv() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
}
