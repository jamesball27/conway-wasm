package render

import (
	"github.com/jamesball27/conway-wasm/canvas"
	"github.com/jamesball27/conway-wasm/conway"
)

type Renderer struct {
	Game   conway.Game
	Canvas canvas.Canvas
}

func (r *Renderer) Render() {
	g := r.Game.CurrentGen

	for x := 0; x < len(g); x++ {
		for y := 0; y < len(g[x]); y++ {
			if g[x][y].IsAlive {
				r.Canvas.FillCell(y, x)
			}
		}
	}
}
