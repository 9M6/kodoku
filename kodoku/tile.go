package kodoku

const (
	T1 = iota
	T2
	T3
	T4
)

// Tile is a tile in the grid.
type Tile uint8

func (t Tile) String() string {
	switch t {
	case T1:
		return "W"
	case T2:
		return "O"
	case T3:
		return "R"
	case T4:
		return "D"
	default:
		return "-"
	}
}
