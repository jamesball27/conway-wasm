package render

import (
	"github.com/jamesball27/conway-wasm/canvas"
	"github.com/jamesball27/conway-wasm/conway"
)

type Renderer struct {
	Game   *conway.Game
	Canvas *canvas.Canvas
}

func (r *Renderer) Render() {
	r.Canvas.Clear()
	r.Canvas.DrawGrid()

	g := r.Game.Grid
	g.Each(func(row int, col int) {
		if g.Get(row, col).IsAlive {
			// Rows/cols of grid are reversed with regards to x/y coordinates
			r.Canvas.FillCell(col, row)
		}
	})
}
