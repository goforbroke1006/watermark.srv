package main

import (
	"fmt"
	osUser "os/user"

	"image"
	"image/png"
	"image/jpeg"
	"io/ioutil"

	"github.com/goforbroke1006/watermarksvc/config"
	"os"

	_ "image/gif"
	_ "golang.org/x/image/webp"
	"image/draw"
)

func openImage(dirPath string, f os.FileInfo) (string, string, *image.Image) {
	inFilename := dirPath + "/" + f.Name()
	inFile, err := os.Open(inFilename)
	doPanic(err)
	defer inFile.Close()

	sourceImage, mimeType, err := image.Decode(inFile)
	doPanic(err)

	return inFilename, mimeType, &sourceImage
}

func addWatermark(watermark image.Image, pic *image.Image) image.Image {
	r := image.Rectangle{image.Point{}, (*pic).Bounds().Size()}
	r2 := image.Rectangle{image.Point{}, watermark.Bounds().Size()}

	rgba := image.NewRGBA(r)

	draw.Draw(rgba, (*pic).Bounds(), *pic, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, watermark, image.Point{0, 0}, draw.Over)

	return rgba
}

func saveFile(outfile *os.File, dist image.Image, mimeType string) {
	if "jpeg" == mimeType {
		jpeg.Encode(outfile, dist, nil)
	} else if "png" == mimeType {
		png.Encode(outfile, dist)
	}
}

func main() {
	user, err := osUser.Current()
	doPanic(err)

	fmt.Println("Hello, " + user.Name)

	cfg, _ := config.LoadConfig("./config.yml")

	wmFile, err := os.Open(cfg.WatermarkFile)
	doPanic(err)
	defer wmFile.Close()

	wmImage, _, err := image.Decode(wmFile)
	doPanic(err)

	files, err := ioutil.ReadDir(cfg.InputDir);
	doPanic(err)
	for _, f := range files {
		inFilename, mimeType, sourceImage := openImage(cfg.InputDir, f)
		fmt.Println(inFilename + " (" + mimeType + ")")

		dist := addWatermark(wmImage, sourceImage)

		outfile, err := os.Create(cfg.OutputDir + "/" + f.Name())
		doPanic(err)

		saveFile(outfile, dist, mimeType)

		outfile.Close()
	}
}

func doPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
