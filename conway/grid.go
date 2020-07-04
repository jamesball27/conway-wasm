package conway

import (
	"math/rand"
)

type Cell struct {
	IsAlive bool
}

type Grid [][]Cell

func NewGrid(h int, w int) Grid {
	g := make(Grid, h)
	for i := 0; i < h; i++ {
		g[i] = make([]Cell, w)
	}

	return g
}

func (g Grid) randomize() Grid {
	for x := 0; x < len(g); x++ {
		for y := 0; y < len(g[x]); y++ {
			c := Cell{
				IsAlive: !(rand.Intn(2) == 0),
			}
			g[x][y] = c
		}
	}
	return g
}
