package genetic_algorithm

import (
	"math/rand"
)

// InvertMutator + SwapMutator
// Mutator selects some part of the chromosome inverts it then swap on gene from it with one gene outside of the interval
type InvertSwapMutator struct {
	*MutatorIntervalBase
	inverter *InvertMutator
}

// Probability is applied to each chromosome
func NewInvertSwapMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor) *InvertSwapMutator {
	mutator := new(InvertSwapMutator)

	mutator.MutatorIntervalBase = NewMutatorIntervalBase(mutator, probability, chromosomeConstructor)
	mutator.inverter = NewInvertMutator(probability, chromosomeConstructor)

	return mutator
}

func (mutator *InvertSwapMutator) MutateGenes(genes GenesInterface, from, to int) {
	mutator.inverter.MutateGenes(genes, from, to)
	mutator.swap(genes, from, to)
}
func (mutator *InvertSwapMutator) swap(genes GenesInterface, from, to int) {
	ind1 := mutator.chooseFirstInd(genes.Len(), from , to)
	ind2 := mutator.chooseSecondInd(genes.Len(), from , to)
	genes.Swap(ind1, ind2)
}
func (mutator *InvertSwapMutator) chooseFirstInd(genesLen, from, to int) int {
	return rand.Intn(to - from) + from
}
func (mutator *InvertSwapMutator) chooseSecondInd(genesLen, from, to int) int {
	ind2 := rand.Intn(genesLen - from)
	if ind2 >= from {
		ind2 += from
	}

	return ind2
}