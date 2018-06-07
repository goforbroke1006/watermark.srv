package fs

import (
	"image"
	"os"
	"strings"
	"image/jpeg"
	"image/png"
	"fmt"
)

func IsFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func ParseFilename(filename string) (directory, name, extension string) {
	parts := strings.Split(filename, "/")
	simpleName := parts[len(parts)-1]

	parts2 := strings.Split(simpleName, ".")

	directory = strings.Replace(filename, "/"+simpleName, "", 1)
	name = parts2[0]
	extension = parts2[1]

	return
}

/*func GetImage(file io.Reader) (image.Image, string, error) {
	img, mime, err := image.Decode(file)
	if nil != err {
		return nil, "", err
	}

	return img, mime, nil
}*/

/*func SaveDraw(file *os.File, dist *draw.Image) {
	_, _, ext := ParseFilename(file.Name())

	if "jpeg" == ext {
		jpeg.Encode(file, dist, nil)
	} else if "png" == ext {
		png.Encode(file, dist)
	}
}*/

func SaveImage(filename string, file *os.File, dist image.Image) error {
	_, _, ext := ParseFilename(filename)

	if "jpeg" == ext || "jpg" == ext {
		return jpeg.Encode(file, dist, nil)
	} else if "png" == ext {
		return png.Encode(file, dist)
	}

	return fmt.Errorf("undefined format: %s", ext)
}
