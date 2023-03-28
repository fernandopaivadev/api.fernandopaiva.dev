package application

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type controllers struct{}

var Controllers controllers

func (*controllers) Welcome(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).SendString(
		"Usage: \n  [...]/qrcode/:text (default 512px)\n  [...]/qrcode/:text/:size",
	)
}

func (*controllers) EncodeText(context *fiber.Ctx) error {
	textToBeEncoded := context.Params("text")

	qrcodePng, err := encodeQrCode(textToBeEncoded, 512)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(
			"Error parsing text to encode => " + err.Error(),
		)
	}

	context.Response().Header.Set("Content-Type", "image/png")
	return context.Status(fiber.StatusOK).Send(qrcodePng)
}

func (controllers *controllers) EncodeTextWithSize(context *fiber.Ctx) error {
	textToBeEncoded := context.Params("text")
	qrcodeImageSize, err := strconv.ParseUint(context.Params("size"), 10, 16)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(
			"Error parsing image size => " + err.Error(),
		)
	}

	qrcodePng, err := encodeQrCode(textToBeEncoded, qrcodeImageSize)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(
			"Error generating qr code image => " + err.Error(),
		)
	}

	context.Response().Header.Set("Content-Type", "image/png")
	return context.Status(fiber.StatusOK).Send(qrcodePng)
}
