package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	ServerPort string
}

var Env env

func (*env) Load() error {
	err := godotenv.Load()

	if err == nil {
		log.Println(
			"environmental variables loaded from .env file",
		)
	}

	Env.ServerPort = os.Getenv("SERVER_PORT")

	if Env.ServerPort == "" {
		Env.ServerPort = "3000"
	}

	return nil
}
