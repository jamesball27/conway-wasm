package render

import (
	"github.com/jamesball27/conway-wasm/canvas"
	"github.com/jamesball27/conway-wasm/conway"
)

type Renderer struct {
	Game   *conway.Game
	Canvas canvas.Canvas
}

func (r *Renderer) Render() {
	r.Canvas.Clear()
	r.Canvas.DrawGrid()

	g := r.Game.Grid
	for x := 0; x < g.Height(); x++ {
		for y := 0; y < g.Width(); y++ {
			if (*g)[x][y].IsAlive {
				// Height and width of grid are reversed with regards to to x/y coordinates
				r.Canvas.FillCell(y, x)
			}
		}
	}
}
