package genetic

import "fmt"

type Generations struct {
	// Base is the user input puzzle
	base []uint8

	// Min and max are the minimum and maximum value for the random
	min, max int

	// Pop is the population size
	pop int

	// Gen is the number of generations
	gen int

	// Iter is the number of iterations per selection
	iter int

	// Mut and cross are the mutation and crossover rate
	mut   float64
	cross float64

	species  *Species
	solution *Genes
	history  []*Genes
}

func New(min, max, pop, gen, iter int, mut, cross float64, base []uint8) *Generations {
	return &Generations{
		base:    base,
		min:     min,
		max:     max,
		mut:     mut,
		cross:   cross,
		pop:     pop,
		gen:     gen,
		iter:    iter,
		species: NewSpecies(min, max, pop, base),
	}
}

// Solve searches for a solution and returns it.
func (g *Generations) Solve() *Genes {
	for g.Next() {
		for g.species.Next() {
			childs, parent := g.species.Breed(g.cross, g.mut, g.base, func(s []*Genes) *Genes {
				return KTournamentSelection(s, g.iter)
			})

			fmt.Printf("\033[2K\r\r%v generations", g.gen)

			g.species.Push(parent)
			g.species.Push(childs)

			if parent.Fitness() > childs.Fitness() {
				g.solution = parent
			} else {
				g.solution = childs
			}

			if g.gen == 0 {
				return g.solution
			}

			if g.solution.Fitness() == 1 {
				g.Stop()
				return g.solution
			}
		}
		g.history = append(g.history, g.species.Purge()...)
	}
	return g.solution
}

// Next decrements the generation counter.
func (g *Generations) Next() bool {
	g.gen--
	return g.gen > 0
}

// Stop stops the genetic algorithm.
func (g *Generations) Stop() {
	g.gen = 0
}
