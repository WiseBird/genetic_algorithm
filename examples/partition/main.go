package main

import (
	ga "github.com/WiseBird/genetic_algorithm"
	"math"
	log "github.com/cihub/seelog"
)

var (
	list = []int{1,2,3,4,5,6,7,8,9,10,11,12}
	finalSum = 49
	finalProduct = 2520
)

func cost(c ga.ChromosomeInterface) float64 {
	bc := c.(*ga.BinaryChromosome)
	sum := 0
	prod := 1

	genes := bc.Genes().(ga.BinaryGenes)

	for i := 0; i < len(genes); i++ {
		if genes[i] {
			prod *= list[i]
		} else {
			sum += list[i]
		}
	}

	sumDiff := sum - finalSum
	prodDiff := prod - finalProduct

	return math.Sqrt(float64(sumDiff * sumDiff + prodDiff * prodDiff))
}

func main() {
	defer log.Flush()
	setupLogger()

	popSize := 32
	chromSize := len(list)
	weedRate := 50.0
	mutationProb := 0.2
	iterations := 100

	weeder := ga.NewSimpleWeeder(weedRate)
	selector := ga.NewRouletteWheelCostWeightingSelector()
	breeder := ga.NewOnePointBreeder(ga.NewEmptyBinaryChromosome)
	mutator := ga.NewBinaryMutator(mutationProb)

	optimizer := ga.NewOptimizer(weeder, selector, breeder, mutator, cost)

	population := ga.BinaryRandomInitializerInstance.Init(popSize, chromSize)
	optimizer.Optimize(population, iterations)
}

func setupLogger() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")

	if err != nil {
	    panic(err)
	}

	log.ReplaceLogger(logger)
}