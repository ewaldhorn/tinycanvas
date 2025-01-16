package main

import (
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

var x, y = 0, 0
var width, height int

const pixelsPerTick = 40

// ----------------------------------------------------------------------------
func performDemoOnCanvasTwo() {
	width, height = canvasTwo.GetDimensions()
	canvasTwo.ClearScreen(*colour.NewColour(80, 80, 180, 255))
	setRefreshCanvasTwoCallback()

	for i := 0; i < 40; i += 2 {
		canvasTwo.Triangle(
			tinycanvas.Point{X: (width / 2), Y: (height / 3) + i},
			tinycanvas.Point{X: (width - (width / 3)) + i, Y: (height - (height / 3)) - i},
			tinycanvas.Point{X: (width / 3) - i, Y: (height - (height / 3)) - i},
		)
	}

	animateCanvasTwo()
}

// ----------------------------------------------------------------------------
func updateCanvasTwo() {
	for range pixelsPerTick {
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
