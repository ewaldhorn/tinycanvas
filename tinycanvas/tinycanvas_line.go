package tinycanvas

// ----------------------------------------------------------------------------
// Draws a one pixel line from X1,Y1 to X2,Y2 in the active colour.
// Based on https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// TODO: Look into anti-aliasing lines
func (t *TinyCanvas) Line(x1, y1, x2, y2 int) {
	// Calculate the absolute differences between the x and y coordinates.
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)

	// Determine the direction of the line.
	sx := -1
	if x1 < x2 {
		sx = 1
	}
	sy := -1
	if y1 < y2 {
		sy = 1
	}

	// Initialize the error value.
	err := dx - dy

	// Draw the line.
	for {
		t.PutColourPixel(x1, y1, t.activeColour)

		// Check if we've reached the end of the line.
		if x1 == x2 && y1 == y2 {
			break
		}

		// Adjust the error value and update the x and y coordinates.
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

// ----------------------------------------------------------------------------
// Draws a one pixel line from X1,Y1 to X2,Y2 in the specified colour.
// Based on https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// TODO: Look into anti-aliasing lines
func (t *TinyCanvas) ColourLine(x1, y1, x2, y2 int, p Colour) {
	tmpC := t.activeColour
	t.activeColour = p

	t.Line(x1, y1, x2, y2)

	t.activeColour = tmpC
}
