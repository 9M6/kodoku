package genetic

import "kodoku/kodoku"

type Population struct {
	candidates []kodoku.Grid
}

// New creates a new population.
func (p *Population) New(Nw int) *Population {
	return p
}

// Seed a new population.
func (p *Population) Seed(Nw int) *Population {
	return p
}
