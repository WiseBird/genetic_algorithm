package genetic_algorithm

import (
	"time"
	"math"
	log "github.com/cihub/seelog"
)

// Default realization of StatisticsInterface
type StatisticsDefault struct {
	started bool
	startTime time.Time
	elapsed time.Duration

	generations int

	minCost float64
	minCostAge int
	prevDifferentMinCost float64
	minCosts []float64
	minCostsVar float64

	meanCost float64
	meanCosts []float64

	worstCost float64
	worstCosts []float64

	options *StatisticsDefaultOptions
}
func NewStatisticsDefault() StatisticsInterface {
	statistics := new(StatisticsDefault)

	statistics.generations = -1

	return statistics
}
func (statistics *StatisticsDefault) SetOptions(options StatisticsOptionsInterface) {
	statistics.options = options.(*StatisticsDefaultOptions)
}
func (statistics *StatisticsDefault) Start() {
	statistics.started = true
	statistics.startTime = time.Now()
}
func (statistics *StatisticsDefault) End() {
	if !statistics.started {
		return
	}

	statistics.started = false
	statistics.elapsed = time.Since(statistics.startTime)
}

// Expects sorted population
func (statistics *StatisticsDefault) OnGeneration(population Chromosomes) {
	if !statistics.started {
		panic("Statistics should be start first")
	}

	statistics.generations++

	if len(population) == 0 {
		return
	}

	statistics.minCost = population[0].Cost()
	log.Tracef("MinCost %v", statistics.minCost)

	if statistics.options.trackMinCosts {
		statistics.minCosts = append(statistics.minCosts, population[0].Cost())
		log.Tracef("MinCosts %v", statistics.minCosts)
	}
	if statistics.options.trackMinCostAge {
		if statistics.generations == 0 || statistics.prevDifferentMinCost != statistics.minCost {
			statistics.prevDifferentMinCost = statistics.minCost
			statistics.minCostAge = 0
		} else {
			statistics.minCostAge++
		}
		log.Tracef("MinCostAge %v", statistics.minCostAge)
	}
	if statistics.options.trackMinCostsVar {
		if statistics.generations == 0 || statistics.generations == statistics.minCostAge {
			statistics.minCostsVar = math.NaN()
		} else {
			statistics.minCostsVar = pvarianceFloat64(statistics.minCosts)
		}
		log.Tracef("MinCostsVar %v", statistics.minCostsVar)
	}

	if statistics.options.trackMeanCost {
		statistics.meanCost = population.MeanCost()
		log.Tracef("MeanCost %v", statistics.meanCost)
	}
	if statistics.options.trackMeanCosts {
		var mean float64
		if statistics.options.trackMeanCost {
			mean = statistics.meanCost
		} else {
			mean = population.MeanCost()
		}

		statistics.meanCosts = append(statistics.meanCosts, mean)
		log.Tracef("MeanCosts %v", statistics.meanCosts)
	}


	if statistics.options.trackWorstCost {
		statistics.worstCost = population[len(population) - 1].Cost()
		log.Tracef("WorstCost %v", statistics.worstCost)
	}
	if statistics.options.trackWorstCosts {
		statistics.worstCosts = append(statistics.worstCosts, population[len(population) - 1].Cost())
		log.Tracef("WorstCosts %v", statistics.worstCosts)
	}
}

// Number of generations
func (statistics *StatisticsDefault) Generations() int {
	return statistics.generations
}
// Returns time elapsed from start to end.
// If optimization are in progress then returns time elapsed from start.
func (statistics *StatisticsDefault) Duration() time.Duration {
	if !statistics.started {
		return statistics.elapsed
	}

	return time.Since(statistics.startTime)
}
// Min cost for last iteration
func (statistics *StatisticsDefault) MinCost() float64 {
	return statistics.minCost
}
// Min cost for each iteration
// Len would be `Generations() + 1` because of initial value
func (statistics *StatisticsDefault) MinCosts() []float64 {
	return statistics.minCosts
}
// Number of generations during which the value remains unchanged 
func (statistics *StatisticsDefault) MinCostAge() int {
	return statistics.minCostAge
}
// Variance of min costs
// Variance equals NaN until two different values of MinCost are obtained
func (statistics *StatisticsDefault) MinCostsVar() float64 {
	return statistics.minCostsVar
}
// Mean cost of last iteration
func (statistics *StatisticsDefault) MeanCost() float64 {
	return statistics.meanCost
}
// Mean cost of each iteration
// Len would be `Iterations() + 1` because of initial value
func (statistics *StatisticsDefault) MeanCosts() []float64 {
	return statistics.meanCosts
}
// Worst cost for last iteration
func (statistics *StatisticsDefault) WorstCost() float64 {
	return statistics.worstCost
}
// Worst cost for each iteration
// Len would be `Iterations() + 1` because of initial value
func (statistics *StatisticsDefault) WorstCosts() []float64 {
	return statistics.worstCosts
}