package tinycanvas

// ----------------------------------------------------------------------------
// Draws a one pixel line from X1,Y1 to X2,Y2 in the specified colour.
// Based on https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// TODO: Look into anti-aliasing lines
func (t *TinyCanvas) Line(x1, y1, x2, y2 int, p Colour) {
	dx := x2 - x1
	dy := y2 - y1
	d := 2*dy - dx
	incrE := 2 * dy
	incrNE := 2 * (dy - dx)
	x := x1
	y := y1
	t.PutPixel(x, y, p)

	for x < x2 {
		if d <= 0 {
			d += incrE
			x++
		} else {
			d += incrNE
			x++
			y++
		}
		t.PutPixel(x, y, p)
	}
}
