package genetic

type Population struct {
	candidates []*Genes
}

// NewPopulation creates a new population.
func NewPopulation(min, max, size int, base []uint8) *Population {
	p := &Population{
		candidates: make([]*Genes, size),
	}

	for i := 0; i < size; i++ {
		g := NewGenes(base)
		g.Seed(min, max)
		p.candidates[i] = g
	}

	return p
}
