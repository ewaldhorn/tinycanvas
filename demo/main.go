package main

import (
	"github.com/ewaldhorn/tinycanvas/colour"
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
	dom.Show("controls")
	bootstrap()
	canvas = tinycanvas.NewTinyCanvas(800, 600)

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

	canvas.ClearScreen(*colour.NewColour(10, 20, 180, 255))

	canvas.FilledRectangle(10, 10, width-20, height-20, *colour.NewColour(0, 255, 0, 255))
	drawLines(canvas)
	canvas.FilledRectangle(10, 10, (width/2)-10, (height/2)-10, *colour.NewColourWhite())
	canvas.FilledRectangle(width-130, height-130, 120, 120, *colour.NewColourWhite())

	drawWithPixels(canvas)
	drawRandomFilledRectangles(canvas)
	drawRandomRectangles(canvas)
	drawRandomCircles(canvas)
	drawSpiral(canvas)

	canvas.Render()
}

// ----------------------------------------------------------------------------
// Canvas supports setting individual pixels.
func drawWithPixels(canvas *tinycanvas.TinyCanvas) {
	redPixel := *colour.NewColour(255, 0, 0, 255)
	const gridSize int = 50
	for x := 0; x < gridSize; x += 5 {
		for y := 0; y < gridSize; y += 5 {
			canvas.PutColourPixel(canvas.Width()/2-gridSize+x, canvas.Height()/2-gridSize+y, redPixel)
		}
	}
}

// ----------------------------------------------------------------------------
func drawRandomFilledRectangles(canvas *tinycanvas.TinyCanvas) {
	for x := range 40 {
		for y := range 40 {
			canvas.FilledRectangle(20+(x*4), 20+(y*4), 20, 20, *colour.NewRandomColour())
		}
	}
}

// ----------------------------------------------------------------------------
func drawRandomRectangles(canvas *tinycanvas.TinyCanvas) {
	for y := range 4 {
		for i := range 10 {
			canvas.ColourRectangle(10+canvas.Width()/2+(i*15), 20+(20*y), 10, 10, 1, *colour.NewRandomColour())
		}
	}
}

// ----------------------------------------------------------------------------
func drawRandomCircles(canvas *tinycanvas.TinyCanvas) {

	for i := range 24 {
		canvas.ColourFilledCircle((i*4)+50+canvas.Width()/4, 60, 40, *colour.NewColour(64+uint8(i*8), 64, 255, 255))
	}

	canvas.BorderCircle(100+canvas.Width()/4, canvas.Height()/4, 35, 3)

	for i := 6; i < 60; i += 3 {
		canvas.ColourCircle(canvas.Width()-70, canvas.Height()-70, i, *colour.NewRandomColour())
	}
}

// ----------------------------------------------------------------------------
func drawLines(canvas *tinycanvas.TinyCanvas) {
	canvas.SetColour(*colour.NewColour(255, 0, 0, 255))
	w, h := canvas.GetDimensions()
	canvas.Line(0, 0, w, h)     // top left to bottom right
	canvas.Line(w/2, h/2, w, 0) // middle to top right

	canvas.SetColour(*colour.NewColourBlack())

	for i := 0; i < 60; i += 5 {
		canvas.Line(40, 10+canvas.Height()/2+i, canvas.Width()/2-i-40, 10+canvas.Height()/2+i)                  // top
		canvas.Line(40+i, 10+canvas.Height()/2+i, 40+i, canvas.Height()-20-i)                                   // left
		canvas.Line(canvas.Width()/2-i-40, 10+canvas.Height()/2+i, canvas.Width()/2-i-40, canvas.Height()-20-i) //right
		canvas.Line(40, canvas.Height()-i-20, canvas.Width()/2-i-40, canvas.Height()-i-20)                      // bottom
	}
}

// ----------------------------------------------------------------------------
func drawSpiral(canvas *tinycanvas.TinyCanvas) {
	centerX, centerY := canvas.Width()/4, canvas.Height()/2+(canvas.Height()/4)

	// Initial direction (0: east, 1: south, 2: west, 3: north)
	direction := 0

	// Number of steps to take in each direction
	steps := 1

	// Number of times to change direction
	targetSize := 32

	stepSize := 2

	// Current number of changes
	currentSize := 0

	// Previous position
	prevX, prevY := centerX, centerY

	for currentSize < targetSize {
		// Take steps in the current direction
		for i := 0; i < steps*2; i++ {
			// Store the current position as the previous position
			prevX, prevY = centerX, centerY

			// Move in the current direction
			switch direction {
			case 0: // east
				centerX += stepSize
			case 1: // south
				centerY -= stepSize
			case 2: // west
				centerX -= stepSize
			case 3: // north
				centerY += stepSize
			}

			// Draw a line from the previous position to the current position
			canvas.Line(prevX, prevY, centerX, centerY)
		}

		// Change direction
		direction = (direction + 1) % 4

		// Increase the number of steps for the next direction
		steps++

		// Increment the change counter
		currentSize++
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
