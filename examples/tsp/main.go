/*
Example of solving tsp.
*/

package main

import (
	. "github.com/WiseBird/genetic_algorithm"
	"math"
	log "github.com/cihub/seelog"
)


type City struct {
	X float64
	Y float64
}

func Cost(c ChromosomeInterface) float64 {
	return CalcPath(cities, c)
}

func CalcPath(cities []City, c ChromosomeInterface) float64 {
	oc := c.(*OrderedChromosome)
	genes := oc.OrderedGenes()

	path := float64(0)

	for i := 0; i < len(genes); i++ {
		if i == 0 {
			path += calcDistance(cities[genes[0]], cities[genes[len(genes)-1]])
		} else {
			path += calcDistance(cities[genes[i-1]], cities[genes[i]])
		}
	}

	return path
}
func calcDistance(c1 City, c2 City) float64 {
	return math.Sqrt(
		math.Pow(c1.X - c2.X, 2) + 
		math.Pow(c1.Y - c2.Y, 2))
}

func main() {
	defer log.Flush()
	setupLogger()

	optimizer := createOptimizer()
	best, _ := optimizer.Optimize()

	log.Warnf("Best: %v", best)
}

func createOptimizer() OptimizerInterface {
	popSize := 32
	chromSize := len(cities)
	weedRate := 50.0
	mutationProb := 0.05
	generations := 200

	return NewIncrementalOptimizer().
		Weeder(NewSimpleWeeder(weedRate)).
		Initializer(NewOrderedRandomInitializer()).
		Selector(NewRouletteWheelRankWeightingSelector()).
		Crossover(NewOrderCrossover()).
		Mutator(NewOrderedSwapMutator(mutationProb)).
		CostFunction(Cost).
		StopCriterion(NewStopCriterionDefault().
			Max_Generations(generations).
			Max_GenerationsWithoutImprovements(15)).
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