package tinycanvas

import (
	"math"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
func (t *TinyCanvas) Circle(midX, midY, radius int, colour colour.Colour) {
	rad := float64(radius)
	deg := 0.0
	for deg < 6.4 {
		x := int(rad * math.Cos(deg))
		y := int(rad * math.Sin(deg))
		t.PutColourPixel(midX+x, midY+y, colour)
		deg += 0.0025
	}
}
