package tinycanvas

// ----------------------------------------------------------------------------
// Draws a triangle between three points.
func (t *TinyCanvas) Triangle(p1, p2, p3 Point) {
	t.LinePoint(p1, p2)
	t.LinePoint(p2, p3)
	t.LinePoint(p1, p3)
}
