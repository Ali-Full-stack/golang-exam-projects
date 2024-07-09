package code

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func GenerateQrcode(id string)[]byte{
	bytes, err := qrcode.Encode(id, qrcode.Medium, 256)
	if err != nil {
		log.Fatal("Failed to generate qrcode:",err)
	}
	return bytes
}	
