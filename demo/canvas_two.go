package main

import (
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
)

var x, y = 0, 0
var width, height int

// ----------------------------------------------------------------------------
func performDemoOnCanvasTwo() {
	width, height = canvasTwo.GetDimensions()
	canvasTwo.ClearScreen(*colour.NewColour(80, 80, 180, 255))
	setRefreshCanvasTwoCallback()
	animateCanvasTwo()
}

// ----------------------------------------------------------------------------
func updateCanvasTwo() {
	for range 10 {
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
	}

	canvasTwo.Render()
}

// ----------------------------------------------------------------------------
func setRefreshCanvasTwoCallback() {
	js.Global().Set("refreshCanvasTwo", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateCanvasTwo()
		return nil
	}))
}
