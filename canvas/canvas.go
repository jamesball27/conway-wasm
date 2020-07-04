package canvas

import (
	"fmt"
	"syscall/js"
)

type Canvas struct {
	height int
	width  int
	ctx    js.Value
	canvas js.Value
}

const boxSize = 50

func New(id string, h int, w int) Canvas {
	canvas := js.Global().Get("document").Call("getElementById", id)
	ctx := canvas.Call("getContext", "2d")

	canvas.Set("height", c.height)
	canvas.Set("width", c.width)
	ctx.Call("clearRect", 0, 0, c.width, c.height)

	return Canvas{
		height: h,
		width:  w,
		ctx:    ctx,
		canvas: canvas,
	}
}

func (c *Canvas) Render() {
	c.drawGrid()
}

func (c *Canvas) drawGrid() {
	for x := 0; x <= c.width; x += boxSize {
		fmt.Println(x)
		c.ctx.Call("moveTo", x, 0)
		c.ctx.Call("lineTo", x, c.height)
	}

	for y := 0; y <= c.height; y += boxSize {
		c.ctx.Call("moveTo", 0, y)
		c.ctx.Call("lineTo", c.width, y)
	}

	c.ctx.Set("strokeStyle", "black")
	c.ctx.Call("stroke")
}
