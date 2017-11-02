package main

import (
	"fmt"
	"io/ioutil"
	"os"

	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"

	"github.com/goforbroke1006/watermarksvc/config"
	"github.com/goforbroke1006/watermarksvc/strategy"
)

func openImage(filename string) (*image.Image, string) {
	inFile, err := os.Open(filename)
	doPanic(err)
	defer inFile.Close()

	sourceImage, mimeType, err := image.Decode(inFile)
	doPanic(err)

	return &sourceImage, mimeType
}

func saveFile(filename string, dist *image.RGBA, mimeType string) {
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

	stg := &strategy.UglySplitStrategy{Rows: 4}

	for _, f := range files {
		inFilename := cfg.InputDir + "/" + f.Name()
		sourceImage, mimeType := openImage(inFilename)
		fmt.Println(inFilename + " (" + mimeType + ")")

		dist := stg.AddWatermark(wmImage, sourceImage)

		saveFile(cfg.OutputDir+"/"+f.Name(), dist, mimeType)
	}
}

func doPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
