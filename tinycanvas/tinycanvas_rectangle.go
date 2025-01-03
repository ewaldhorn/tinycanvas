package tinycanvas

// ----------------------------------------------------------------------------
// Draws a rectangle of the specified width and height from the top left corner
// filled with the given colour.
func (t *TinyCanvas) FilledRectangle(xStart, yStart, width, height int, colour Colour) {
	for x := range width {
		for y := range height {
			t.PutPixel(xStart+x, yStart+y, colour)
		}
	}
}

// ----------------------------------------------------------------------------
// Draws a single pixel wide Rectangle
func (t *TinyCanvas) _rectangle(xStart, yStart, width, height int, colour Colour) {
	t.Line(xStart, yStart, xStart+width, yStart, colour)               // top
	t.Line(xStart+width, yStart, xStart+width, yStart+height, colour)  // right
	t.Line(xStart, yStart+height, xStart+width, yStart+height, colour) // bottom
	t.Line(xStart, yStart, xStart, yStart+height, colour)              // left
}

// ----------------------------------------------------------------------------
// Draws a rectangle with the specified width and height from the top left corner
// having a border of the specified thickness and colour.
func (t *TinyCanvas) Rectangle(xStart, yStart, width, height, thickness int, colour Colour) {
	// TODO We need LINES
	t._rectangle(xStart, yStart, width, height, colour)
	t.Line(400, 100, 400, 400, *NewColour(255, 128, 128, 255))
	t.Line(401, 100, 401, 400, *NewColour(255, 128, 128, 255))
	t.Line(402, 100, 402, 400, *NewColour(255, 128, 128, 255))
	t.Line(403, 100, 403, 400, *NewColour(255, 128, 128, 255))

	for i := range 50 {
		t.PutPixel(450+i, 100, *NewColourWhite())
	}

}
