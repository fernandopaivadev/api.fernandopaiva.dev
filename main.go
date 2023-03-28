package main

import (
	"log"
	"qr-code-generator-api/application"
)

func main() {
	application.Environment.Load()

	application.Server.Setup()
	application.Server.LoadRoutes()

	err := application.Server.Start()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
