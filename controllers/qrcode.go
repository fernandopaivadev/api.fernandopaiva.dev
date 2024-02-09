package controllers

import (
	"main/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type qrCodeController struct{}

var QRCode qrCodeController

func (*qrCodeController) EncodeText(c echo.Context) error {
	textToBeEncoded := c.Param("text")
	qrcodeImageSize, _ := strconv.ParseUint(c.Param("size"), 10, 16)

	if qrcodeImageSize == 0 {
		qrcodeImageSize = 500
	}

	qrcodePng, err := services.QRCode.TextToPng(textToBeEncoded, qrcodeImageSize)

	if err != nil {
		return c.String(
			500,
			"Error encoding text: "+err.Error(),
		)
	}

	return c.Blob(200, "image/png", qrcodePng)
}
