package canvas

import (
	"syscall/js"
)

type Canvas struct {
	height int
	width  int
	ctx    js.Value
}

const boxSize = 10

func New(id string, h int, w int) Canvas {
	canvas := js.Global().Get("document").Call("getElementById", id)
	ctx := canvas.Call("getContext", "2d")

	height := h * boxSize
	width := w * boxSize

	canvas.Set("height", height)
	canvas.Set("width", width)
	ctx.Call("clearRect", 0, 0, height, width)
	ctx.Set("fillStyle", "black")

	c := Canvas{
		height: height,
		width:  width,
		ctx:    ctx,
	}

	c.drawGrid()

	return c
}

func (c *Canvas) FillCell(x int, y int) {
	c.ctx.Call("fillRect", x*boxSize, y*boxSize, boxSize, boxSize)
}

func (c *Canvas) drawGrid() {
	for x := 0; x <= c.width; x += boxSize {
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
