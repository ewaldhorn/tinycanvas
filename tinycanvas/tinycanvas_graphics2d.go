package tinycanvas

// Sometimes you just need proper access to the canvas, so, here we go.

import "syscall/js"

// ----------------------------------------------------------------------------
//
// Context2D is a convenience wrapper around the Graphics2D context object.
// It's useful for using the HTML Canvas directly instead of the buffered image.
type Context2D struct {
	ctx *js.Value
}

// ----------------------------------------------------------------------------
// Sometimes you need access to the Context.
func (t *TinyCanvas) GetContext() *Context2D {
	return &Context2D{&t.ctx}
}

// ---------------------------------------------------------- GENERIC FUNCTIONS
// ----------------------------------------------------------------------------
// Generally be able to Set something
func (c *Context2D) Set(what string, toWhat any) {
	c.ctx.Set(what, toWhat)
}

// ----------------------------------------------------------------------------
// Generally be able to Call something
func (c *Context2D) Call(what string, toWhat ...any) {
	c.ctx.Call(what, toWhat...)
}

// --------------------------------------------------------- GRAPHICS FUNCTIONS
// ----------------------------------------------------------------------------
// arc supports a counterclockwise parameter now
func (c *Context2D) Arc(xPos, yPos, radius, startAngle, endAngle, counterclockwise any) {
	c.Call("arc", xPos, yPos, radius, startAngle, endAngle, counterclockwise)
}

// ----------------------------------------------------------------------------
func (c *Context2D) BeginPath() {
	c.Call("beginPath")
}

// ----------------------------------------------------------------------------
func (c *Context2D) ClearRect(x, y, width, height any) {
	c.Call("clearRect", x, y, width, height)
}

// ----------------------------------------------------------------------------
func (c *Context2D) DrawImage(img, x, y any) {
	c.Call("drawImage", img, x, y)
}

// ----------------------------------------------------------------------------
func (c *Context2D) Fill() {
	c.Call("fill")
}

// ----------------------------------------------------------------------------
func (c *Context2D) FillRect(x, y, width, height any) {
	c.Call("fillRect", x, y, width, height)
}

// ----------------------------------------------------------------------------
func (c *Context2D) FillStyle(style string) {
	c.Set("fillStyle", style)
}

// ----------------------------------------------------------------------------
func (c *Context2D) GlobalCompositeOperation(value any) {
	c.Set("globalCompositeOperation", value)
}

// ----------------------------------------------------------------------------
func (c *Context2D) LineTo(x, y any) {
	c.Call("lineTo", x, y)
}

// ----------------------------------------------------------------------------
func (c *Context2D) LineWidth(width int) {
	c.Set("lineWidth", width)
}

// ----------------------------------------------------------------------------
func (c *Context2D) Restore() {
	c.Call("restore")
}

// ----------------------------------------------------------------------------
func (c *Context2D) Rotate(angle any) {
	c.Call("rotate", angle)
}

// ----------------------------------------------------------------------------
func (c *Context2D) Save() {
	c.Call("save")
}

// ----------------------------------------------------------------------------
func (c *Context2D) Stroke() {
	c.Call("stroke")
}

// ----------------------------------------------------------------------------
func (c *Context2D) StrokeStyle(style string) {
	c.Set("strokeStyle", style)
}

// ----------------------------------------------------------------------------
func (c *Context2D) Translate(x, y any) {
	c.Call("translate", x, y)
}
