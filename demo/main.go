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

	canvas.ClearScreen(*tinycanvas.NewColour(10, 20, 180, 255))
	canvas.FilledRectangle(10, 10, width-20, height-20, *tinycanvas.NewColour(0, 255, 0, 255))
	canvas.FilledRectangle(10, 10, width/2, height/2, *tinycanvas.NewColourWhite())
	canvas.FilledRectangle(510, 350, 120, 120, *tinycanvas.NewColourWhite())
	drawWithPixels(canvas)
	drawRandomFilledRectangles(canvas)
	drawRandomRectangles(canvas)
	drawRandomCircles(canvas)

	drawLines(canvas)

	canvas.Render()
}

// ----------------------------------------------------------------------------
// Canvas supports setting individual pixels.
func drawWithPixels(canvas *tinycanvas.TinyCanvas) {
	redPixel := *tinycanvas.NewColour(255, 0, 0, 255)

	for x := 0; x < 50; x += 5 {
		for y := 0; y < 50; y += 5 {
			canvas.PutColourPixel(280+x, 200+y, redPixel)
		}
	}
}

// ----------------------------------------------------------------------------
func drawRandomFilledRectangles(canvas *tinycanvas.TinyCanvas) {
	for x := range 40 {
		for y := range 40 {
			canvas.FilledRectangle(20+(x*4), 20+(y*4), 20, 20, *tinycanvas.NewRandomColour())
		}
	}
}

// ----------------------------------------------------------------------------
func drawRandomRectangles(canvas *tinycanvas.TinyCanvas) {
	for i := range 10 {
		canvas.ColourRectangle(340+(i*15), 20, 10, 10, 1, *tinycanvas.NewRandomColour())
	}
}

// ----------------------------------------------------------------------------
func drawRandomCircles(canvas *tinycanvas.TinyCanvas) {
	for i := 6; i < 60; i += 3 {
		canvas.Circle(570, 410, i, *tinycanvas.NewRandomColour())
	}
}

// ----------------------------------------------------------------------------
func drawLines(canvas *tinycanvas.TinyCanvas) {

	canvas.Line(10, 400, 630, 400)
	canvas.Line(10, 405, 630, 405)
	canvas.Line(10, 410, 630, 410)
	canvas.Line(10, 415, 630, 415)
	canvas.Line(10, 420, 630, 420)

	for i := 0; i < 50; i += 5 {
		canvas.ColourLine(20, 300+i, 620, 300+i, *tinycanvas.NewRandomColour())
	}

	canvas.ColourLine(215, 15, 250, 250, *tinycanvas.NewColour(0, 0, 0, 255))
}
