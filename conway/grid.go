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

func (g *Grid) Height() int {
	return len(*g)
}

func (g *Grid) Width() int {
	return len((*g)[0])
}

func (g *Grid) Get(row int, col int) Cell {
	return (*g)[row][col]
}

func (g *Grid) Set(row int, col int, c Cell) {
	(*g)[row][col] = c
}

func (g *Grid) Each(f func(row int, col int)) {
	for row := 0; row < g.Height(); row++ {
		for col := 0; col < g.Width(); col++ {
			f(row, col)
		}
	}
}

func (g *Grid) randomize() *Grid {
	g.Each(func(row int, col int) {
		c := Cell{
			IsAlive: !(rand.Intn(2) == 0),
			row:     row,
			col:     col,
		}
		g.Set(row, col, c)
	})

	return g
}

// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// Any live cell with two or three live neighbours lives on to the next generation.
// Any live cell with more than three live neighbours dies, as if by overpopulation.
// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func (g *Grid) populate() *Grid {
	next := NewGrid(g.Height(), g.Width())

	g.Each(func(row int, col int) {
		cell := g.Get(row, col)
		new := Cell{row: row, col: col}
		count := cell.aliveNeighbors(g)

		if cell.IsAlive {
			if count < 2 || count > 3 {
				new.IsAlive = false
			} else {
				new.IsAlive = true
			}
		} else {
			if count == 3 {
				new.IsAlive = true
			}
		}

		next.Set(row, col, new)
	})

	return next
}
