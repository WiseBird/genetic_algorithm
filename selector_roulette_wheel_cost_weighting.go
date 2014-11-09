package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math/rand"
)

// Selects individual with probability proportional to it fitness value.
// Warning! In order to use this selector cost value must be normalized, i.e. chromosome with cost=0 is the best solution.
type RouletteWheelCostWeightingSelector struct {
	Population Chromosomes
	FitnessSum float64
}

func NewRouletteWheelCostWeightingSelector() *SelectorBase {
	selector := new(RouletteWheelCostWeightingSelector)

	return NewSelectorBase(selector)
}
func (selector *RouletteWheelCostWeightingSelector) Prepare(population Chromosomes) {
	log.Tracef("Preparing")

	selector.Population = population

	fitnessSum := 0.0

	for i := 0; i < len(selector.Population); i++ {
		chrom := selector.Population[i]
		fitnessSum += chrom.Fitness()
	}

	selector.FitnessSum = fitnessSum

	log.Tracef("Prepared fs=%f\n", selector.FitnessSum)
}
func (selector *RouletteWheelCostWeightingSelector) Select() ChromosomeInterface {
	rnd := rand.Float64() * selector.FitnessSum

	sum := 0.0
	for i := 0; i < len(selector.Population); i++ {
		chrom := selector.Population[i]
		sum += chrom.Fitness()

		if rnd < sum {
			log.Tracef("Found chrom %v, on %d", chrom, i)
			return chrom
		}
	}

	panic("Select can't select")
}