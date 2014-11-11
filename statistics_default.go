package genetic_algorithm

import (
	"time"
	//log "github.com/cihub/seelog"
)

type StatisticsDefaultInterface interface {
	StatisticsInterface

	// Returns time elapsed from start to end.
	// If optimization are in progress then returns time elapsed from start.
	Duration() time.Duration

	Iterations() int

	// Min cost for last iteration
	MinCost() float64
	// Min cost for each iteration
	// Len would be `Iterations() + 1` because of initial value
	MinCosts() []float64
	// Number of iterations during which the value remains unchanged 
	MinCostAge() int

	// Mean cost of last iteration
	MeanCost() float64
	// Mean cost of each iteration
	// Len would be `Iterations() + 1` because of initial value
	MeanCosts() []float64

	// Worst cost for last iteration
	WorstCost() float64
	// Worst cost for each iteration
	// Len would be `Iterations() + 1` because of initial value
	WorstCosts() []float64

	TrackMinCosts()
	TrackMinCostAge()
	TrackMeanCost()
	TrackMeanCosts()
	TrackWorstCost()
	TrackWorstCosts()
}

// Default realization of StatisticsInterface
type StatisticsDefault struct {
	started bool
	startTime time.Time
	elapsed time.Duration

	iterations int

	minCost float64
	minCostAge int
	prevDifferentMinCost float64
	minCosts []float64

	meanCost float64
	meanCosts []float64

	worstCost float64
	worstCosts []float64

	trackMinCosts bool
	trackMinCostAge bool
	trackMeanCost bool
	trackMeanCosts bool
	trackWorstCost bool
	trackWorstCosts bool
}
func NewStatisticsDefault() StatisticsInterface {
	statistics := new(StatisticsDefault)

	statistics.iterations = -1

	return statistics
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
func (statistics *StatisticsDefault) Duration() time.Duration {
	if !statistics.started {
		return statistics.elapsed
	}

	return time.Since(statistics.startTime)
}

// Expects sorted population
func (statistics *StatisticsDefault) OnInteration(population Chromosomes) {
	if !statistics.started {
		panic("Statistics should be start first")
	}

	statistics.iterations++

	if len(population) == 0 {
		return
	}

	statistics.minCost = population[0].Cost()

	if statistics.trackMinCosts {
		statistics.minCosts = append(statistics.minCosts, population[0].Cost())
	}
	if statistics.trackMinCostAge {
		if statistics.iterations == 0 || statistics.prevDifferentMinCost != statistics.minCost {
			statistics.prevDifferentMinCost = statistics.minCost
			statistics.minCostAge = 0
		} else {
			statistics.minCostAge++
		}
	}

	if statistics.trackMeanCost {
		statistics.meanCost = population.MeanCost()
	}
	if statistics.trackMinCosts {
		var mean float64
		if statistics.trackMinCosts {
			mean = statistics.meanCost
		} else {
			mean = population.MeanCost()
		}

		statistics.meanCosts = append(statistics.meanCosts, mean)
	}


	if statistics.trackWorstCost {
		statistics.worstCost = population[len(population) - 1].Cost()
	}
	if statistics.trackWorstCosts {
		statistics.worstCosts = append(statistics.worstCosts, population[len(population) - 1].Cost())
	}
}

func (statistics *StatisticsDefault) Iterations() int {
	return statistics.iterations
}
func (statistics *StatisticsDefault) MinCost() float64 {
	return statistics.minCost
}
func (statistics *StatisticsDefault) MinCosts() []float64 {
	return statistics.minCosts
}
func (statistics *StatisticsDefault) MinCostAge() int {
	return statistics.minCostAge
}
func (statistics *StatisticsDefault) MeanCost() float64 {
	return statistics.meanCost
}
func (statistics *StatisticsDefault) MeanCosts() []float64 {
	return statistics.meanCosts
}
func (statistics *StatisticsDefault) WorstCost() float64 {
	return statistics.worstCost
}
func (statistics *StatisticsDefault) WorstCosts() []float64 {
	return statistics.worstCosts
}

func (statistics *StatisticsDefault) TrackMinCosts() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackMinCosts = true
}
func (statistics *StatisticsDefault) TrackMinCostAge() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackMinCostAge = true
}
func (statistics *StatisticsDefault) TrackMeanCost() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackMeanCost = true
}
func (statistics *StatisticsDefault) TrackMeanCosts() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackMeanCosts = true
}
func (statistics *StatisticsDefault) TrackWorstCost() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackWorstCost = true
}
func (statistics *StatisticsDefault) TrackWorstCosts() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackWorstCosts = true
}