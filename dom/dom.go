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
func GetElementById(elem string) js.Value {
	return document.Call("getElementById", elem)
}

// ----------------------------------------------------------------------------
func GetElementValue(elem string, value string) js.Value {
	return GetElementById(elem).Get(value)
}

// ----------------------------------------------------------------------------
func Hide(elem string) {
	GetElementValue(elem, "style").Call("setProperty", "display", "none")
}

// ----------------------------------------------------------------------------
func Show(elem string) {
	GetElementValue(elem, "style").Call("setProperty", "display", "block")
}

// ----------------------------------------------------------------------------
func SetFocus(elem string) {
	GetElementById(elem).Call("focus")
}

// ----------------------------------------------------------------------------
func GetString(elem string, value string) string {
	return GetElementValue(elem, value).String()
}

// ----------------------------------------------------------------------------
func SetValue(elem string, key string, value string) {
	GetElementById(elem).Set(key, value)
}

// ----------------------------------------------------------------------------
func AddClass(elem string, class string) {
	GetElementValue(elem, "classList").Call("add", class)
}

// ----------------------------------------------------------------------------
func RemoveClass(elem string, class string) {
	classList := GetElementValue(elem, "classList")
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
	GetElementById(elem).Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
}
