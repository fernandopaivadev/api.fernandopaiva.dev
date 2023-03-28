package main

import (
	"log"
	"qr-code-generator-api/application"
	"qr-code-generator-api/environment"
)

func main() {
	environment.Load()

	application.Server.Setup()
	application.Server.LoadRoutes()

	err := application.Server.Start()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
