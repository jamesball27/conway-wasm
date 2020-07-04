package conway

type Game struct {
	CurrentGen Grid
	NextGen    Grid
}

func NewGame(h int, w int) Game {
	g := NewGrid(h, w)
	return Game{
		CurrentGen: g.randomize(),
		NextGen:    nil,
	}
}
