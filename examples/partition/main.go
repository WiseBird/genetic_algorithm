/*
The task is to split list of numbers in two groups such that the sum of one is 178 and the product of the other is 3120.
*/

package main

import (
	. "github.com/WiseBird/genetic_algorithm"
	"math"
	log "github.com/cihub/seelog"
)

var (
	List = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	finalSum = 178
	finalProduct = 3120
)

func Cost(c ChromosomeInterface) float64 {
	bc := c.(*BinaryChromosome)
	sum := 0
	prod := 1

	genes := bc.Genes().(BinaryGenes)

	for i := 0; i < len(genes); i++ {
		if genes[i] {
			prod *= List[i]
		} else {
			sum += List[i]
		}
	}

	sumDiff := float64(sum - finalSum)
	prodDiff := float64(prod - finalProduct)

	return math.Abs(sumDiff) + math.Abs(prodDiff)
}

func main() {
	defer log.Flush()
	setupLogger()

	optimizer := createOptimizer()
	best, _ := optimizer.Optimize()

	log.Warnf("Baest: %v", best)
}

func createOptimizer() OptimizerInterface {
	popSize := 32
	chromSize := len(List)
	weedRate := 50.0
	mutationProb := 0.05
	generations := 200

	return NewIncrementalOptimizer().
		Weeder(NewSimpleWeeder(weedRate)).
		Initializer(NewBinaryRandomInitializer()).
		Selector(NewRouletteWheelCostWeightingSelector()).
		Breeder(NewOnePointBreeder(NewEmptyBinaryChromosome)).
		Mutator(NewBinaryMutator(mutationProb)).
		CostFunction(Cost).
		StopCriterion(NewStopCriterionDefault().
			Max_Generations(generations).
			Min_Cost(0).
			Max_MinCostAge(15)).
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