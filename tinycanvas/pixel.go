package tinycanvas

import (
	"math/rand"
)

const MAX_PIXEL_VALUE uint8 = 255

// ----------------------------------------------------------------------------
type Pixel struct {
	r, g, b, a uint8
}

// ----------------------------------------------------------------------------
func NewPixel(r, g, b, a uint8) *Pixel {
	return &Pixel{r: r, g: g, b: b, a: a}
}

// ----------------------------------------------------------------------------
func NewWhitePixel() *Pixel {
	return &Pixel{r: MAX_PIXEL_VALUE, g: MAX_PIXEL_VALUE, b: MAX_PIXEL_VALUE, a: MAX_PIXEL_VALUE}
}

// ----------------------------------------------------------------------------
func NewRandomPixel() *Pixel {
	return &Pixel{
		r: uint8(rand.Float32() * float32(MAX_PIXEL_VALUE)),
		g: uint8(rand.Float32() * float32(MAX_PIXEL_VALUE)),
		b: uint8(rand.Float32() * float32(MAX_PIXEL_VALUE)),
		a: MAX_PIXEL_VALUE,
	}
}
