package config

import (
	"errors"
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
	host = getEnvValue("HOST")
	if host == "" {
		panic(errors.New("can't get Host env from .env file"))
	}

	return
}

func GetPort() (port string) {
	port = getEnvValue("PORT")
	if port == "" {
		panic(errors.New("can't get Port env from .env file"))
	}
	return
}
