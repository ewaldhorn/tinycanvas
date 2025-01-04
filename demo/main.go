package main

import (
	"github.com/ewaldhorn/tinycanvas/dom"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ------------------------------------------ Externally provided by JavaScript

//export bootstrap
func bootstrap()

// -------------------------------------------------------------------- GLOBALS
var canvas *tinycanvas.TinyCanvas

// ----------------------------------------------------------------------------
func main() {
	setCallbacks()
	dom.Hide("loading")
	bootstrap()
	canvas = tinycanvas.NewTinyCanvas(640, 480)

	performDemo()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
	setRefreshCallback()
}

// ----------------------------------------------------------------------------
func performDemo() {
	width, height := canvas.GetDimensions()

	canvas.ClearScreen(*tinycanvas.NewColour(10, 20, 180, 255))
	canvas.FilledRectangle(10, 10, width-20, height-20, *tinycanvas.NewColour(0, 255, 0, 255))
	canvas.FilledRectangle(10, 10, (width/2)-10, (height/2)-10, *tinycanvas.NewColourWhite())
	canvas.FilledRectangle(510, 350, 120, 120, *tinycanvas.NewColourWhite())
	drawLines(canvas)
	drawWithPixels(canvas)
	drawRandomFilledRectangles(canvas)
	drawRandomRectangles(canvas)
	drawRandomCircles(canvas)

	canvas.Render()
}

// ----------------------------------------------------------------------------
// Canvas supports setting individual pixels.
func drawWithPixels(canvas *tinycanvas.TinyCanvas) {
	redPixel := *tinycanvas.NewColour(255, 0, 0, 255)

	for x := 0; x < 50; x += 5 {
		for y := 0; y < 50; y += 5 {
			canvas.PutColourPixel(270+x, 190+y, redPixel)
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
	for y := range 4 {
		for i := range 10 {
			canvas.ColourRectangle(340+(i*15), 20+(20*y), 10, 10, 1, *tinycanvas.NewRandomColour())
		}
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
	canvas.SetColour(*tinycanvas.NewColour(255, 0, 0, 255))
	w, h := canvas.GetDimensions()
	canvas.Line(0, 0, w, h) // top left to bottom right
	canvas.Line(0, h, w, 0) // bottom left to top right

	for i := 0; i < 50; i += 5 {
		canvas.ColourLine(20, 300+i, 620, 300+i, *tinycanvas.NewRandomColour())
	}
}

// ----------------------------------------------------------------------------
func setRefreshCallback() {
	// js.Global().Set("rerender", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	if canvas != nil {
	// 		performDemo()
	// 	}
	// 	return nil
	// }))

	dom.AddEventListener("refreshButton", "click", func() {
		if canvas != nil {
			performDemo()
		}
	})
}
