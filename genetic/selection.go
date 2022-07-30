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
func KTournamentSelection(p *Population, iter int) *Genes {
	var best *Genes
	for i := 0; i < iter; i++ {
		g := p.candidates[rand.Intn(len(p.candidates))]
		if best == nil || g.Fitness() > best.Fitness() {
			best = g
		}
	}
	return best
}

// FitnessProportionateSelection selects a candidate from a population
// using a fitness proportionate selection.
func FitnessProportionateSelection(p *Population) *Genes {
	var total float64
	for _, g := range p.candidates {
		total += g.Fitness()
	}

	var selected *Genes
	pick := rand.Float64() * total
	for _, g := range p.candidates {
		pick -= g.Fitness()
		if pick <= 0 {
			selected = g
			break
		}
	}
	return selected
}
