package application

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type controllers struct{}

var Controllers controllers

func (controllers *controllers) EncodeTextWithSize(context echo.Context) error {
	textToBeEncoded := context.QueryParam("text")
	qrcodeImageSize, _ := strconv.ParseUint(context.QueryParam("size"), 10, 16)

	qrcodePng, err := textToPng(textToBeEncoded, qrcodeImageSize)

	if err != nil {
		return context.String(
			500,
			"Error encoding text => "+err.Error(),
		)
	}

	return context.Blob(200, "image/png", qrcodePng)
}
