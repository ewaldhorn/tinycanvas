package main

import "github.com/ewaldhorn/tinycanvas/colour"

// ----------------------------------------------------------------------------
func performDemoOnCanvasTwo() {
	// width, height := canvasTwo.GetDimensions()

	canvasTwo.ClearScreen(*colour.NewColour(80, 80, 180, 255))

	canvasTwo.Render()
}
