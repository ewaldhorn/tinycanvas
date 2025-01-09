package main

import "github.com/ewaldhorn/tinycanvas/colour"

var x, y = 0, 0
var width, height int

// ----------------------------------------------------------------------------
func performDemoOnCanvasTwo() {
	width, height = canvasTwo.GetDimensions()
	canvasTwo.ClearScreen(*colour.NewColour(80, 80, 180, 255))

	canvasTwo.Render()
}

// ----------------------------------------------------------------------------
func updateCanvasTwo() {
	x += 1

	if x >= width {
		x = 1
		y += 1
	}

	if y >= height {
		y = 1
	}

	canvasTwo.SetColour(*colour.NewRandomColour())
	canvasTwo.PutPixel(x, y)
	canvasTwo.Render()
}
