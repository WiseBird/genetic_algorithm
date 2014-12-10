package genetic_algorithm

type StatisticsConstructor func(StatisticsOptionsInterface) StatisticsInterface
type StatisticsAggregatorConstructor func(StatisticsOptionsInterface) StatisticsAggregatorInterface

// Accumulate statistics during optimization
type StatisticsInterface interface {
	// Method for duration tracking
	// Call without parameters means that optimization is started
	Start(keys ...string)
	// Should be called when tracked operation is complete.
	End()
	// Optimizer will call this method on each generation
	// First call would be with initial population
	OnGeneration(Chromosomes)
	Data() StatisticsDataInterface
}

type StatisticsDataInterface interface{}

// Options for statistics
// Defines what data gather and what not
type StatisticsOptionsInterface interface {
	// Ensures that other tracks what ensurer tracks
	Ensure(other StatisticsOptionsInterface)
}

// Statistics aggregator
type StatisticsAggregatorInterface interface {
	Aggregate(StatisticsDataInterface)
	// Calcs mean of aggregated data
	// Should be called on the end of aggregation
	Compute() StatisticsDataInterface
}
type StatisticsAggregatorWithOptions interface {
	StatisticsAggregatorInterface
	Options() StatisticsOptionsInterface
}
