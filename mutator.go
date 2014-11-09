package genetic_algorithm

type MutatorInterface interface {
	Mutate([]ChromosomeInterface)
}