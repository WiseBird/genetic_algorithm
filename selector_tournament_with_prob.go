package genetic_algorithm

import (
	"math"
	"math/rand"
	"sort"
)

// Selects several random individuals from population and then runs tournament among them.
// See http://en.wikipedia.org/wiki/Tournament_selection
type TournamentSelector struct {
	*SelectorBase

	probability float64
	contestants int
}

// Constructor for tournament selector.
//
// Probability specify the probability that the best individual will be selected.
// Must be equal or grater than 0.5
func NewTournamentSelector(probability float64, contestants int) *TournamentSelector {
	if probability < 0.5 || probability > 1 {
		panic("Probability out of range")
	}
	if contestants < 1 {
		panic("Must be at least one contestant")
	}

	selector := new(TournamentSelector)

	selector.SelectorBase = NewSelectorBase(selector)
	selector.probability = probability
	selector.contestants = contestants

	return selector
}
func (selector *TournamentSelector) Select() ChromosomeInterface {
	return selector.population[selector.SelectInd()]
}
func (selector *TournamentSelector) SelectInd() int {
	contestants := make([]int, selector.contestants)

	for i := 0; i < selector.contestants; i++ {
		contestants[i] = rand.Intn(len(selector.population))
	}
	sort.Sort(sort.IntSlice(contestants))

	r := rand.Float64()
	prob := float64(0)
	for i := 0; i < selector.contestants - 1; i++ {
		prob += selector.ithProbability(i)
		if r < prob {
			return contestants[i]
		}
	}
	return contestants[selector.contestants - 1]
}
func (selector *TournamentSelector) ithProbability(i int) float64 {
	return selector.probability * math.Pow((float64(1) - selector.probability), float64(i))
}