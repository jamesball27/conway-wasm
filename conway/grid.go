package conway

import (
	"math/rand"
)

type Grid [][]Cell

func NewGrid(h int, w int) *Grid {
	g := make(Grid, h)
	for i := 0; i < h; i++ {
		g[i] = make([]Cell, w)
	}

	return &g
}

func (g *Grid) randomize() *Grid {
	for x := 0; x < g.Height(); x++ {
		for y := 0; y < g.Width(); y++ {
			c := Cell{
				x:       x,
				y:       y,
				IsAlive: !(rand.Intn(2) == 0),
			}
			(*g)[x][y] = c
		}
	}
	return g
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Width() int {
	return len(g[0])
}

// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// Any live cell with two or three live neighbours lives on to the next generation.
// Any live cell with more than three live neighbours dies, as if by overpopulation.
// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func (g *Grid) populate() *Grid {
	next := NewGrid(g.Height(), g.Width())

	for x := 0; x < g.Height(); x++ {
		for y := 0; y < g.Width(); y++ {
			cell := (*g)[x][y]
			new := Cell{x: x, y: y}
			alive := cell.aliveNeighbors(*g)

			if cell.IsAlive {
				if alive < 2 || alive > 3 {
					new.IsAlive = false
				} else {
					new.IsAlive = true
				}
			} else {
				if alive == 3 {
					new.IsAlive = true
				}
			}

			(*next)[x][y] = new
		}
	}

	return next
}
