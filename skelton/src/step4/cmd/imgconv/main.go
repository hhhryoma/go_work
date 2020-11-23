package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

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

func main() {
	encodeImage(os.Args[1], os.Args[2])
}
