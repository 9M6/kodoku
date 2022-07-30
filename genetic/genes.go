package genetic

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	EMPTY = 0x00
)

// Genes is a struct that holds a gene which is an array of uint8,
// we crease a 1 Dimensional array that represents the solution to
// the puzzle.
//
// The puzzle in our problem is a 4x4 matrix. The matrix is represented
// by a 1 Dimensional array with a dynamic length in our case, but given
// the size of the matrix, we calculate the length to 16 words.
//
// Each word is Encoded as a number from 1 to 4, and with 0 being
// an empty, or nil placeholder.
//
// When we crease a Gene, we pre-fill it with the base solution that
// the user has initially provided.
type Genes struct {
	// Gene structure
	gene []uint8
	lock map[int]bool

	// Fitness structure
	dirty   bool
	fitness float64
}

// NewGenes creates a new genes, and prefills the gene array with the
// base solution provided by the user.
func NewGenes(base []uint8) *Genes {
	var g = &Genes{
		gene: make([]uint8, len(base)),
		lock: make(map[int]bool),
	}

	for i := 0; i < len(base); i++ {
		g.FillAndLock(i, base[i])
	}

	return g
}

// Seed function fills the genes with a random values from min, max range.
func (g *Genes) Seed(min, max int) {
	for i := 0; i < len(g.gene); i++ {
		g.Fill(i, g.Rand(min, max))
	}
}

// Rand returns a random number between min and max.
func (g *Genes) Rand(min, max int) uint8 {
	return uint8(rand.Intn(max-min+1) + min)
}

// Fill fills the gene at the given index with the given value.
func (g *Genes) Fill(i int, v uint8) {
	if !g.IsLocked(i) || v == EMPTY {
		g.gene[i] = v
	}
}

// FillAndLock fills the gene at the given index with the given value,
// and locks the gene at the given index.
func (g *Genes) FillAndLock(i int, v uint8) {
	if !g.IsLocked(i) && v != EMPTY {
		g.gene[i] = v
		g.lock[i] = true
	}
}

// IsLocked returns true if the gene at the given index is locked.
func (g *Genes) IsLocked(i int) bool {
	b, ok := g.lock[i]
	return b && ok
}

// Rows returns the genes as a matrix chunking the elements given the size.
//  	- size is the number of items.
func (g *Genes) Rows(size int) [][]uint8 {
	var rows [][]uint8
	for i := 0; i < len(g.gene); i += size {
		end := i + size

		if end > len(g.gene) {
			end = len(g.gene)
		}

		rows = append(rows, g.gene[i:end])
	}

	return rows
}

// RowScore returns the score of the fitness of the row, it does so by
// counting the number of duplicates in each row, and then it returns
// the number of duplicates and the percentage of duplicates.
func (g *Genes) RowScore() (int, int) {
	score := make(map[int]map[uint8]int, 0)

	rows := g.Rows(4)
	for i := 0; i < len(rows); i++ {
		score[i] = make(map[uint8]int, 0)
		for j := 0; j < len(rows[i]); j++ {
			score[i][rows[i][j]]++
		}
	}

	var duple int
	for i := 0; i < len(score); i++ {
		for _, s := range score[i] {
			if s > 1 {
				duple += s
			}
		}
	}

	return duple, len(rows) * len(rows[0])
}

// Cols returns the genes as a matrix which has been transposed given a matrix.
// 		- size is the number of columns.
func (g *Genes) Cols(size int) [][]uint8 {
	rows := g.Rows(size)

	xl := len(rows[0])
	yl := len(rows)
	cols := make([][]uint8, xl)

	for i := range cols {
		cols[i] = make([]uint8, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			cols[i][j] = rows[j][i]
		}
	}

	return cols
}

// ColScore returns the score of the fitness of the column, it does so by
// counting the number of duplicates in each column, and then it returns
// the number of duplicates and the percentage of duplicates.
func (g *Genes) ColScore() (int, int) {
	score := make(map[int]map[uint8]int, 0)

	cols := g.Rows(4)
	for i := 0; i < len(cols); i++ {
		score[i] = make(map[uint8]int, 0)
		for j := 0; j < len(cols[i]); j++ {
			score[i][cols[i][j]]++
		}
	}

	var duple int
	for i := 0; i < len(score); i++ {
		for _, s := range score[i] {
			if s > 1 {
				duple += s
			}
		}
	}

	return duple, len(cols) * len(cols[0])
}

// Subs returns the genes as a matrix.
func (g *Genes) Subs(size int) [][]uint8 {
	rows := g.Rows(2)
	subs := make([][]uint8, size)
	step := size / 2

	for i := 0; i < size; i++ {
		subs[i] = make([]uint8, 0)
		for j := 0; j < step; j++ {
			if i < step {
				subs[i] = append(subs[i], rows[i+step*j]...)
			} else {
				subs[i] = append(subs[i], rows[step+i+step*j]...)
			}
		}
	}

	return subs
}

// SubScore returns the score of the fitness of the sub.
func (g *Genes) SubScore() (int, int) {
	score := make(map[int]map[uint8]int, 0)

	subs := g.Subs(4)
	for i := 0; i < len(subs); i++ {
		score[i] = make(map[uint8]int, 0)
		for j := 0; j < len(subs[i]); j++ {
			score[i][subs[i][j]]++
		}
	}

	var duple int
	for i := 0; i < len(score); i++ {
		for _, s := range score[i] {
			if s > 1 {
				duple += s
			}
		}
	}

	return duple, len(subs) * len(subs[0])
}

// Fitness returns the score of the fitness of the genes, it does so by
// calculating both row and column, and subgrid scores, and then it
// returns the sum of all scores.
//
// The score calculates simply the number of duplicates in each row, column,
// and subgrid, and then it returns the sum of all scores, but to calculate
// the fitness, we have to subtract it, since that will give us the % of
// non duplicates words, and thus how close we are to having all unique words.
func (g *Genes) Fitness() float64 {
	if g.fitness != 0.0 && !g.dirty {
		return g.fitness
	}

	rd, rs := g.RowScore()
	cd, cs := g.ColScore()
	sd, ss := g.SubScore()

	return 1 - (float64(rd+cd+sd) / float64(rs+cs+ss))
}

// Mutate mutates the genes, it does so by filling the genes with a random
// value from min, max range, the function also has a probability variable,
// 'u' given that the function will actually mutate/swap the genes.
func (g *Genes) Mutate(u float32) {
	gene := g.Export()

	for i := 0; i < len(gene); i++ {
		if rand.Float32() < u {
		RAND:
			rnd := g.Rand(0, len(gene)-1)
			if gene[rnd] == EMPTY {
				goto RAND
			}

			if gene[i] == EMPTY {
				continue
			}

			gene[i], gene[rnd] = gene[rnd], gene[i]
		}
	}

	g.Import(gene)
}

// Import imports the genes from a slice, it does so by filling the genes with
// the values from the slice, by not overwriting the locked genes.
func (g *Genes) Import(gene []uint8) {
	g.dirty = true
	for i := 0; i < len(gene); i++ {
		if gene[i] != EMPTY {
			g.Fill(i, gene[i])
		}
	}
}

// Export exports the genes to a slice, it does so by returning the genes, as
// a slice of uint8, the genes returned are the ones that are unlocked.
func (g *Genes) Export() []uint8 {
	gene := make([]uint8, len(g.gene))
	for i := 0; i < len(g.gene); i++ {
		if g.IsLocked(i) {
			gene[i] = EMPTY
			continue
		}
		gene[i] = g.gene[i]
	}
	return gene
}
