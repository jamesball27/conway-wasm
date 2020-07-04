package main

import (
	"math/rand"
	"time"

	"github.com/jamesball27/conway-wasm/canvas"
	"github.com/jamesball27/conway-wasm/conway"
	"github.com/jamesball27/conway-wasm/render"
)

const (
	canvasID = "canvas"
	height   = 100
	width    = 200
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := canvas.New(canvasID, height, width)
	g := conway.NewGame(height, width)

	r := render.Renderer{
		Canvas: c,
		Game:   g,
	}

	r.Render()
}
