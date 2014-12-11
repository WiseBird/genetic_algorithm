package genetic_algorithm

import (
	"math/rand"
)

// Mutator simply swaps two elements
type SwapMutator struct {
}

// Probability is applied to each element separately.
func NewSwapMutator(probability float64) *MutatorGeneBase {
	mutator := NewGeneBaseMutator(new(SwapMutator), probability)

	return mutator
}
func (mutator *SwapMutator) MutateCromosome(chrom ChromosomeInterface, ind int) {
	for {
		ind2 := rand.Intn(chrom.Genes().Len())
		if ind == ind2 {
			continue
		}

		chrom.Genes().Swap(ind, ind2)
		break
	}
}
