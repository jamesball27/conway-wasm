package main

import (
	"math/rand"
	"syscall/js"
	"time"

	"github.com/jamesball27/conway-wasm/canvas"
	"github.com/jamesball27/conway-wasm/conway"
	"github.com/jamesball27/conway-wasm/render"
)

const (
	canvasID = "canvas"
	height   = 50
	width    = 100
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := canvas.New(canvasID, height, width)
	g := conway.NewGame(height, width)

	r := render.Renderer{
		Canvas: c,
		Game:   g,
	}

	renderer := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		r.Render()
		g.PopulateNextGen()

		return nil
	})
	window := js.Global()
	window.Call("setInterval", renderer, 100)

	// Ensure Go program doesn't exit
	runner := make(chan bool)
	<-runner
}
