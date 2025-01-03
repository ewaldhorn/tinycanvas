package tinycanvas

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
type DOMElements struct {
	window   js.Value
	document js.Value
	body     js.Value
}

// ----------------------------------------------------------------------------
type TinyCanvas struct {
	DOMElements

	canvas        js.Value
	ctx           js.Value
	imageData     js.Value
	imageBuffer   js.Value
	wasmImageData []uint8
	width         int
	height        int
}

// ----------------------------------------------------------------------------
func NewTinyCanvas(width, height int) *TinyCanvas {
	tinyCanvas := &TinyCanvas{width: width, height: height}

	tinyCanvas.window = js.Global()
	tinyCanvas.document = tinyCanvas.window.Get("document")
	tinyCanvas.body = tinyCanvas.document.Get("body")

	// make it fit the window if needed
	if width == 0 || height == 0 {
		tinyCanvas.width = tinyCanvas.window.Get("innerWidth").Int()
		tinyCanvas.height = tinyCanvas.window.Get("innerHeight").Int()
	}

	tinyCanvas.createHTMLCanvasElement()

	return tinyCanvas
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) createHTMLCanvasElement() {
	t.canvas = t.document.Call("createElement", "canvas")
	t.canvas.Set("height", t.height)
	t.canvas.Set("width", t.width)
	t.body.Call("appendChild", t.canvas)

	t.refreshCanvasProperties()
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) refreshCanvasProperties() {
	t.ctx = t.canvas.Call("getContext", "2d")
	t.imageData = t.ctx.Call("createImageData", t.width, t.height)
	t.wasmImageData = make([]uint8, t.width*t.height*4) // width*height*pixel_size
	t.imageBuffer = js.Global().Get("Uint8Array").New(len(t.wasmImageData))
	t.wasmImageData = make([]uint8, t.width*t.height*4)
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) GetDimensions() (int, int) {
	return t.width, t.height
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) ClearScreen(p Colour) {
	for x := range t.width {
		for y := range t.height {
			t.PutPixel(x, y, p)
		}
	}
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) PutPixel(x, y int, p Colour) {
	offset := (x * 4) + (y * 4 * t.width)

	// don't bother if we are outside our area
	if offset < 0 || offset > int(len(t.wasmImageData)) {
		return
	}

	t.wasmImageData[offset] = p.r
	t.wasmImageData[offset+1] = p.g
	t.wasmImageData[offset+2] = p.b
	t.wasmImageData[offset+3] = p.a
}

// ----------------------------------------------------------------------------
// Copies buffer to canvas
func (t *TinyCanvas) Render() {
	js.CopyBytesToJS(t.imageBuffer, t.wasmImageData)   // copy local buffer to JS buffer
	t.imageData.Get("data").Call("set", t.imageBuffer) // copy that data to the canvas image data buffer
	t.ctx.Call("putImageData", t.imageData, 0, 0)      // finally render the image to the canvas
}
