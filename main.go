package main

import (
	"./zhg_qrcode/decode"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
)

func main() {
	var img image.Image
	var err error
	var body string
	file, err := os.Open("qr.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()

	img, _ = png.Decode(file)

	newImg := decode.NewImage(img)
	scanner := decode.NewScanner().SetEnabledAll(true)

	symbols, _ := scanner.ScanImage(newImg)
	for _, s := range symbols {
		body += s.Data
	}
	fmt.Println("识别结果:", body)
	time.Sleep(time.Second * 1000)
}
