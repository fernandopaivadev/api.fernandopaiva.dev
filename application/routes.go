package application

func (server *server) LoadRoutes() {
	routesGroup := server.App.Group("/qrcode")
	routesGroup.Get("/", Controllers.Welcome)
	routesGroup.Get("/:text", Controllers.EncodeText)
	routesGroup.Get("/:text/:size", Controllers.EncodeTextWithSize)
}
