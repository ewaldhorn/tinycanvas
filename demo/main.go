package main

import (
	"github.com/ewaldhorn/tinycanvas/dom"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ------------------------------------------ Externally provided by JavaScript

//export bootstrap
func bootstrap()

//export animateCanvasTwo
func animateCanvasTwo()

// -------------------------------------------------------------------- GLOBALS
var canvasOne, canvasTwo *tinycanvas.TinyCanvas

// ----------------------------------------------------------------------------
func main() {
	setCallbacks()
	dom.Hide("loading")
	dom.Show("controls")
	bootstrap()
	canvasOne = tinycanvas.NewTinyCanvas(800, 600)
	canvasTwo = tinycanvas.NewTinyCanvas(320, 200)

	performDemoOnCanvasOne()
	performDemoOnCanvasTwo()

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
func setRefreshCallback() {
	// js.Global().Set("rerender", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	if canvas != nil {
	// 		performDemo()
	// 	}
	// 	return nil
	// }))

	dom.AddEventListener("refreshButton", "click", func() {
		if canvasOne != nil {
			performDemoOnCanvasOne()
		}
	})
}
