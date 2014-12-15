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
	ind2 := mutator.chooseSecondInd(chrom.Genes().Len(), ind)
	chrom.Genes().Swap(ind, ind2)
}
func (mutator *SwapMutator) chooseSecondInd(genesLen, ind int) int {
	ind2 := rand.Intn(genesLen - 1)
	if ind2 >= ind {
		ind2++
	}
	return ind2
}