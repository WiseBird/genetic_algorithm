package genetic_algorithm

type StopCriterionInterface interface {
	ShouldStop(StatisticsInterface) bool
	Setup(StatisticsInterface)
}