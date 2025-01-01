package tinycanvas

import "syscall/js"

// ----------------------------------------------------------------------------
type DOM struct {
	window   js.Value
	document js.Value
	body     js.Value
}

// ----------------------------------------------------------------------------
type TinyCanvas struct {
	DOM

	canvas    js.Value
	ctx       js.Value
	imageData js.Value
	width     int
	height    int
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
}

// ----------------------------------------------------------------------------
func (t *TinyCanvas) GetDimensions() (int, int) {
	return t.width, t.height
}
