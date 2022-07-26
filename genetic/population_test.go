package genetic

import (
	"testing"
)

func TestNewPopulation(t *testing.T) {
	var p *Population
	var base = []uint8{1, EMPTY, 3, 4, EMPTY, 2, 3, EMPTY, 1, 4, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}

	p = NewPopulation(1, 4, 10, base)
	if len(p.candidates) != 10 {
		t.Errorf("Expected 10 candidates, got %d", len(p.candidates))
		t.Log(p.candidates)
	}
}
