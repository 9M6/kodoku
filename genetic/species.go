package genetic

type Species struct {
	next []*Genes
	curr []*Genes
}

// NewSpecies creates a new Species for Breeding.
func NewSpecies(min, max, size int, base []uint8) *Species {
	s := &Species{
		next: make([]*Genes, 0),
		curr: make([]*Genes, size),
	}

	for i := 0; i < size; i++ {
		g := NewGenes(base)
		g.Seed(min, max)
		s.curr[i] = g
	}

	return s
}

// Breed generates a new pair of the species, given two parents and a child.
func (s *Species) Breed(cx, mt float64, base []uint8, fn func([]*Genes) *Genes) (*Genes, *Genes) {
	mother := fn(s.curr)
	father := fn(s.curr)
	childs := NewGenes(base)
	childs.CrossOver(cx, mother, father)
	mother.Mutate(mt)
	father.Mutate(mt)
	childs.Mutate(mt * 0.1)
	parent := fn([]*Genes{mother, father})
	return childs, parent
}

// Push generates a new generation of the species.
func (s *Species) Push(g *Genes) {
	s.next = append(s.next, g)
}

// Next generates the next generation of the species.
func (s *Species) Next() bool {
	return len(s.next) < len(s.curr)
}

// Purge removes the current generation from the species.
func (s *Species) Purge() []*Genes {
	history := s.curr
	s.curr = s.next
	s.next = make([]*Genes, 0)
	return history
}
