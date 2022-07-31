package genetic

import (
	"testing"
)

func TestNewPopulation(t *testing.T) {
	var p *Species
	var base = []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}

	p = NewSpecies(1, 4, 10, base)
	if len(p.curr) != 10 {
		t.Errorf("Expected 10 candidates, got %d", len(p.curr))
		t.Log(p.curr)
	}
}
