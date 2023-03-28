package application

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (server *server) LoadRoutes() {
	routesGroup := server.App.Group("/qrcode")

	server.App.Get("/metrics", monitor.New(monitor.Config{
		Title:   "QR Code Generator API Metrics",
		Refresh: time.Second * 5,
	}))

	routesGroup.Get("/", Controllers.Welcome)
	routesGroup.Get("/:text", Controllers.EncodeText)
	routesGroup.Get("/:text/:size", Controllers.EncodeTextWithSize)
}
