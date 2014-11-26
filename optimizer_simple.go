package genetic_algorithm

import (
	"math/rand"
	log "github.com/cihub/seelog"
)

type SimpleOptimizer struct {
	*OptimizerBase

	crossoverProbability float64
	elitism int

	secondPopulation Chromosomes
}

func NewSimpleOptimizer() *SimpleOptimizer {
	optimizer := &SimpleOptimizer{}

	optimizer.OptimizerBase = NewOptimizerBase(optimizer)

	return optimizer
}

func (optimizer *SimpleOptimizer) Elitism(elitism int) *SimpleOptimizer {
	optimizer.elitism = elitism
	return optimizer
}
func (optimizer *SimpleOptimizer) CrossoverProbability(crossoverProbability float64) *SimpleOptimizer {
	optimizer.crossoverProbability = crossoverProbability
	return optimizer
}

func (optimizer *SimpleOptimizer) optimizeInner() {
	optimizer.breed()
	optimizer.mutate()
}
func (optimizer *SimpleOptimizer) breed() {
	if optimizer.secondPopulation == nil {
		optimizer.secondPopulation = make(Chromosomes, optimizer.elitism, optimizer.popSize)
	} else {
		optimizer.secondPopulation = optimizer.secondPopulation[:optimizer.elitism]
	}
	copy(optimizer.secondPopulation, optimizer.population[:optimizer.elitism])

	optimizer.selector.Prepare(optimizer.population)

	for ;; {
		chromsToCross := optimizer.selector.SelectMany(optimizer.breeder.ParentsCount())
		log.Debugf("Parents:\n%v", chromsToCross)

		var newChromosomes Chromosomes

		if optimizer.crossoverProbability > rand.Float64() {
			newChromosomes = optimizer.breeder.Crossover(chromsToCross)
			log.Debugf("Children\n%v\n", newChromosomes)
		} else {
			log.Debugf("Parents go to the new generation\n")
			newChromosomes = chromsToCross
		}

		for i := 0; i < len(newChromosomes); i++ {
			optimizer.secondPopulation = append(optimizer.secondPopulation, newChromosomes[i])

			if len(optimizer.secondPopulation) == optimizer.popSize {
				optimizer.population, optimizer.secondPopulation = optimizer.secondPopulation, optimizer.population
				return
			}
		}
	}
}
func (optimizer *SimpleOptimizer) mutate() {
	optimizer.mutator.Mutate(optimizer.population)
}

func (optimizer *SimpleOptimizer) check() {
	if optimizer.elitism < 0 {
		panic("Elitism must be positive")
	}
	if optimizer.crossoverProbability <= 0 || optimizer.crossoverProbability > 1 {
		panic("CrossoverProbability must be in (0, 1] range")
	}
}