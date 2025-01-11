package tinycanvas

import "github.com/ewaldhorn/tinycanvas/colour"

// ----------------------------------------------------------------------------
// Determine the direction of the line.
func getSlopes(x1, y1, x2, y2 int) (int, int) {
	slopeX := -1
	if x1 < x2 {
		slopeX = 1
	}

	slopeY := -1
	if y1 < y2 {
		slopeY = 1
	}

	return slopeX, slopeY
}

// ----------------------------------------------------------------------------
// Draws a one pixel line from X1,Y1 to X2,Y2 in the active colour.
// Based on https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// TODO: Look into anti-aliasing lines
func (t *TinyCanvas) Line(x1, y1, x2, y2 int) {
	// Calculate the absolute differences between the x and y coordinates.
	diffX := abs(x2 - x1)
	diffY := abs(y2 - y1)

	slopeX, slopeY := getSlopes(x1, y1, x2, y2)

	// Initialize the error value.
	errVal := diffX - diffY

	// Draw the line.
	for {
		t.PutPixel(x1, y1)

		// Check if we've reached the end of the line.
		if x1 == x2 && y1 == y2 {
			break
		}

		// Adjust the error value and update the x and y coordinates.
		errVal2 := 2 * errVal
		if errVal2 > -diffY {
			errVal -= diffY
			x1 += slopeX
		}
		if errVal2 < diffX {
			errVal += diffX
			y1 += slopeY
		}
	}
}

// ----------------------------------------------------------------------------
// Draws a one pixel line from X1,Y1 to X2,Y2 in the specified colour.
// Based on https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// TODO: Look into anti-aliasing lines
func (t *TinyCanvas) ColourLine(x1, y1, x2, y2 int, colour colour.Colour) {
	t.SwitchAndSaveColour(colour)
	t.Line(x1, y1, x2, y2)
	t.RestoreColour()
}
