package main

import (
	"github.com/ewaldhorn/tinycanvas/dom"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

//export bootstrap
func bootstrap()

// ----------------------------------------------------------------------------
func main() {
	setCallbacks()
	dom.Hide("loading")
	bootstrap()

	canvas := tinycanvas.NewTinyCanvas(640, 480)
	width, height := canvas.GetDimensions()

	canvas.ClearScreen(*tinycanvas.NewPixel(10, 20, 180, 255))
	canvas.Rectangle(10, 10, width-20, height-20, *tinycanvas.NewPixel(0, 255, 0, 255))

	canvas.Render()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
}
