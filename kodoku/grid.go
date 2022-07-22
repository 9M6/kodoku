package kodoku

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

// Grid is a grid of tiles.
type Grid struct {
	rows []Row
	lock map[int]map[int]bool
}

// NewGrid returns a new grid.
func NewGrid(x, y int) *Grid {
	r := make(Row, x)
	for i := 0; i < x; i++ {
		r[i] = Tile(T1)
	}

	g := &Grid{
		rows: make([]Row, x),
		lock: make(map[int]map[int]bool, y),
	}
	for i := 0; i < y; i++ {
		g.rows[i] = r
		g.lock[i] = make(map[int]bool, x)
	}

	return g
}

// Clone returns a clone of the grid.
// TODO:
func (g *Grid) Clone() *Grid {
	clone := NewGrid(len(g.rows), len(g.rows[0]))

	copy(clone.rows, g.rows)
	for i, row := range g.lock {
		clone.lock[i] = make(map[int]bool, len(row))
		for j, locked := range row {
			clone.lock[i][j] = locked
		}
	}

	return clone
}

// Fill fills the grid with the given tile.
func (g *Grid) Fill(row, col int, tile Tile) {
	if len(g.rows) > row && len(g.rows[row]) > col {
		locked, ok := g.lock[row][col]
		if !ok || !locked {
			g.rows[row][col] = tile
		}
	}
}

// FillAndLock the grid with the given tile, and locks the variable.
func (g *Grid) FillAndLock(row, col int, tile Tile) {
	g.Fill(row, col, tile)
	g.Lock(row, col, true)
}

// TODO: Replace csvString with csvReader
func (g *Grid) FillFromCSV(csvString string) {
	r := csv.NewReader(strings.NewReader(csvString))

	record, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, row := range record {
		for j, col := range row {
			tile, err := strconv.Atoi(col)
			if err != nil {
				fmt.Println(err)
				return
			}

			g.FillAndLock(i, j, Tile(tile))
		}
	}
}

// Lock locks the grid at the given row and column.
func (g *Grid) Lock(row, col int, locked bool) {
	g.lock[row][col] = locked
}

// IsLocked returns whether the grid at the given row and column is locked.
func (g *Grid) IsLocked(row, col int) bool {
	locked, ok := g.lock[row][col]
	return ok || locked
}

// TODO:
func (g *Grid) UniqGrid() {}

func (g *Grid) UniqRows() {}
func (g *Grid) UniqCols() {}
func (g *Grid) UniqSubs() {}

// String returns the string representation of the grid.
func (g *Grid) String() string {
	var sb strings.Builder
	for _, r := range g.rows {
		sb.WriteString(fmt.Sprintf("%v\n", r))
	}
	return sb.String()
}
