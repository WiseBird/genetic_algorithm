package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math"
	"time"
)

type StatisticsDataDefault interface {
	Generations() int
	Duration() time.Duration
	MinCost() float64
	MinCosts() []float64
	MinCostsVar() float64
	MeanCost() float64
	MeanCosts() []float64
	WorstCost() float64
	WorstCosts() []float64
}

// Default realization of StatisticsInterface
type StatisticsDefault struct {
	started   bool
	startTime time.Time
	elapsed   time.Duration

	generations int

	minCost              float64
	gensWoImprv          int
	prevDifferentMinCost float64
	minCosts             []float64
	minCostsVar          float64

	meanCost  float64
	meanCosts []float64

	worstCost  float64
	worstCosts []float64

	options *StatisticsDefaultOptions
}

func NewStatisticsDefault(options StatisticsOptionsInterface) StatisticsInterface {
	opts, ok := options.(*StatisticsDefaultOptions)
	if !ok {
		panic("Expects instance of StatisticsOptionsInterface")
	}

	statistics := new(StatisticsDefault)

	statistics.generations = -1
	statistics.options = opts

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

// Expects sorted population
func (statistics *StatisticsDefault) OnGeneration(population Chromosomes) {
	if !statistics.started {
		panic("Statistics should be start first")
	}

	statistics.generations++

	if len(population) == 0 {
		panic("Population is empty")
	}

	statistics.minCost = population[0].Cost()
	log.Tracef("MinCost %v", statistics.minCost)

	if statistics.options.trackMinCosts {
		statistics.minCosts = append(statistics.minCosts, population[0].Cost())
		log.Tracef("MinCosts %v", statistics.minCosts)
	}
	if statistics.options.trackGensWoImprv {
		if statistics.generations == 0 || statistics.prevDifferentMinCost != statistics.minCost {
			statistics.prevDifferentMinCost = statistics.minCost
			statistics.gensWoImprv = 0
		} else {
			statistics.gensWoImprv++
		}
		log.Tracef("MinCostAge %v", statistics.gensWoImprv)
	}
	if statistics.options.trackMinCostsVar {
		if statistics.generations == 0 || statistics.generations == statistics.gensWoImprv {
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
		statistics.worstCost = population[len(population)-1].Cost()
		log.Tracef("WorstCost %v", statistics.worstCost)
	}
	if statistics.options.trackWorstCosts {
		statistics.worstCosts = append(statistics.worstCosts, population[len(population)-1].Cost())
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

// Number of generations during which the min cost remains unchanged
func (statistics *StatisticsDefault) GenerationsWithoutImprovements() int {
	return statistics.gensWoImprv
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

func (statistics *StatisticsDefault) Data() StatisticsDataInterface {
	return statistics
}
