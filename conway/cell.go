package conway

type Cell struct {
	IsAlive bool
	x       int
	y       int
}

func (c Cell) aliveNeighbors(g Grid) int {
	count := 0
	for _, n := range c.neighbors(g) {
		if n.IsAlive {
			count++
		}
	}
	return count
}

func (c Cell) neighbors(g Grid) []Cell {
	x := c.x
	y := c.y
	// -1, -1
	// -1 0
	// -1 1
	return []Cell{
		// g[x-1][y-1],
		// g[x-1][y],
		// g[x-1][y+1],
		// g[x][y-1],
		// g[x][y+1],
		// g[x+1][y-1],
		// g[x+1][y],
		// g[x+1][y+1],
		g[(x-1+g.Height())%g.Height()][(y-1+g.Width())%g.Width()],
		g[(x-1+g.Height())%g.Height()][y],
		g[(x-1+g.Height())%g.Height()][(y+1+g.Width())%g.Width()],
		g[x][(y-1+g.Width())%g.Width()],
		g[x][(y+1+g.Width())%g.Width()],
		g[(x+1+g.Height())%g.Height()][(y-1+g.Width())%g.Width()],
		g[(x+1+g.Height())%g.Height()][y],
		g[(x+1+g.Height())%g.Height()][(y+1+g.Width())%g.Width()],
	}
}
