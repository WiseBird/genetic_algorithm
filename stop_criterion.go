package genetic_algorithm

// 
type StopCriterionInterface interface {
	// Method sets up statistics tracking
	// Executes one time before optimization
	Setup(StatisticsInterface)

	// Method executes each time before iteration
	ShouldStop(StatisticsInterface) bool
}