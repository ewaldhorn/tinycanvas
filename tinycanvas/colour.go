package tinycanvas

import (
	"math/rand"
)

const MAX_COLOUR_VALUE uint8 = 255

// ----------------------------------------------------------------------------
type Colour struct {
	r, g, b, a uint8
}

// ----------------------------------------------------------------------------
// an empty colour is used for transparent effects
func (c *Colour) IsEmpty() bool {
	return c.a == 0 && c.b == 0 && c.g == 0 && c.r == 0
}

// ----------------------------------------------------------------------------
func NewColour(r, g, b, a uint8) *Colour {
	return &Colour{r: r, g: g, b: b, a: a}
}

// ----------------------------------------------------------------------------
func NewColourWhite() *Colour {
	return &Colour{r: MAX_COLOUR_VALUE, g: MAX_COLOUR_VALUE, b: MAX_COLOUR_VALUE, a: MAX_COLOUR_VALUE}
}

// ----------------------------------------------------------------------------
func NewColourBlack() *Colour {
	return &Colour{a: MAX_COLOUR_VALUE}
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
		r: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		g: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		b: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		a: MAX_COLOUR_VALUE,
	}
}
