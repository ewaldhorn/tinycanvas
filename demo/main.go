package main

import (
	"fmt"

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
	dom.Log(fmt.Sprintf("Canvas has a height of %d and width of %d", width, height))

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
}
