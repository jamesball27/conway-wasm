package main

import (
	"math/rand"
	"syscall/js"
	"time"

	"github.com/jamesball27/conway-wasm/listener"

	"github.com/jamesball27/conway-wasm/canvas"
	"github.com/jamesball27/conway-wasm/conway"
	"github.com/jamesball27/conway-wasm/render"
)

const (
	canvasID = "canvas"
	startID  = "start"
	resetID  = "reset"
	stopID   = "stop"
	height   = 50
	width    = 100
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := canvas.New(canvasID, height, width)
	g := conway.NewGame(height, width)

	r := &render.Renderer{
		Canvas: c,
		Game:   g,
	}
	reset(r)

	listener.New(startID).OnClick(func() {
		start(r, r.Game)
	})
	listener.New(resetID).OnClick(func() {
		reset(r)
	})
	listener.New(stopID).OnClick(func() {
		stop()
	})

	// Ensure Go program never exits
	runner := make(chan bool)
	<-runner
}

func start(r *render.Renderer, g *conway.Game) {
	renderer := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		g.PopulateNextGen()
		r.Render()

		return nil
	})
	js.Global().Set("interval", js.Global().Call("setInterval", renderer, 100))
}

func reset(r *render.Renderer) {
	r.Canvas = canvas.New(canvasID, height, width)
	r.Game = conway.NewGame(height, width)
	r.Render()
}

func stop() {
	js.Global().Call("clearInterval", js.Global().Get("interval"))
}
