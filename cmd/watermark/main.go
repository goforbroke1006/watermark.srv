package main

import (
	"fmt"
	"io/ioutil"
	"os"

	_ "golang.org/x/image/webp"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	"image/png"

	"github.com/goforbroke1006/watermarksvc/config"
)

func openImage(filename string) (*image.Image, string) {
	inFile, err := os.Open(filename)
	doPanic(err)
	defer inFile.Close()

	sourceImage, mimeType, err := image.Decode(inFile)
	doPanic(err)

	return &sourceImage, mimeType
}

func addWatermark(watermark *image.Image, pic *image.Image) image.Image {
	r := image.Rectangle{image.Point{}, (*pic).Bounds().Size()}
	r2 := image.Rectangle{image.Point{}, (*watermark).Bounds().Size()}

	rgba := image.NewRGBA(r)

	draw.Draw(rgba, (*pic).Bounds(), *pic, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, *watermark, image.Point{0, 0}, draw.Over)

	return rgba
}

func saveFile(filename string, dist image.Image, mimeType string) {
	outfile, err := os.Create(filename)
	doPanic(err)
	defer outfile.Close()

	if "jpeg" == mimeType {
		jpeg.Encode(outfile, dist, nil)
	} else if "png" == mimeType {
		png.Encode(outfile, dist)
	}
}

func main() {
	cfg, _ := config.LoadConfig("./config.yml")

	wmImage, _ := openImage(cfg.WatermarkFile)

	files, err := ioutil.ReadDir(cfg.InputDir)
	doPanic(err)

	for _, f := range files {
		inFilename := cfg.InputDir + "/" + f.Name()
		sourceImage, mimeType := openImage(inFilename)
		fmt.Println(inFilename + " (" + mimeType + ")")

		dist := addWatermark(wmImage, sourceImage)

		saveFile(cfg.OutputDir+"/"+f.Name(), dist, mimeType)
	}
}

func doPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
