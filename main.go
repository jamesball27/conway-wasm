package main

import (
	"github.com/jamesball27/conway-wasm/canvas"
)

const (
	canvasID     = "canvas"
	canvasHeight = 400
	canvasWidth  = 1000
)

func main() {
	c := canvas.New(canvasID, canvasHeight, canvasWidth)
	c.Render()
}
