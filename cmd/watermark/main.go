package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/goforbroke1006/watermarksvc/config"
	"github.com/goforbroke1006/watermarksvc/strategy"
	"github.com/goforbroke1006/watermarksvc/util/fs"
)

const (
	BackupSuffix = "_origin"
)

func main() {
	cfg, _ := config.LoadConfig("./config.yml")

	watermark, err := os.Open(cfg.WatermarkFile)
	checkErr(err)
	wmImage, _, err := image.Decode(watermark)
	checkErr(err)
	watermark.Close()

	stg := &strategy.UglySplitStrategy{Rows: 4}

	semaphore := make(chan bool, 4)

	for {
		for _, directory := range cfg.Directories {

			files, err := ioutil.ReadDir(directory)
			checkErr(err)

			for _, f := range files {
				if strings.Contains(f.Name(), BackupSuffix+".") {
					continue
				}

				semaphore <- true

				go func(f os.FileInfo) {
					defer func(semaphore <-chan bool) { <-semaphore }(semaphore)

					filename := directory + "/" + f.Name()

					if hasBackupFile(filename) {
						return
					}

					fmt.Println(filename)

					doBackupFile(filename)
					addWatermark(filename, wmImage, stg)
				}(f)
			}

			time.Sleep(5 * time.Second)
		}
	}
}

func getBackupFilename(filename string) string {
	dir, name, ext := fs.ParseFilename(filename)
	return dir + "/" + name + BackupSuffix + "." + ext
}

func hasBackupFile(filename string) bool {
	return fs.IsFileExists(
		getBackupFilename(filename),
	)
}

func doBackupFile(filename string) {
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	srcImage, _, err := image.Decode(file)
	checkErr(err)

	origFilename := getBackupFilename(filename)

	origFile, err := os.Create(origFilename)
	checkErr(err)
	defer origFile.Close()

	err = fs.SaveImage(filename, origFile, srcImage)
	checkErr(err)
}

func addWatermark(filename string, watermark image.Image, strategy strategy.BaseStrategy) {
	file, err := os.Open(filename)
	checkErr(err)

	src, _, err := image.Decode(file)
	checkErr(err)

	file.Close()

	proceedImg := strategy.AddWatermark(watermark, src)

	file2, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	err = fs.SaveImage(filename, file2, proceedImg)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
