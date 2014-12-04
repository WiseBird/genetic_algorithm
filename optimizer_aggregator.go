package genetic_algorithm

import ()

type OptimizerAggregator struct {
	optimizer                       OptimizerInterface
	statisticsAggregatorConstructor StatisticsAggregatorConstructor
	statisticsOptions               StatisticsOptionsInterface
	iterations                      int
}

func NewOptimizerAggregator() *OptimizerAggregator {
	aggregator := new(OptimizerAggregator)

	aggregator.statisticsAggregatorConstructor = NewStatisticsDefaultAggregator
	aggregator.statisticsOptions = NewStatisticsDefaultOptions()

	return aggregator
}

func (aggregator *OptimizerAggregator) Optimizer(optimizer OptimizerInterface) *OptimizerAggregator {
	aggregator.optimizer = optimizer
	return aggregator
}
func (aggregator *OptimizerAggregator) StatisticsAggregatorConstructor(constr StatisticsAggregatorConstructor) *OptimizerAggregator {
	aggregator.statisticsAggregatorConstructor = constr
	return aggregator
}
func (aggregator *OptimizerAggregator) StatisticsOptions(statisticsOptions StatisticsOptionsInterface) *OptimizerAggregator {
	aggregator.statisticsOptions = statisticsOptions
	return aggregator
}
func (aggregator *OptimizerAggregator) Iterations(iterations int) *OptimizerAggregator {
	aggregator.iterations = iterations
	return aggregator
}

func (aggregator *OptimizerAggregator) check() {
	if aggregator.optimizer == nil {
		panic("Optimizer must be set")
	}
	if aggregator.statisticsAggregatorConstructor == nil {
		panic("StatisticsAggregatorConstructor must be set")
	}
	if aggregator.statisticsOptions == nil {
		panic("StatisticsOptions must be set")
	}
	if aggregator.iterations <= 0 {
		panic("Iterations must be positive value")
	}
}

func (aggregator *OptimizerAggregator) Optimize() (ChromosomeInterface, StatisticsDataInterface) {
	aggregator.check()

	statisticsAggregator := aggregator.statisticsAggregatorConstructor(aggregator.statisticsOptions)
	aggregator.ensureOptimizerOptions(statisticsAggregator)

	var bestChrom ChromosomeInterface
	for i := 0; i < aggregator.iterations; i++ {
		chrom, stats := aggregator.optimizer.Optimize()
		if bestChrom == nil || bestChrom.Cost() > chrom.Cost() {
			bestChrom = chrom
		}

		statisticsAggregator.Aggregate(stats)
	}

	return bestChrom, statisticsAggregator.Compute()
}
func (aggregator *OptimizerAggregator) ensureOptimizerOptions(statisticsAggregator StatisticsAggregatorInterface) {
	optimizerWithStatisticsOptionsSetup, ok := aggregator.optimizer.(OptimizerWithStatisticsOptionsSetup)
	if !ok {
		return
	}

	statisticsAggregatorWithOptions, ok := statisticsAggregator.(StatisticsAggregatorWithOptions)
	if !ok {
		return
	}

	statisticsAggregatorWithOptions.Options().Ensure(
		optimizerWithStatisticsOptionsSetup.SetupStatisticsOptions())
}
