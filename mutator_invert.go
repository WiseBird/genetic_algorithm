package genetic_algorithm

// Mutator selects some part of the chromosome and inverts it
type InvertMutator struct {
	*MutatorIntervalBase
}

// Probability is applied to each chromosome
func NewInvertMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor) *InvertMutator {
	mutator := new(InvertMutator)

	mutator.MutatorIntervalBase = NewMutatorIntervalBase(mutator, probability, chromosomeConstructor)

	return mutator
}

func (mutator *InvertMutator) MutateGenes(genes GenesInterface, from, to int) {
	temp := mutator.getIntervalCopy(genes, from, to)

	for i := from; i < to; i++ {
		ind := to + from - i - 1
		genes.Set(i, temp.Get(ind))
	}
}
