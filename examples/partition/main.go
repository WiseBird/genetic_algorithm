package main

import (
	ga "github.com/WiseBird/genetic_algorithm"
	"math"
	log "github.com/cihub/seelog"
)

var (
	list = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	finalSum = 178
	finalProduct = 3120
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

	sumDiff := float64(sum - finalSum)
	prodDiff := float64(prod - finalProduct)

	return math.Abs(sumDiff) + math.Abs(prodDiff * prodDiff)
}

func main() {
	defer log.Flush()
	setupLogger()

	popSize := 32
	chromSize := len(list)
	weedRate := 50.0
	mutationProb := 0.2
	iterations := 100

	initializer := ga.BinaryRandomInitializerInstance
	weeder := ga.NewSimpleWeeder(weedRate)
	selector := ga.NewRouletteWheelCostWeightingSelector()
	breeder := ga.NewOnePointBreeder(ga.NewEmptyBinaryChromosome)
	mutator := ga.NewBinaryMutator(mutationProb)
	stopCriterion := ga.NewStopCriterionDefault().
		MaxIterations(iterations).
		MinCost(0).
		MaxMinCostAge(15)

	optimizer := ga.NewOptimizer(initializer, weeder, selector, breeder, mutator, cost, popSize, chromSize)
	optimizer.Optimize(stopCriterion)
}

func setupLogger() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")

	if err != nil {
	    panic(err)
	}

	log.ReplaceLogger(logger)
}