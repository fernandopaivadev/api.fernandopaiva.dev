package server

import (
	"errors"
	"main/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

type httpServer struct {
	instance *echo.Echo
}

var HTTPServer httpServer

func (server *httpServer) Setup() {
	server.instance = echo.New()

	server.instance.Use(echomiddleware.RequestID())
	server.instance.Use(echomiddleware.Recover())
	server.instance.Use(echomiddleware.BodyLimit("2M"))
	server.instance.Use(echomiddleware.GzipWithConfig(echomiddleware.GzipConfig{
		Level: 5,
	}))
	server.instance.Use(echomiddleware.SecureWithConfig(echomiddleware.SecureConfig{
		XSSProtection:      "1; mode=block",
		XFrameOptions:      "deny",
		ContentTypeNosniff: "nosniff",
	}))
	server.instance.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodHead,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))
	server.instance.Use(middleware.Logger.CustomLogger())

	server.instance.HTTPErrorHandler = func(err error, c echo.Context) {
		c.String(500, "Error: "+err.Error())
	}
}

func (server *httpServer) Start(serverPort string) error {
	serverPort = ":" + serverPort

	err := server.instance.Start(serverPort)

	if err != nil {
		return errors.New("error starting server: " + err.Error())
	}

	return nil
}
