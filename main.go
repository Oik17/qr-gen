package main

import (
	"image/color"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	//var png []byte
	//png, err := qrcode.Encode("https://github.com/Oik17", qrcode.Medium, 256)

	err := qrcode.WriteColorFile("https://github.com/Oik17", qrcode.Medium, 256, color.Black, color.White, "qr.png")

	if err != nil {
		log.Println(err)
	}
}
