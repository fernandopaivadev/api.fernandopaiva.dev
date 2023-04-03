package application

func (server *server) LoadRoutes() {
	routesGroup := server.Instance.Group("/qrcode")

	routesGroup.GET("/", Controllers.Welcome)
	routesGroup.GET("/:text", Controllers.EncodeText)
	routesGroup.GET("/:text/:size", Controllers.EncodeTextWithSize)
}
