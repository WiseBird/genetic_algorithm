package genetic_algorithm

import (
	log "github.com/cihub/seelog"
)

type StopCriterionDefault struct {
	maxGenerations     int
	maxGenerationsCrit bool

	minCost     float64
	minCostCrit bool

	maxGensWoImprv     int
	maxGensWoImprvCrit bool

	minMinCostsVar     float64
	minMinCostsVarCrit bool
}

func NewStopCriterionDefault() *StopCriterionDefault {
	criterion := new(StopCriterionDefault)

	return criterion
}

// Stop when specified number of generations is reached
func (criterion *StopCriterionDefault) Max_Generations(value int) *StopCriterionDefault {
	if value < 0 {
		panic("Value can't be negative")
	}

	criterion.maxGenerationsCrit = true
	criterion.maxGenerations = value
	return criterion
}

// Stop when min cost less than or equals value
func (criterion *StopCriterionDefault) Min_Cost(value float64) *StopCriterionDefault {
	criterion.minCostCrit = true
	criterion.minCost = value
	return criterion
}

// Stop when min cost wasn't changed for value generations
func (criterion *StopCriterionDefault) Max_GenerationsWithoutImprovements(value int) *StopCriterionDefault {
	criterion.maxGensWoImprvCrit = true
	criterion.maxGensWoImprv = value
	return criterion
}

// Stop when variance of min costs less than or equals value
func (criterion *StopCriterionDefault) Min_MinCostsVar(value float64) *StopCriterionDefault {
	criterion.minMinCostsVarCrit = true
	criterion.minMinCostsVar = value
	return criterion
}

func (criterion *StopCriterionDefault) Setup(opts StatisticsOptionsInterface) {
	options, ok := opts.(*StatisticsDefaultOptions)
	if !ok {
		panic("Method expects StatisticsDefault")
	}

	if criterion.maxGensWoImprvCrit {
		options.TrackGenerationsWithoutImprovements()
	}
	if criterion.minMinCostsVarCrit {
		options.TrackMinCostsVar()
	}
}

func (criterion *StopCriterionDefault) ShouldStop(statistics StatisticsDataInterface) bool {
	stats, ok := statistics.(StatisticsDataDefault)
	if !ok {
		panic("Method expects StatisticsDefault")
	}

	if criterion.maxGenerationsCrit && stats.Generations() >= criterion.maxGenerations {
		log.Info("Stop by max generations")
		return true
	}
	if criterion.minCostCrit {
		log.Debugf("MinCost %v", stats.MinCost())
		if stats.MinCost() <= criterion.minCost {
			log.Info("Stop by min cost")
			return true
		}
	}
	if criterion.maxGensWoImprvCrit {
		log.Debugf("GenerationsWithoutImprovements %v", stats.GenerationsWithoutImprovements())
		if stats.GenerationsWithoutImprovements() >= criterion.maxGensWoImprv {
			log.Info("Stop by max min cost's age")
			return true
		}
	}
	if criterion.minMinCostsVarCrit {
		log.Debugf("MinCostsVar %v", stats.MinCostsVar())
		if stats.MinCostsVar() >= criterion.minMinCostsVar {
			log.Info("Stop by max min costs var")
			return true
		}
	}

	return false
}
