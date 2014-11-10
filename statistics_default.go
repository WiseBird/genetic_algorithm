package genetic_algorithm

import (
	"time"
	//log "github.com/cihub/seelog"
)

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

	trackMinCost bool
	trackMinCosts bool
	trackMinAge bool
	trackMeanCost bool
	trackMeanCosts bool
	trackWorstCost bool
	trackWorstCosts bool
}
func NewStatisticsDefault() *StatisticsDefault {
	statistics := new(StatisticsDefault)

	statistics.iterations = -1

	return statistics
}
func StatisticsDefaultConstructor() StatisticsInterface {
	return NewStatisticsDefault()
}
func (statistics *StatisticsDefault) Start() {
	statistics.started = true
	statistics.startTime = time.Now()
}
func (statistics *StatisticsDefault) End() {
	statistics.elapsed = time.Since(statistics.startTime)
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

	if statistics.trackMinCost {
		statistics.minCost = population[0].Cost()
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

func (statistics *StatisticsDefault) TrackMinCost() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackMinCost = true
}
func (statistics *StatisticsDefault) TrackMinCosts() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackMinCosts = true
}
func (statistics *StatisticsDefault) TrackMinAge() {
	if statistics.started {
		panic("Statistics should be set up before optimization")
	}
	statistics.trackMinAge = true
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