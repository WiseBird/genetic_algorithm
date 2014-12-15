package genetic_algorithm

// InvertMutator + DisplacementMutator
// Mutator selects some part of the chromosome inverts it then place at other position
type InvertDisplacementMutator struct {
	*MutatorIntervalBase
	inverter *InvertMutator
	displacer *DisplacementMutator
}

// Probability is applied to each chromosome
func NewInvertDisplacementMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor) *InvertDisplacementMutator {
	mutator := new(InvertDisplacementMutator)

	mutator.MutatorIntervalBase = NewMutatorIntervalBase(mutator, probability, chromosomeConstructor)
	mutator.inverter = NewInvertMutator(probability, chromosomeConstructor)
	mutator.displacer = NewDisplacementMutator(probability, chromosomeConstructor)

	return mutator
}

func (mutator *InvertDisplacementMutator) MutateGenes(genes GenesInterface, from, to int) {
	mutator.inverter.MutateGenes(genes, from, to)
	mutator.displacer.MutateGenes(genes, from, to)
}