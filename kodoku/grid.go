package kodoku

import (
	"fmt"
	"strings"

	"kodoku/genetic"
)

// Grid is a grid of tiles.
type Grid struct {
	rows [][]uint8
	lock map[int]map[int]bool
	code map[uint8]string
}

// NewGrid returns a new grid.
func NewGrid(x, y int) *Grid {
	g := &Grid{
		rows: make([][]uint8, x),
		code: make(map[uint8]string),
	}

	for i := 0; i < x; i++ {
		g.rows[i] = make([]uint8, y)
	}

	return g
}

// String returns the string representation of the grid.
func (g *Grid) String() string {
	var rb strings.Builder
	for _, r := range g.rows {
		var cb strings.Builder
		for _, v := range r {
			cb.WriteString(fmt.Sprintf(" %s ", g.Decode(v)))
		}
		rb.WriteString(fmt.Sprintf("\n%v", cb.String()))
	}
	return rb.String()
}

// Fill fills the grid with the given tile.
func (g *Grid) Fill(row, col int, v uint8) {
	if len(g.rows) > row && len(g.rows[row]) > col {
		locked, ok := g.lock[row][col]
		if !ok || !locked {
			g.rows[row][col] = v
		}
	}
}

func (g *Grid) FillFromCSV(csv [][]string) {
	for i, row := range csv {
		for j, col := range row {
			g.Fill(i, j, g.Encode(col))
		}
	}
}

func (g *Grid) Encode(v string) uint8 {
	if v == "-" || v == "" {
		return 0
	}

	var max uint8
	for i := range g.code {
		if g.code[i] == v {
			return i
		}

		if i > max {
			max = i
		}
	}

	max = max + 1
	g.code[max] = v

	return max
}

func (g *Grid) Decode(v uint8) string {
	if v == genetic.EMPTY {
		return "-"
	}
	return g.code[v]
}

func (g *Grid) Import(p []uint8) {
	x := len(g.rows)
	y := len(g.rows[0])
	r := make([]uint8, x)

	for i := 0; i < x; i++ {
		r, p = p[:y], p[y:]
		for j := 0; j < y; j++ {
			if r[j] != genetic.EMPTY {
				g.Fill(i, j, r[j])
			}
		}
	}
}

func (g *Grid) Export() []uint8 {
	p := make([]uint8, len(g.rows)*len(g.rows[0]))
	for i := 0; i < len(g.rows); i++ {
		for j := 0; j < len(g.rows[i]); j++ {
			p[i*len(g.rows[i])+j] = g.rows[i][j]
		}
	}
	return p
}
