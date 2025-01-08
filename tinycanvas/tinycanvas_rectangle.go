package tinycanvas

import (
	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
func (t *TinyCanvas) FilledRectangle(xStart, yStart, width, height int) {
	for x := range width {
		for y := range height {
			t.PutPixel(xStart+x, yStart+y)
		}
	}
}

// ----------------------------------------------------------------------------
// Draws a rectangle of the specified width and height from the top left corner
// filled with the given colour.
func (t *TinyCanvas) ColourFilledRectangle(xStart, yStart, width, height int, colour colour.Colour) {
	t.SwitchAndSaveColour(colour)
	t.FilledRectangle(xStart, yStart, width, height)
	t.RestoreColour()
}

// ----------------------------------------------------------------------------
// Draws a single pixel wide Rectangle in the default canvas colour
func (t *TinyCanvas) _rectangle(xStart, yStart, width, height int) {
	t.Line(xStart, yStart, xStart+width, yStart)               // top
	t.Line(xStart+width, yStart, xStart+width, yStart+height)  // right
	t.Line(xStart, yStart+height, xStart+width, yStart+height) // bottom
	t.Line(xStart, yStart, xStart, yStart+height)              // left
}

// ----------------------------------------------------------------------------
// Draws a rectangle with the specified width and height from the top left corner
// having a border of the specified thickness and colour.
func (t *TinyCanvas) ColourRectangle(xStart, yStart, width, height, thickness int, colour colour.Colour) {
	t.SwitchAndSaveColour(colour)
	t._rectangle(xStart, yStart, width, height)
	t.RestoreColour()
}

// ----------------------------------------------------------------------------
// Draws a rectangle with the specified width and height from the top left corner
// having a border of the specified thickness and default colour.
func (t *TinyCanvas) Rectangle(xStart, yStart, width, height, thickness int) {
	t._rectangle(xStart, yStart, width, height)
}
