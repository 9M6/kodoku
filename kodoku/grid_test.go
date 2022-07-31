package kodoku

import (
	"encoding/csv"
	"fmt"
	"strings"
	"testing"

	"kodoku/genetic"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid(4, 4)
	t.Log(g.String())
}

func TestGrid_Fill(t *testing.T) {
	g := NewGrid(4, 4)
	g.Fill(0, 0, 1)
	g.Fill(0, 1, 2)
	g.Fill(0, 2, 3)
	g.Fill(0, 3, 4)
	g.Fill(1, 0, 1)
	g.Fill(1, 1, 2)
	g.Fill(1, 2, 3)
	g.Fill(1, 3, 4)
	g.Fill(2, 0, 1)
	g.Fill(2, 1, 2)
	g.Fill(2, 2, 3)
	g.Fill(2, 3, 4)
	g.Fill(3, 0, 1)
	g.Fill(3, 1, 2)
	g.Fill(3, 2, 3)
	g.Fill(3, 3, 4)
	t.Log(g.String())
	t.Log(fmt.Sprintf("%b", g.rows))
	t.Log(fmt.Sprintf("%b, %b", g.rows[0][1], g.rows[2][2]))
	t.Log(fmt.Sprintf("%b", g.rows[0][1]^g.rows[2][2]))
}

func TestGrid_Import(t *testing.T) {
	gene := genetic.NewGenes(make([]uint8, 16))
	gene.Seed(1, 4)
	grid := NewGrid(4, 4)
	grid.Import(gene.Export())
	t.Log(gene.Export())
	t.Log(grid.rows)
}

func TestGrid_Export(t *testing.T) {
	gene := genetic.NewGenes(make([]uint8, 16))
	gene.Seed(1, 4)
	grid := NewGrid(4, 4)

	text := strings.NewReader("1,2,,\n2,3,1,3\n2,,,1\n1,,,3")
	str, err := csv.NewReader(text).ReadAll()
	if err != nil {
		panic(err)
	}

	grid.FillFromCSV(str)
	t.Log(grid)
	t.Log(grid.rows)

	t.Log(grid.Export())
	gene.Import(grid.Export())
}
