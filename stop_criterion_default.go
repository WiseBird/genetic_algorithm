package genetic_algorithm

type StopCriterionDefault struct {
	maxIterations int
	maxIterationsCrit bool

	minCost float64
	minCostCrit bool
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
func (criterion *StopCriterionDefault) Setup(statistics StatisticsInterface) {
	if criterion.minCostCrit {
		statistics.TrackMinCost()
	}
}

func (criterion *StopCriterionDefault) ShouldStop(statistics StatisticsInterface) bool {
	if criterion.maxIterationsCrit && statistics.Iterations() >= criterion.maxIterations {
		return true
	}
	if criterion.minCostCrit && statistics.MinCost() <= criterion.minCost {
		return true
	}

	return false
}
