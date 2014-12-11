package genetic_algorithm

// Mutator for binary chromosomes
// Simply inverts specified bit
type BinaryMutator struct {
}

// Probability is applied to each element separately.
func NewBinaryMutator(probability float64) *MutatorGeneBase {
	mutator := NewGeneBaseMutator(new(BinaryMutator), probability)

	return mutator
}
func (mutator *BinaryMutator) MutateCromosome(chrom ChromosomeInterface, ind int) {
	bc, ok := chrom.(*BinaryChromosome)
	if !ok {
		panic("Expects BinaryChromosome")
	}

	bc.genes[ind] = !bc.genes[ind]
}
