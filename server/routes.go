package server

import "main/controllers"

func (server *httpServer) LoadQRCodeRoutes() {
	server.instance.GET(
		"/qrcode/:text",
		controllers.QRCode.EncodeText,
	)
	server.instance.GET(
		"/qrcode/:text/:size",
		controllers.QRCode.EncodeText,
	)
}
