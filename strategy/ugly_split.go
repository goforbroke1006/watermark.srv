package strategy

import (
	"image"
	"image/draw"
)
import "github.com/nfnt/resize"

type UglySplitStrategy struct {
	Rows int
}

func (s *UglySplitStrategy) AddWatermark(watermark *image.Image, pic *image.Image) *image.RGBA {
	var partHeight int = (*pic).Bounds().Size().Y / s.Rows
	var proportion float32 = float32(partHeight) / float32((*watermark).Bounds().Size().Y)
	var partWidth int = int(float32((*watermark).Bounds().Size().X) * proportion)
	var cols int = (*pic).Bounds().Size().X / partWidth

	r := image.Rectangle{Min: image.Point{}, Max: (*pic).Bounds().Size()}
	rgba := image.NewRGBA(r)
	draw.Draw(rgba, (*pic).Bounds(), *pic, image.Point{0, 0}, draw.Src)

	newImage := resize.Resize(uint(partWidth), uint(partHeight), *watermark, resize.Lanczos3)

	for row := 0; row <= s.Rows; row++ {
		for column := 0; column <= cols; column++ {
			wShift := column * partWidth
			hShift := row * partHeight
			r2 := image.Rectangle{
				Min: image.Point{
					X: wShift,
					Y: hShift,
				},
				Max: image.Point{
					X: wShift + partWidth,
					Y: hShift + partHeight,
				},
			}
			draw.Draw(rgba, r2, newImage, image.Point{0, 0}, draw.Over)
		}
	}

	return rgba
}
