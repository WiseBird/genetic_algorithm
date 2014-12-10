package genetic_algorithm

import (
	"time"
)

type StatisticsDefaultAggregator struct {
	options    *StatisticsDefaultOptions
	statistics []*StatisticsDefault

	duration time.Duration

	generations int

	minCost     float64
	gensWoImprv          int
	minCosts    []float64
	minCostsVar float64

	meanCost  float64
	meanCosts []float64

	worstCost  float64
	worstCosts []float64
}

func NewStatisticsDefaultAggregator(options StatisticsOptionsInterface) StatisticsAggregatorInterface {
	opts, ok := options.(*StatisticsDefaultOptions)
	if !ok {
		panic("Expects instance of StatisticsOptionsInterface")
	}

	aggregator := new(StatisticsDefaultAggregator)

	aggregator.options = opts

	return aggregator
}
func (aggregator *StatisticsDefaultAggregator) Options() StatisticsOptionsInterface {
	return aggregator.options.Copy()
}
func (aggregator *StatisticsDefaultAggregator) Aggregate(statistics StatisticsDataInterface) {
	stats, ok := statistics.(*StatisticsDefault)
	if !ok {
		panic("Expects instance of StatisticsDefault")
	}

	aggregator.statistics = append(aggregator.statistics, stats)
}
func (aggregator *StatisticsDefaultAggregator) Compute() StatisticsDataInterface {
	count := len(aggregator.statistics)

	aggregator.duration = time.Duration(
		meanInt64Iter(count, func(i int) int64 {
			return int64(aggregator.statistics[i].Duration())
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
	if aggregator.options.trackGensWoImprv {
		aggregator.gensWoImprv = int(
		meanInt64Iter(count, func(i int) int64 {
			return int64(aggregator.statistics[i].gensWoImprv)
		}))
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

	return aggregator
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
func (aggregator *StatisticsDefaultAggregator) GenerationsWithoutImprovements() int {
	return aggregator.gensWoImprv
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

func (aggregator *StatisticsDefaultAggregator) Data() StatisticsDataInterface {
	return aggregator
}
