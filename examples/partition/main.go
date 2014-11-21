package main

import (
	. "github.com/WiseBird/genetic_algorithm"
	"github.com/WiseBird/genetic_algorithm/plotting"
	"math"
	log "github.com/cihub/seelog"
)

var (
	list = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	finalSum = 178
	finalProduct = 3120
)

func cost(c ChromosomeInterface) float64 {
	bc := c.(*BinaryChromosome)
	sum := 0
	prod := 1

	genes := bc.Genes().(BinaryGenes)

	for i := 0; i < len(genes); i++ {
		if genes[i] {
			prod *= list[i]
		} else {
			sum += list[i]
		}
	}

	sumDiff := float64(sum - finalSum)
	prodDiff := float64(prod - finalProduct)

	return math.Abs(sumDiff) + math.Abs(prodDiff * prodDiff)
}

func main() {
	defer log.Flush()
	setupLogger()

	iterations := 100

	statisticsOptions := NewStatisticsDefaultOptions().
			TrackMinCosts().
			TrackMeanCosts()
	statisticsAggregator := NewStatisticsDefaultAggregator(statisticsOptions)
	optimizer := createOptimizer(statisticsOptions)

	plotting.NewPlotter().
		AddPlotWithComputations(optimizer, statisticsAggregator, iterations).
			AddMinCostDataSet().YConverter(plotting.Log10).Done().
			AddMeanCostDataSet().YConverter(plotting.Log10).Done().
		Done().
		Draw(8, 4, "points.png")

	log.Warnf("Duration: %v", statisticsAggregator.Duration())
	log.Warnf("MinCosts: %v", statisticsAggregator.MinCosts())
	log.Warnf("MeanCosts: %v", statisticsAggregator.MeanCosts())
}

func createOptimizer(statisticsOptions StatisticsOptionsInterface) OptimizerInterface {
	popSize := 32
	chromSize := len(list)
	weedRate := 50.0
	mutationProb := 0.05
	generations := 200

	return NewIncrementalOptimizer().
		Initializer(NewBinaryRandomInitializer()).
		Weeder(NewSimpleWeeder(weedRate)).
		Selector(NewRouletteWheelCostWeightingSelector()).
		Breeder(NewOnePointBreeder(NewEmptyBinaryChromosome)).
		Mutator(NewBinaryMutator(mutationProb)).
		CostFunction(cost).
		StopCriterion(NewStopCriterionDefault().
			Max_Generations(generations).
			Min_Cost(0).
			Max_MinCostAge(15)).
		StatisticsOptions(statisticsOptions).
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