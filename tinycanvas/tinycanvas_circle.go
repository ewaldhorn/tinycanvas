package tinycanvas

import (
	"math"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
func (t *TinyCanvas) Circle(midX, midY, radius int) {
	rad := float64(radius)
	deg := 0.0
	for deg < 6.4 {
		x := int(rad * math.Cos(deg))
		y := int(rad * math.Sin(deg))
		t.PutPixel(midX+x, midY+y)
		deg += 0.0025
	}
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) ColourCircle(midX, midY, radius int, colour colour.Colour) {
	// preserve canvas colour
	tmpC := t.activeColour
	t.SetColour(colour)

	t.Circle(midX, midY, radius)

	// restore original canvas colour
	t.SetColour(tmpC)
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) FilledCircle(midX, midY, radius int) {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				t.PutPixel(midX+x, midY+y)
			}
		}
	}
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) ColourFilledCircle(midX, midY, radius int, colour colour.Colour) {
	// preserve canvas colour
	tmpC := t.activeColour
	t.SetColour(colour)

	t.FilledCircle(midX, midY, radius)

	// restore original canvas colour
	t.SetColour(tmpC)
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) BorderCircle(midX, midY, radius, borderWidth int) {
	innerRadius := radius - borderWidth
	innerRadiusSquared := innerRadius * innerRadius
	radiusSquared := radius * radius

	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			distSquared := x*x + y*y
			if distSquared <= radiusSquared && distSquared > innerRadiusSquared {
				t.PutPixel(midX+x, midY+y)
			}
		}
	}
}
