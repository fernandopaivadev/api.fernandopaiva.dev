package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envVars struct {
	ServerPort string
}

var Variables envVars

func Load() {
	err := godotenv.Load()

	if err == nil {
		log.Println(
			"Environment variables loaded from .env file",
		)
	}

	Variables.ServerPort = os.Getenv("SERVER_PORT")
}
