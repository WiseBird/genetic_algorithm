/*
	Draws plot of min and mean costs versus generations.
	Uses logarithmic scale for costs.
*/
package main

import (
	. "github.com/WiseBird/genetic_algorithm"
	"github.com/WiseBird/genetic_algorithm/plotting"
	log "github.com/cihub/seelog"
	partition "github.com/WiseBird/genetic_algorithm/examples/partition"
)

func main() {
	defer log.Flush()
	setupLogger()

	iterations := 1000

	plotting.NewPlotter().
		AddPlot("Partition").
			AddDataProvider(NewOptimizerAggregator().
					Optimizer(createOptimizer()).
					StatisticsOptions(NewStatisticsDefaultOptions().
						TrackMinCosts().
						TrackMeanCosts()).
					Iterations(iterations)).
				AddMinCostDataSet().YConverter(plotting.Log10).Done().
				AddMeanCostDataSet().YConverter(plotting.Log10).Done().
				Done().
			Done().
		Draw(8, 4, "plot.png")
}

func createOptimizer() OptimizerInterface {
	popSize := 32
	chromSize := len(partition.List)
	weedRate := 50.0
	mutationProb := 0.05
	generations := 200
	maxGenerationsWithoutImprovements := 15

	return NewIncrementalOptimizer().
		Weeder(NewSimpleWeeder(weedRate)).
		Initializer(NewBinaryRandomInitializer()).
		Selector(NewRouletteWheelCostWeightingSelector()).
		Crossover(NewOnePointCrossover(NewEmptyBinaryChromosome)).
		Mutator(NewBinaryMutator(mutationProb)).
		CostFunction(partition.Cost).
		StopCriterion(NewStopCriterionDefault().
			Max_Generations(generations).
			Min_Cost(0).
			Max_GenerationsWithoutImprovements(maxGenerationsWithoutImprovements)).
		PopSize(popSize).
		ChromSize(chromSize)
}

func setupLogger() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
	    panic(err)
	}

	log.ReplaceLogger(logger)
}