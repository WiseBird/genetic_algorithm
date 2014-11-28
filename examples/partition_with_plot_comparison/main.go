/*
	Compares performance of two optimizers with different mutation probability.
	Draws min/mean cost plots and logs some statistics.
*/

package main

import (
	. "github.com/WiseBird/genetic_algorithm"
	"github.com/WiseBird/genetic_algorithm/plotting"
	log "github.com/cihub/seelog"
	partition "github.com/WiseBird/genetic_algorithm/examples/partition"
	"fmt"
)

func main() {
	defer log.Flush()
	setupLogger()

	pmutate1 := 0.05
	pmutate2 := 0.2
	iterations := 1000

	statisticsAggregator1 := createAggregator()
	optimizer1 := createOptimizer(statisticsAggregator1, pmutate1)

	statisticsAggregator2 := createAggregator()
	optimizer2 := createOptimizer(statisticsAggregator2, pmutate2)

	plotting.NewPlotter().
		AddPlotWithComputations(optimizer1, statisticsAggregator1, iterations).
			Title(fmt.Sprintf("Partition m=%.2f", pmutate1)).
			AddMinCostDataSet().YConverter(plotting.Log10).Done().
			AddMeanCostDataSet().YConverter(plotting.Log10).Done().
			Done().
		AddPlotWithComputations(optimizer2, statisticsAggregator2, iterations).
			Title(fmt.Sprintf("Partition m=%.2f", pmutate2)).
			AddMinCostDataSet().YConverter(plotting.Log10).Done().
			AddMeanCostDataSet().YConverter(plotting.Log10).Done().
			Done().
		Draw(8, 4, "plot.png")

	log.Warnf("Avg duration1: %v", statisticsAggregator1.Duration())
	log.Warnf("Avg duration2: %v", statisticsAggregator2.Duration())

	log.Warnf("Avg generations1: %v", statisticsAggregator1.Generations())
	log.Warnf("Avg generations2: %v", statisticsAggregator2.Generations())

	log.Warnf("Avg minCost1: %v", statisticsAggregator1.MinCost())
	log.Warnf("Avg minCost2: %v", statisticsAggregator2.MinCost())
}

func createAggregator() *StatisticsDefaultAggregator {
	return NewStatisticsDefaultAggregator(NewStatisticsDefaultOptions().
		TrackMinCosts().
		TrackMeanCosts())
}
func createOptimizer(statisticsAggregator *StatisticsDefaultAggregator, mutationProb float64) OptimizerInterface {
	popSize := 32
	chromSize := len(partition.List)
	weedRate := 50.0
	generations := 200

	return NewIncrementalOptimizer().
		Weeder(NewSimpleWeeder(weedRate)).
		Initializer(NewBinaryRandomInitializer()).
		Selector(NewRouletteWheelCostWeightingSelector()).
		Breeder(NewOnePointBreeder(NewEmptyBinaryChromosome)).
		Mutator(NewBinaryMutator(mutationProb)).
		CostFunction(partition.Cost).
		StopCriterion(NewStopCriterionDefault().
			Max_Generations(generations).
			Min_Cost(0).
			Max_GenerationsWithoutImprovements(15)).
		StatisticsOptions(statisticsAggregator.Options()).
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