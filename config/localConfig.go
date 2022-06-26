package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func CheckEnvValues() {
	GetHost()
	GetPort()
}

// getEnvValue will get env values from .env file.
func getEnvValue(key string) string {
	if RunMode == "normal" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	} else if RunMode == "test" {
		path, _ := os.Getwd()
		err := godotenv.Load(strings.Split(path, ProjectName)[0] + ProjectName + "/.env.test")
		if err != nil {
			log.Fatalf("Error loading .env.test file")
		}
	}

	return os.Getenv(key)
}

func GetHost() (host string) {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == "HOST" {
			host = pair[1]
			return
		}
	}

	host = getEnvValue("HOST")
	return
}

func GetPort() (port string) {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == "PORT" {
			port = pair[1]
			return
		}

	}
	port = "12345" //getEnvValue("PORT")
	return
}
