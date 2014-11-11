package genetic_algorithm

type StopCriterionDefault struct {
	maxIterations int
	maxIterationsCrit bool

	minCost float64
	minCostCrit bool

	maxMinCostAge int
	maxMinCostAgeCrit bool
}
func NewStopCriterionDefault() *StopCriterionDefault {
	criterion := new(StopCriterionDefault)

	return criterion
}
func (criterion *StopCriterionDefault) MaxIterations(value int) *StopCriterionDefault {
	if value < 0 {
		panic("Value can't be negative")
	}

	criterion.maxIterationsCrit = true
	criterion.maxIterations = value
	return criterion
}
func (criterion *StopCriterionDefault) MinCost(value float64) *StopCriterionDefault {
	criterion.minCostCrit = true
	criterion.minCost = value
	return criterion
}
func (criterion *StopCriterionDefault) MaxMinCostAge(value int) *StopCriterionDefault {
	criterion.maxMinCostAgeCrit = true
	criterion.maxMinCostAge = value
	return criterion
}

func (criterion *StopCriterionDefault) Setup(statistics StatisticsInterface) {
	stats, ok := statistics.(StatisticsDefaultInterface)
	if !ok {
		panic("Method expects StatisticsDefaultInterface")
	}

	if criterion.maxMinCostAgeCrit {
		stats.TrackMinCostAge()
	}
}

func (criterion *StopCriterionDefault) ShouldStop(statistics StatisticsInterface) bool {
	stats, ok := statistics.(StatisticsDefaultInterface)
	if !ok {
		panic("Method expects StatisticsDefaultInterface")
	}

	if criterion.maxIterationsCrit && stats.Iterations() >= criterion.maxIterations {
		return true
	}
	if criterion.minCostCrit && stats.MinCost() <= criterion.minCost {
		return true
	}
	if criterion.maxMinCostAgeCrit && stats.MinCostAge() >= criterion.maxMinCostAge {
		return true
	}

	return false
}
