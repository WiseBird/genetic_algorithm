package genetic_algorithm

// Interface for selectors -
type SelectorInterface interface {
	// Called once before selection for each population
	Prepare(Chromosomes)

	// Selects one parent from population
	Select() ChromosomeInterface
	// Selects c parents from population
	SelectMany(c int) Chromosomes
}
