package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math/rand"
)

// Selects individual with probability proportional to it fitness value.
// Warning! In order to use this selector cost value must be normalized, i.e. chromosome with cost=0 is the best solution.
type RouletteWheelRankWeightingSelector struct {
	*SelectorBase

	weights []float64
}

func NewRouletteWheelRankWeightingSelector() *RouletteWheelRankWeightingSelector {
	selector := new(RouletteWheelRankWeightingSelector)

	selector.SelectorBase = NewSelectorBase(selector)

	return selector
}
func (selector *RouletteWheelRankWeightingSelector) Prepare(population Chromosomes) {
	log.Tracef("Preparing")

	selector.SelectorBase.Prepare(population)

	if selector.weights == nil || len(selector.weights) != len(population) {
		selector.recalcWeights(len(population))
	}
}
func (selector *RouletteWheelRankWeightingSelector) recalcWeights(n int) {
	selector.weights = make([]float64, n)

	sum := float64(n * (n + 1) / 2)
	for i := 1; i <= n; i++ {
		selector.weights[i-1] = float64(n + 1 - i) / sum
	}

	log.Tracef("Recalced weights %v\n", selector.weights)
}
func (selector *RouletteWheelRankWeightingSelector) SelectInd() int {
	rnd := rand.Float64()

	sum := 0.0
	for i := 0; i < len(selector.population); i++ {
		sum += selector.weights[i]

		if rnd < sum {
			log.Tracef("Found chrom on %d", i)
			return i
		}
	}

	panic("Select can't select")
}
