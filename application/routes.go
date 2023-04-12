package application

func (server *server) LoadRoutes() {
	server.Instance.GET("/qrcode", Controllers.EncodeTextWithSize)
}
