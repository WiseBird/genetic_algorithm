package genetic_algorithm

import (
	"math/rand"
)

// Selects several random individuals from population and then selectd best.
type SimpleTournamentSelector struct {
	*SelectorBase

	contestants int
}

// Constructor for simple tournament selector.
// When contestants = 1 behaves like random selector
func NewSimpleTournamentSelector(contestants int) *SimpleTournamentSelector {
	if contestants < 1 {
		panic("Must be at least one contestant")
	}

	selector := new(SimpleTournamentSelector)

	selector.SelectorBase = NewSelectorBase(selector)
	selector.contestants = contestants

	return selector
}
func (selector *SimpleTournamentSelector) Select() ChromosomeInterface {
	return selector.population[selector.SelectInd()]
}
func (selector *SimpleTournamentSelector) SelectInd() int {
	var bestCost float64
	bestInd := -1

	for i := 0; i < selector.contestants; i++ {
		ind := rand.Intn(len(selector.population))
		chrom := selector.population[ind]

		if bestInd == -1 || chrom.Cost() < bestCost {
			bestCost = chrom.Cost()
			bestInd = ind
		}
	}
	return bestInd
}