package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math/rand"
)

// Selects individual with probability proportional to it fitness value.
// Warning! In order to use this selector cost value must be normalized, i.e. chromosome with cost=0 is the best solution.
type RouletteWheelCostWeightingSelector struct {
	*SelectorBase

	fitnessSum float64
}

func NewRouletteWheelCostWeightingSelector() *RouletteWheelCostWeightingSelector {
	selector := new(RouletteWheelCostWeightingSelector)

	selector.SelectorBase = NewSelectorBase(selector)

	return selector
}
func (selector *RouletteWheelCostWeightingSelector) Prepare(population Chromosomes) {
	log.Tracef("Preparing")

	selector.SelectorBase.Prepare(population)

	fitnessSum := 0.0

	for i := 0; i < len(selector.population); i++ {
		chrom := selector.population[i]
		fitnessSum += selector.fitness(chrom.Cost())
	}

	selector.fitnessSum = fitnessSum

	log.Tracef("Prepared fs=%f\n", selector.fitnessSum)
}
func (selector *RouletteWheelCostWeightingSelector) fitness(cost float64) float64 {
	if cost < 0 {
		panic("Can't calc fitness for negative cost")
	}

	return 1 / (cost + 1)
}
func (selector *RouletteWheelCostWeightingSelector) SelectInd() int {
	rnd := rand.Float64() * selector.fitnessSum

	sum := 0.0
	for i := 0; i < len(selector.population); i++ {
		chrom := selector.population[i]
		sum += selector.fitness(chrom.Cost())

		if rnd < sum {
			log.Tracef("Found chrom %v, on %d", chrom, i)
			return i
		}
	}

	panic("Select can't select")
}
