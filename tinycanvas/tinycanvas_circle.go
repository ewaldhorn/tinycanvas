package tinycanvas

import "math"

// ----------------------------------------------------------------------------
func (t *TinyCanvas) Circle(midX, midY, radius int, colour Colour) {
	rad := float64(radius)
	deg := 0.0
	for deg < 6.4 {
		x := int(rad * math.Cos(deg))
		y := int(rad * math.Sin(deg))
		t.PutPixel(midX+x, midY+y, colour)
		deg += 0.0025
	}
}
