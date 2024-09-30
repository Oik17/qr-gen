package main

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	//var png []byte
	//png, err := qrcode.Encode("https://github.com/Oik17", qrcode.Medium, 256)
	err := qrcode.WriteFile("https://github.com/Oik17", qrcode.Medium, 256, "qr.png")

	if err != nil {
		log.Println(err)
	}
}
