package conway

type Cell struct {
	IsAlive bool
	row     int
	col     int
}

func (c Cell) aliveNeighbors(g *Grid) int {
	count := 0
	for _, n := range c.neighbors(g) {
		if n.IsAlive {
			count++
		}
	}
	return count
}

func (c Cell) neighbors(g *Grid) []Cell {
	cells := []Cell{}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Do not consider self as a neighbor
			if i == 0 && j == 0 {
				continue
			}

			// Wrap around grid if index out of bounds
			wrapRow := (c.row + i + g.Height()) % g.Height()
			wrapCol := (c.col + j + g.Width()) % g.Width()

			cells = append(cells, g.Get(wrapRow, wrapCol))
		}
	}

	return cells
}
