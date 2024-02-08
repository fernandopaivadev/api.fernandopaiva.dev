package services

import (
	"errors"

	"github.com/skip2/go-qrcode"
)

type qrCodeService struct{}

var QRCode qrCodeService

func (*qrCodeService) TextToPng(text string, size uint64) ([]byte, error) {
	if size > 5000 {
		return nil, errors.New("image size must be less or equal to 5000")
	}

	if size == 0 {
		return nil, errors.New("image size must be greater than 0")
	}

	return qrcode.Encode(text, qrcode.Highest, int(size))
}
