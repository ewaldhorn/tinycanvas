package dom

// Package dom provides a small DOM library for TinyGo, enabling DOM manipulation
// in WebAssembly TinyGo applications.

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
var document js.Value
var console js.Value

// ----------------------------------------------------------------------------
func init() {
	document = js.Global().Get("document")
	console = js.Global().Get("console")
}

// ----------------------------------------------------------------------------
//
// Log calls JavaScript console.log in the browser.
func Log(message string) {
	console.Call("log", message)
}

// ----------------------------------------------------------------------------
func getElementById(elem string) js.Value {
	return document.Call("getElementById", elem)
}

// ----------------------------------------------------------------------------
func getElementValue(elem string, value string) js.Value {
	return getElementById(elem).Get(value)
}

// ----------------------------------------------------------------------------
func Hide(elem string) {
	getElementValue(elem, "style").Call("setProperty", "display", "none")
}

// ----------------------------------------------------------------------------
func Show(elem string) {
	getElementValue(elem, "style").Call("setProperty", "display", "block")
}

// ----------------------------------------------------------------------------
func SetFocus(elem string) {
	getElementById(elem).Call("focus")
}

// ----------------------------------------------------------------------------
func GetString(elem string, value string) string {
	return getElementValue(elem, value).String()
}

// ----------------------------------------------------------------------------
func SetValue(elem string, key string, value string) {
	getElementById(elem).Set(key, value)
}

// ----------------------------------------------------------------------------
func AddClass(elem string, class string) {
	getElementValue(elem, "classList").Call("add", class)
}

// ----------------------------------------------------------------------------
func RemoveClass(elem string, class string) {
	classList := getElementValue(elem, "classList")
	if classList.Call("contains", class).Bool() {
		classList.Call("remove", class)
	}
}

// ----------------------------------------------------------------------------
func wrapGoFunction(fn func()) func(js.Value, []js.Value) interface{} {
	return func(_ js.Value, _ []js.Value) interface{} {
		fn()
		return nil
	}
}

// ----------------------------------------------------------------------------
func AddEventListener(elem string, event string, fn func()) {
	getElementById(elem).Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
}
