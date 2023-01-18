package main

import "log"

func main() {
	Environment.Load()

	Server.Setup()
	Server.LoadRoutes()

	err := Server.Start()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
