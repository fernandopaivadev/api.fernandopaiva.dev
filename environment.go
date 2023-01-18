package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envVars struct {
	ServerPort string
}

type environment struct {
	Variables envVars
}

var Environment environment

func (environment *environment) Load() {
	err := godotenv.Load()

	if err == nil {
		log.Println(
			"Environment variables loaded from .env file",
		)
	}

	environment.Variables.ServerPort = os.Getenv("SERVER_PORT")
}
