package genetic_algorithm

// Mutator for binary chromosomes
// Simply inverts specified bit
type BinaryMutator struct {

}
func NewBinaryMutator(probability float64) *MutatorBase {
	mutator := NewMutator(new(BinaryMutator), probability)

	return mutator
}
func(mutator *BinaryMutator) MutateCromosome(chrom ChromosomeInterface, ind int) {
	bc, ok := chrom.(*BinaryChromosome)
	if !ok {
		panic("Expects BinaryChromosome")
	}

	bc.genes[ind] = !bc.genes[ind]
}