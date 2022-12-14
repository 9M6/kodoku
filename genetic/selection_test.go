package genetic

import (
	"testing"
)

func TestKTournamentSelection(t *testing.T) {
	p := NewSpecies(1, 4, 100, []uint8{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})
	g1 := KTournamentSelection(p.curr, 10)
	t.Log(g1)
	t.Log(g1.Fitness())
	g2 := KTournamentSelection(p.curr, 10)
	t.Log(g2)
	t.Log(g2.Fitness())
	g3 := KTournamentSelection(p.curr, 100)
	t.Log(g3.base)
	t.Log(g3.gene)
	t.Log(g3.Fitness())
}

func TestFitnessProportionateSelection(t *testing.T) {
	p := NewSpecies(1, 4, 1000, []uint8{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})
	g1 := FitnessProportionateSelection(p.curr)
	t.Log(g1)
	t.Log(g1.Fitness())
	g2 := FitnessProportionateSelection(p.curr)
	t.Log(g2)
	t.Log(g2.Fitness())
	g3 := FitnessProportionateSelection(p.curr)
	t.Log(g3)
	t.Log(g3.Fitness())
}
