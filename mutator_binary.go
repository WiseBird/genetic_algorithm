package genetic_algorithm

// Mutator for binary chromosomes
// Simply inverts specified bit
type BinaryMutator struct {

}
func NewBinaryMutator(probability float64) MutatorInterface {
	mutator := NewMutator(new(BinaryMutator), probability)

	return mutator
}
func(mutator *BinaryMutator) MutateCromosome(chrom ChromosomeInterface, ind int) {
	bc := chrom.(*BinaryChromosome)

	bc.Genes[ind] = !bc.Genes[ind]
}