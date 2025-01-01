package main

import (
	"github.com/ewaldhorn/tinycanvas/dom"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ------------------------------------------ Externally provided by JavaScript

//export bootstrap
func bootstrap()

// ----------------------------------------------------------------------------
func main() {
	setCallbacks()
	dom.Hide("loading")
	bootstrap()

	performDemo()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
}

// ----------------------------------------------------------------------------
func performDemo() {
	canvas := tinycanvas.NewTinyCanvas(640, 480)
	width, height := canvas.GetDimensions()

	canvas.ClearScreen(*tinycanvas.NewPixel(10, 20, 180, 255))
	canvas.Rectangle(10, 10, width-20, height-20, *tinycanvas.NewPixel(0, 255, 0, 255))
	canvas.Rectangle(10, 10, width/2, height/2, *tinycanvas.NewWhitePixel())
	drawWithPixels(canvas)
	drawRandomRectangles(canvas)

	canvas.Render()
}

// ----------------------------------------------------------------------------
// Canvas supports setting individual pixels.
func drawWithPixels(canvas *tinycanvas.TinyCanvas) {
	redPixel := *tinycanvas.NewPixel(255, 0, 0, 255)

	for x := 0; x < 50; x += 5 {
		for y := 0; y < 50; y += 5 {
			canvas.PutPixel(280+x, 200+y, redPixel)
		}
	}
}

// ----------------------------------------------------------------------------
func drawRandomRectangles(canvas *tinycanvas.TinyCanvas) {
	for x := range 40 {
		for y := range 40 {
			canvas.Rectangle(20+(x*4), 20+(y*4), 20, 20, *tinycanvas.NewRandomPixel())
		}
	}
}
