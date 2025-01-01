package tinycanvas

const MAX_PIXEL_VALUE uint8 = 255

// ----------------------------------------------------------------------------
type Pixel struct {
	r, g, b, a uint8
}

// ----------------------------------------------------------------------------
func NewPixel(r, g, b, a uint8) *Pixel {
	return &Pixel{r: r, g: g, b: b, a: a}
}

// ----------------------------------------------------------------------------
func NewWhitePixel() *Pixel {
	return &Pixel{r: MAX_PIXEL_VALUE, g: MAX_PIXEL_VALUE, b: MAX_PIXEL_VALUE, a: MAX_PIXEL_VALUE}
}
