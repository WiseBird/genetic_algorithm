package genetic_algorithm

import (
	"math/rand"
)

// Mutator for ordered chromosomes
// Simply swaps two elements
type OrderedSwapMutator struct {

}
func NewOrderedSwapMutator(probability float64) *MutatorBase {
	mutator := NewMutator(new(OrderedSwapMutator), probability)

	return mutator
}
func(mutator *OrderedSwapMutator) MutateCromosome(chrom ChromosomeInterface, ind int) {
	bc, ok := chrom.(*OrderedChromosome)
	if !ok {
		panic("Expects OrderedChromosome")
	}

	for ;; {
		ind2 := rand.Intn(chrom.Genes().Len())
		if ind == ind2 {
			continue
		}

		bc.genes[ind], bc.genes[ind2] = bc.genes[ind2], bc.genes[ind]
		break
	}
}