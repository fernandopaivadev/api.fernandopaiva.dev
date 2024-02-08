package main

import (
	"log"
	"main/config"
	"main/server"
)

func main() {
	err := config.Env.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}

	server.HTTPServer.Setup()
	server.HTTPServer.LoadQRCodeRoutes()

	err = server.HTTPServer.Start(config.Env.ServerPort)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
