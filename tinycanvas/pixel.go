package tinycanvas

// ----------------------------------------------------------------------------
type Pixel struct {
	r, g, b, a uint8
}

// ----------------------------------------------------------------------------
func NewPixel(r, g, b, a uint8) *Pixel {
	return &Pixel{r: r, g: g, b: b, a: a}
}
