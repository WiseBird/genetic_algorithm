package genetic_algorithm

import (
	log "github.com/cihub/seelog"
)

type StopCriterionDefault struct {
	maxGenerations int
	maxGenerationsCrit bool

	minCost float64
	minCostCrit bool

	maxMinCostAge int
	maxMinCostAgeCrit bool

	minMinCostsVar float64
	minMinCostsVarCrit bool
}
func NewStopCriterionDefault() *StopCriterionDefault {
	criterion := new(StopCriterionDefault)

	return criterion
}
// Stop when specified number of generations is reached
func (criterion *StopCriterionDefault) MaxGenerations(value int) *StopCriterionDefault {
	if value < 0 {
		panic("Value can't be negative")
	}

	criterion.maxGenerationsCrit = true
	criterion.maxGenerations = value
	return criterion
}
// Stop when min cost less then or equals value
func (criterion *StopCriterionDefault) MinCost(value float64) *StopCriterionDefault {
	criterion.minCostCrit = true
	criterion.minCost = value
	return criterion
}
// Stop when min cost wasn't changed for value generations
func (criterion *StopCriterionDefault) MaxMinCostAge(value int) *StopCriterionDefault {
	criterion.maxMinCostAgeCrit = true
	criterion.maxMinCostAge = value
	return criterion
}
// Stop when variance of min cost less then or equals value
func (criterion *StopCriterionDefault) MinMinCostsVar(value float64) *StopCriterionDefault {
	criterion.minMinCostsVarCrit = true
	criterion.minMinCostsVar = value
	return criterion
}

func (criterion *StopCriterionDefault) Setup(statistics StatisticsInterface) {
	stats, ok := statistics.(*StatisticsDefault)
	if !ok {
		panic("Method expects StatisticsDefault")
	}

	if criterion.maxMinCostAgeCrit {
		stats.TrackMinCostAge()
	}
	if criterion.minMinCostsVarCrit {
		stats.TrackMinCostsVar()
	}
}

func (criterion *StopCriterionDefault) ShouldStop(statistics StatisticsInterface) bool {
	stats, ok := statistics.(*StatisticsDefault)
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
	if criterion.maxMinCostAgeCrit {
		log.Debugf("MinCostAge %v", stats.MinCostAge())
		if stats.MinCostAge() >= criterion.maxMinCostAge {
			log.Info("Stop by max min cost's age")
			return true
		}
	}
	if criterion.minMinCostsVarCrit {
		log.Debugf("MinCostsVar %v", stats.MinCostsVar())
		if stats.MinCostsVar() >= criterion.minMinCostsVar {
			log.Info("Stop by max min cost's age")
			return true
		}
	}

	return false
}
