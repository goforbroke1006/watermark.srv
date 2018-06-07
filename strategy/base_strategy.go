package strategy

import (
	"image"
)

type BaseStrategy interface {
	AddWatermark(watermark image.Image, pic image.Image) image.Image
}
