package genetic

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestNewGenes(t *testing.T) {
	g := NewGenes([]uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})
	t.Log(g)
	if g.gene[0] != 1 {
		t.Errorf("expected gene[0] to be 1, got %d", g.gene[0])
	}
	if g.gene[7] != EMPTY {
		t.Errorf("expected gene[7] to be 1, got %d", g.gene[7])
	}
	if g.gene[9] != 4 {
		t.Errorf("expected gene[9] to be 10, got %d", g.gene[9])
	}
}

func TestGenes_Generate(t *testing.T) {
	g := NewGenes([]uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})
	g.Seed(1, 4)
	if g.gene[0] != 1 {
		t.Errorf("expected gene[0] to be 1, got %d", g.gene[0])
	}
	if g.gene[7] == EMPTY {
		t.Errorf("expected gene[7] to be 1, got %d", g.gene[7])
	}
	if g.gene[9] != 4 {
		t.Errorf("expected gene[9] to be 10, got %d", g.gene[9])
	}
}

func TestGenes_Mutate(t *testing.T) {
	ar := []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	g := NewGenes(ar)
	g.Seed(1, 4)
	t.Log(g.gene)
	g.Mutate(1)
	t.Log(g.gene)
}

func TestGenes_Rows(t *testing.T) {
	g := NewGenes([]uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})
	g.Seed(1, 4)
	t.Log(g.gene)
	t.Log(g.Rows(4))
}

func TestGenes_RowScore(t *testing.T) {
	ar := []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	g := NewGenes(ar)
	g.Seed(1, 4)
	t.Log(g.gene)
	t.Log(g.RowScore())
}

func TestGenes_Cols(t *testing.T) {
	g := NewGenes([]uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 2, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})
	g.Seed(1, 4)
	t.Log(g.gene)
	t.Log(g.Cols(4))
}

func TestGenes_ColScore(t *testing.T) {
	ar := []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	g := NewGenes(ar)
	g.Seed(1, 4)
	t.Log(g.gene)
	t.Log(g.ColScore())
}

func TestGenes_Fitness(t *testing.T) {
	ar := []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	g := NewGenes(ar)
	g.Seed(1, 4)
	t.Log(g.gene)
	t.Log(1 - g.Fitness())

}

func TestGenes_Subs(t *testing.T) {
	ar := []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	g := NewGenes(ar)
	g.Seed(1, 4)
	t.Log(g.gene)
	t.Log(g.Rows(4))
	t.Log(g.Subs(4))
}

func TestGenes_SubScore(t *testing.T) {
	ar := []uint8{3, 4, 2, 1, 1, 2, 3, 4, 2, 1, 3, 4, 3, 4, 2, 1}
	g := NewGenes(ar)
	g.Seed(1, 4)
	t.Log(g.gene)
	t.Log(g.SubScore())
}

func TestGenes_Import(t *testing.T) {
	ar := []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	g := NewGenes(ar)
	g.Seed(1, 4)
	t.Log(g.gene)
	g.Import([]uint8{EMPTY, 3, EMPTY, EMPTY, 2, EMPTY, EMPTY, 1, EMPTY, EMPTY, 1, 3, 2, 3, 2, 4})
	t.Log(g.gene)
}

func TestGenes_Export(t *testing.T) {
	ar := []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	g := NewGenes(ar)
	g.Seed(1, 4)
	t.Log(g.gene)
	t.Log(g.Export())
}

func TestGenes_CrossOver(t *testing.T) {
	ch := []uint8{1, EMPTY, 3, EMPTY, 1, 4, 3, EMPTY, 1, EMPTY, EMPTY, EMPTY, EMPTY, 3, 1, EMPTY}
	g := NewGenes(ch)
	p1 := []uint8{1, EMPTY, 3, EMPTY, 1, 4, 3, EMPTY, 1, EMPTY, EMPTY, EMPTY, EMPTY, 3, 1, EMPTY}
	p2 := []uint8{3, EMPTY, 1, EMPTY, EMPTY, EMPTY, 1, EMPTY, 3, EMPTY, 1, 4, EMPTY, 3, 1, EMPTY}
	g1 := NewGenes(p1)
	g1.Seed(1, 4)
	g2 := NewGenes(p2)
	g2.Seed(1, 4)
	t.Log(g1.gene)
	t.Log(g2.gene)
	g.CrossOver(0.5, g1, g2)
	t.Log(g.gene)
}
