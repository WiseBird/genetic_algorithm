/*
Example of solving tsp.
*/

package main

import (
	. "github.com/WiseBird/genetic_algorithm"
	log "github.com/cihub/seelog"
	"github.com/WiseBird/genetic_algorithm/plotting"
	tsp "github.com/WiseBird/genetic_algorithm/examples/tsp"
)

func main() {
	defer log.Flush()
	setupLogger()

	iterations := 100

	plotting.NewPlotter().
		AddPlot("TSP").
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer(1)).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().Name("OX1.1").Done().
				Done().
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer(2)).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().Name("OX1.2").Done().
				Done().
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer(3)).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().Name("PMX").Done().
				Done().
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer(4)).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().Name("POS").Done().
				Done().
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer(5)).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().Name("ROX").Done().
				Done().
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer(6)).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().Name("PPX").Done().
				Done().
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer(7)).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().Name("CX").Done().
				Done().
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer(8)).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().Name("OX2").Done().
				Done().
			Done().
		Draw(8, 4, "plot.png")
}

func createOptimizer(ver int) OptimizerInterface {
	popSize := 32
	chromSize := len(tsp.Cities)
	weedRate := 50.0
	mutationProb := 0.05
	generations := 100

	optimizer := NewIncrementalOptimizer().
		Weeder(NewSimpleWeeder(weedRate)).
		Initializer(NewOrderedRandomInitializer()).
		Selector(NewRouletteWheelRankWeightingSelector()).
		Mutator(NewOrderedSwapMutator(mutationProb)).
		CostFunction(tsp.Cost).
		StopCriterion(NewStopCriterionDefault().
			Max_Generations(generations).
			Max_GenerationsWithoutImprovements(15)).
		PopSize(popSize).
		ChromSize(chromSize)

	if ver == 1 {
		optimizer.Crossover(NewOrderCrossoverVer1())
	} else if ver == 2 {
		optimizer.Crossover(NewOrderCrossoverVer2())
	} else if ver == 3 {
		optimizer.Crossover(NewPartiallyMappedCrossover())
	} else if ver == 4 {
		optimizer.Crossover(NewPositionBasedCrossover())
	} else if ver == 5 {
		optimizer.Crossover(NewRelativeOrderingCrossover(chromSize/2))
	} else if ver == 6 {
		optimizer.Crossover(NewPrecedencePreservativeCrossover())
	} else if ver == 7 {
		optimizer.Crossover(NewCycleCrossover())
	} else if ver == 8 {
		optimizer.Crossover(NewOrderBasedCrossover())
	}

	return optimizer
}

func setupLogger() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
	    panic(err)
	}

	log.ReplaceLogger(logger)
}