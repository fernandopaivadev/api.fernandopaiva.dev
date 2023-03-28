package application

import (
	"errors"

	"github.com/skip2/go-qrcode"
)

func encodeQrCode(text string, size uint64) ([]byte, error) {
	if size > 5000 {
		return nil, errors.New("image size must be less or equal to 5000")
	}

	return qrcode.Encode(text, qrcode.Highest, int(size))
}
