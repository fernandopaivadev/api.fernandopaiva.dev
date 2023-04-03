package application

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type controllers struct{}

var Controllers controllers

func (*controllers) Welcome(context echo.Context) error {
	return context.String(
		200,
		"QR Code Generator API v1.2.0\n\nUsage: \n  [...]/qrcode/:text (default 512px)\n  [...]/qrcode/:text/:size",
	)
}

func (*controllers) EncodeText(context echo.Context) error {
	textToBeEncoded := context.Param("text")

	qrcodePng, err := textToPng(textToBeEncoded, 512)

	if err != nil {
		return context.String(
			500,
			"Error encoding text => "+err.Error(),
		)
	}

	return context.Blob(200, "image/png", qrcodePng)
}

func (controllers *controllers) EncodeTextWithSize(context echo.Context) error {
	textToBeEncoded := context.Param("text")
	qrcodeImageSize, _ := strconv.ParseUint(context.Param("size"), 10, 16)

	qrcodePng, err := textToPng(textToBeEncoded, qrcodeImageSize)

	if err != nil {
		return context.String(
			500,
			"Error encoding text => "+err.Error(),
		)
	}

	return context.Blob(200, "image/png", qrcodePng)
}
