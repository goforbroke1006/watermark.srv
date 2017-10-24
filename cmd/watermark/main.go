package main

import (
	"fmt"
	"log"
	osUser "os/user"

	"image"
	"image/png"
	"io"
	_ "golang.org/x/image/webp"
	_ "image/jpeg"
)

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}

func main() {
	user, err := osUser.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello, " + user.Name)
}
