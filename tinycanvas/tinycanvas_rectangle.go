package tinycanvas

import "github.com/ewaldhorn/tinycanvas/colour"

// ----------------------------------------------------------------------------
// Draws a rectangle of the specified width and height from the top left corner
// filled with the given colour.
func (t *TinyCanvas) FilledRectangle(xStart, yStart, width, height int, colour colour.Colour) {
	tmpC := t.activeColour
	t.activeColour = colour

	for x := range width {
		for y := range height {
			t.PutPixel(xStart+x, yStart+y)
		}
	}

	t.activeColour = tmpC
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
	tmpC := t.activeColour
	t.activeColour = colour

	t._rectangle(xStart, yStart, width, height)

	t.activeColour = tmpC
}

// ----------------------------------------------------------------------------
// Draws a rectangle with the specified width and height from the top left corner
// having a border of the specified thickness and default colour.
func (t *TinyCanvas) Rectangle(xStart, yStart, width, height, thickness int) {
	t._rectangle(xStart, yStart, width, height)
}
