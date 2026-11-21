package rules

type Stone uint8

const (
	Empty Stone = iota
	Black
	White
)

type Board struct {
	Size int
	Grid [][]Stone
}

func NewBoard(size int) *Board {
	grid := make([][]Stone, size)
	for i := range grid {
		grid[i] = make([]Stone, size)
	}
	return &Board{Size: size, Grid: grid}
}
