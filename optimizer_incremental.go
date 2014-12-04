package genetic_algorithm

import (
	log "github.com/cihub/seelog"
)

type IncrementalOptimizer struct {
	*OptimizerBase
	weeder WeederInterface
}

func NewIncrementalOptimizer() *IncrementalOptimizer {
	optimizer := &IncrementalOptimizer{}

	optimizer.OptimizerBase = NewOptimizerBase(optimizer)

	return optimizer
}

func (optimizer *IncrementalOptimizer) Weeder(weeder WeederInterface) *IncrementalOptimizer {
	optimizer.weeder = weeder
	return optimizer
}

func (optimizer *IncrementalOptimizer) optimizeInner() {
	optimizer.weed()
	optimizer.breed()
	optimizer.mutate()
}
func (optimizer *IncrementalOptimizer) weed() {
	optimizer.population = optimizer.weeder.Weed(optimizer.population)

	log.Tracef("Weeded population:\n%v\n", optimizer.population)
}
func (optimizer *IncrementalOptimizer) breed() {
	newPopulation := optimizer.population

	optimizer.selector.Prepare(optimizer.population)

	for {
		chromsToCross := optimizer.selector.SelectMany(optimizer.crossover.ParentsCount())
		log.Debugf("Parents:\n%v", chromsToCross)

		children := optimizer.crossover.Crossover(chromsToCross)

		log.Debugf("Children\n%v\n", children)

		for i := 0; i < len(children); i++ {
			newPopulation = append(newPopulation, children[i])

			if len(newPopulation) == optimizer.popSize {
				optimizer.population = newPopulation
				return
			}
		}
	}
}
func (optimizer *IncrementalOptimizer) mutate() {
	optimizer.mutator.Mutate(optimizer.population)
}

func (optimizer *IncrementalOptimizer) check() {
	if optimizer.weeder == nil {
		panic("Weeder must be set")
	}
}
