package genetic_algorithm

type StatisticsConstructor func() StatisticsInterface

// Accumulate statistics during optimization.
type StatisticsInterface interface {
	Start()
	End()
	// Optimizer will call this method on each generation
	// First call would be with initial population
	OnGeneration(Chromosomes)
}