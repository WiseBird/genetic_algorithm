package genetic_algorithm

type StatisticsConstructor func(StatisticsOptionsInterface) StatisticsInterface

// Accumulate statistics during optimization.
type StatisticsInterface interface {
	// Optimizer will call this method before optimization
	Start()
	// Optimizer will call this method after optimization
	End()
	// Optimizer will call this method on each generation
	// First call would be with initial population
	OnGeneration(Chromosomes)
}

// Options for statistics
// Defines what info gather and what not
type StatisticsOptionsInterface interface{}

type StatisticsAggregatorInterface interface {
	Aggregate(StatisticsInterface)
	Compute()
}
