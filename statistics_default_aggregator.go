package genetic_algorithm

import (
	"time"
)

type StatisticsDefaultAggregator struct {
	options *StatisticsDefaultOptions
	statistics []*StatisticsDefault

	duration time.Duration

	generations int

	minCost float64
	minCosts []float64
	minCostsVar float64

	meanCost float64
	meanCosts []float64

	worstCost float64
	worstCosts []float64
}
func NewStatisticsDefaultAggregator(options *StatisticsDefaultOptions) *StatisticsDefaultAggregator {
	aggregator := new(StatisticsDefaultAggregator)

	aggregator.options = options

	return aggregator
}
func (aggregator *StatisticsDefaultAggregator) Aggregate(statistics StatisticsInterface) {
	stats, ok := statistics.(*StatisticsDefault)
	if !ok {
		panic("Expects instance of StatisticsDefault")
	}

	aggregator.statistics = append(aggregator.statistics, stats)
}
func (aggregator *StatisticsDefaultAggregator) Compute() {
	count := len(aggregator.statistics)

	aggregator.duration = time.Duration(
		meanInt64Iter(count, func(i int) int64 {
			return int64(aggregator.statistics[i].elapsed)
		}))
	aggregator.generations = int(
		meanInt64Iter(count, func(i int) int64 {
			return int64(aggregator.statistics[i].generations)
		}))

	aggregator.minCost = 
		meanFloat64Iter(count, func(i int) float64 {
			return aggregator.statistics[i].minCost
		})

	if aggregator.options.trackMinCosts {
		aggregator.minCosts =
			meanFloat64ArrIter(count, func(i int) []float64 {
				return aggregator.statistics[i].minCosts
			})
	}
	if aggregator.options.trackMinCostsVar {
		aggregator.minCostsVar = 
			meanFloat64Iter(count, func(i int) float64 {
				return aggregator.statistics[i].minCostsVar
			})
	}

	if aggregator.options.trackMeanCost {
		aggregator.meanCost = 
			meanFloat64Iter(count, func(i int) float64 {
				return aggregator.statistics[i].meanCost
			})
	}
	if aggregator.options.trackMeanCosts {
		aggregator.meanCosts =
			meanFloat64ArrIter(count, func(i int) []float64 {
				return aggregator.statistics[i].meanCosts
			})
	}

	if aggregator.options.trackWorstCost {
		aggregator.worstCost = 
			meanFloat64Iter(count, func(i int) float64 {
				return aggregator.statistics[i].worstCost
			})
	}
	if aggregator.options.trackWorstCosts {
		aggregator.worstCosts =
			meanFloat64ArrIter(count, func(i int) []float64 {
				return aggregator.statistics[i].worstCosts
			})
	}
}

func (aggregator *StatisticsDefaultAggregator) Generations() int {
	return aggregator.generations
}
func (aggregator *StatisticsDefaultAggregator) Duration() time.Duration {
	return aggregator.duration
}
func (aggregator *StatisticsDefaultAggregator) MinCost() float64 {
	return aggregator.minCost
}
func (aggregator *StatisticsDefaultAggregator) MinCosts() []float64 {
	return aggregator.minCosts
}
func (aggregator *StatisticsDefaultAggregator) MinCostsVar() float64 {
	return aggregator.minCostsVar
}
func (aggregator *StatisticsDefaultAggregator) MeanCost() float64 {
	return aggregator.meanCost
}
func (aggregator *StatisticsDefaultAggregator) MeanCosts() []float64 {
	return aggregator.meanCosts
}
func (aggregator *StatisticsDefaultAggregator) WorstCost() float64 {
	return aggregator.worstCost
}
func (aggregator *StatisticsDefaultAggregator) WorstCosts() []float64 {
	return aggregator.worstCosts
}