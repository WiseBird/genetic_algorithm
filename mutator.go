package genetic_algorithm

// Mutator interface
type MutatorInterface interface {
	Mutate(Chromosomes)
}

var (
	NopMutator = &nopMutator{}
)

type nopMutator struct{}

func (mutator *nopMutator) Mutate(c Chromosomes) {}
