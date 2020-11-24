package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

var clip string

func init() {
	flag.StringVar(&clip, "clip", "", "切り取る画像サイズ（'幅[px|%]x高さ[px|%]'）")
	flag.Parse()
}

func encodeImage(sourceFileName string, fileName string) {
	out, _ := os.Create(fileName)
	defer out.Close()

	source, _ := os.Open(sourceFileName)
	defer source.Close()

	img, _, _ := image.Decode(source)

	ext := strings.Split(fileName, ".")[1]
	if ext == "png" {
		png.Encode(out, img)
	} else if ext == "jpg" {
		jpeg.Encode(out, img, nil)
	}
}

func getFlag() string {
	if clip != "" {
		fmt.Println("切り抜きを行う予定", clip)
	}
	return flag.Args()[1]
}

func main() {
	getFlag()
}
