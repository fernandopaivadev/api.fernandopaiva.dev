package main

func (server *server) LoadRoutes() {
	server.App.Get("/qrcode", Controllers.Welcome)
	server.App.Get("/qrcode/:text", Controllers.EncodeText)
	server.App.Get("/qrcode/:text/:size", Controllers.EncodeTextWithSize)
}
