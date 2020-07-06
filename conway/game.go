package conway

type Game struct {
	Grid *Grid
}

func NewGame(h int, w int) *Game {
	g := NewGrid(h, w).randomize()

	return &Game{Grid: g}
}

func (g *Game) PopulateNextGen() {
	g.Grid = g.Grid.populate()
}
