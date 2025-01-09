package colour

import (
	"math/rand"
)

const MAX_COLOUR_VALUE uint8 = 255

// ----------------------------------------------------------------------------
type Colour struct {
	R, G, B, A uint8
}

// ----------------------------------------------------------------------------
// an empty colour is used for transparent effects
func (c *Colour) IsEmpty() bool {
	return c.A == 0 && c.B == 0 && c.G == 0 && c.R == 0
}

// ----------------------------------------------------------------------------
func NewColour(r, g, b, a uint8) *Colour {
	return &Colour{R: r, G: g, B: b, A: a}
}

// ----------------------------------------------------------------------------
func NewColourWhite() *Colour {
	return &Colour{R: MAX_COLOUR_VALUE, G: MAX_COLOUR_VALUE, B: MAX_COLOUR_VALUE, A: MAX_COLOUR_VALUE}
}

// ----------------------------------------------------------------------------
func NewColourBlack() *Colour {
	return &Colour{A: MAX_COLOUR_VALUE}
}

// ----------------------------------------------------------------------------
// empty colour signals to the renderer not to draw anything and is used to create
// transparent "gaps" in images. This is a legacy feature to support some very
// old file formats.
func NewColourEmpty() *Colour {
	return &Colour{}
}

// ----------------------------------------------------------------------------
func NewRandomColour() *Colour {
	return &Colour{
		R: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		G: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		B: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		A: MAX_COLOUR_VALUE,
	}
}

// ----------------------------------------------------------------------------
// Built using information from https://en.wikipedia.org/wiki/Grayscale
// and https://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang
func (c *Colour) ConvertToGrayscale() {
	shadeOfGray := uint8((0.299*(float64(c.R)) + 0.587*(float64(c.G)) + 0.144*(float64(c.B))) / 256)

	c.R = shadeOfGray
	c.G = shadeOfGray
	c.B = shadeOfGray
}
