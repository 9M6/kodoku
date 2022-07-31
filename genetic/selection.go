package genetic

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// KTournamentSelection selects a candidate from a population
// using a k-tournament selection.
func KTournamentSelection(p []*Genes, iter int) *Genes {
	var best *Genes
	for i := 0; i < iter; i++ {
		g := p[rand.Intn(len(p))]
		if best == nil || g.Fitness() > best.Fitness() {
			best = g
		}
	}
	return best
}

// FitnessProportionateSelection selects a candidate from a population
// using a fitness proportionate selection.
func FitnessProportionateSelection(p []*Genes) *Genes {
	var total float64
	for _, g := range p {
		total += g.Fitness()
	}

	var selected *Genes
	pick := rand.Float64() * total
	for _, g := range p {
		pick -= g.Fitness()
		if pick <= 0 {
			selected = g
			break
		}
	}
	return selected
}
