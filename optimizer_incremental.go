package genetic_algorithm

import (
	"sort"
	log "github.com/cihub/seelog"
)

type IncrementalOptimizer struct {
	initializer InitializerInterface
	weeder WeederInterface
	selector SelectorInterface
	breeder BreederInterface
	mutator MutatorInterface

	costFunction CostFunction
	statisticsConstructor StatisticsConstructor
	statisticsOptions StatisticsOptionsInterface
	stopCriterion StopCriterionInterface

	popSize int
	chromSize int

	population Chromosomes
}

func NewIncrementalOptimizer() *IncrementalOptimizer {
	optimizer := &IncrementalOptimizer{}

	optimizer.statisticsConstructor = NewStatisticsDefault
	optimizer.statisticsOptions = NewStatisticsDefaultOptions()

	return optimizer
}

func (optimizer *IncrementalOptimizer) Initializer(initializer InitializerInterface) *IncrementalOptimizer {
	optimizer.initializer = initializer
	return optimizer
}
func (optimizer *IncrementalOptimizer) Weeder(weeder WeederInterface) *IncrementalOptimizer {
	optimizer.weeder = weeder
	return optimizer
}
func (optimizer *IncrementalOptimizer) Selector(selector SelectorInterface) *IncrementalOptimizer {
	optimizer.selector = selector
	return optimizer
}
func (optimizer *IncrementalOptimizer) Breeder(breeder BreederInterface) *IncrementalOptimizer {
	optimizer.breeder = breeder
	return optimizer
}
func (optimizer *IncrementalOptimizer) Mutator(mutator MutatorInterface) *IncrementalOptimizer {
	optimizer.mutator = mutator
	return optimizer
}
func (optimizer *IncrementalOptimizer) CostFunction(cost CostFunction) *IncrementalOptimizer {
	optimizer.costFunction = cost
	return optimizer
}
func (optimizer *IncrementalOptimizer) StopCriterion(stopCriterion StopCriterionInterface) *IncrementalOptimizer {
	optimizer.stopCriterion = stopCriterion
	return optimizer
}
func (optimizer *IncrementalOptimizer) StatisticsConstructor(constr StatisticsConstructor) *IncrementalOptimizer {
	optimizer.statisticsConstructor = constr
	return optimizer
}
func (optimizer *IncrementalOptimizer) StatisticsOptions(statisticsOptions StatisticsOptionsInterface) *IncrementalOptimizer {
	optimizer.statisticsOptions = statisticsOptions
	return optimizer
}
func (optimizer *IncrementalOptimizer) PopSize(popSize int) *IncrementalOptimizer {
	optimizer.popSize = popSize
	return optimizer
}
func (optimizer *IncrementalOptimizer) ChromSize(chromSize int) *IncrementalOptimizer {
	optimizer.chromSize = chromSize
	return optimizer
}

func (optimizer *IncrementalOptimizer) Optimize() (ChromosomeInterface, StatisticsInterface) {
	optimizer.check()
	optimizer.stopCriterion.Setup(optimizer.statisticsOptions)

	statistics := optimizer.statisticsConstructor(optimizer.statisticsOptions)
	statistics.Start()

	optimizer.population = optimizer.initializer.Init(optimizer.popSize, optimizer.chromSize)
	if len(optimizer.population) == 0 {
		panic("Init population is empty")
	}

	iter := 0
	for ;; {
		log.Infof("GENERATION %d", iter)

		optimizer.sort()
		statistics.OnGeneration(optimizer.population)

		if optimizer.stopCriterion.ShouldStop(statistics) {
			break
		}

		optimizer.weed()
		optimizer.breed()
		optimizer.mutate()

		iter++
	}

	statistics.End()

	return optimizer.population[0], statistics
}
func (optimizer *IncrementalOptimizer) sort() {
	optimizer.population.SetCost(optimizer.costFunction)
	sort.Sort(optimizer.population)

	log.Infof("Best: %v", optimizer.population[0])
	log.Debugf("Population:\n%v\n", optimizer.population)
}
func (optimizer *IncrementalOptimizer) weed() {
	optimizer.population = optimizer.weeder.Weed(optimizer.population)

	log.Tracef("Weeded population:\n%v\n", optimizer.population)
}
func (optimizer *IncrementalOptimizer) breed() {
	newPopulation := optimizer.population

	optimizer.selector.Prepare(optimizer.population)

	for ;; {
		chromsToCross := optimizer.selector.SelectMany(optimizer.breeder.ParentsCount())
		log.Debugf("Parents:\n%v", chromsToCross)

		children := optimizer.breeder.Crossover(chromsToCross)

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
	if optimizer.initializer == nil {
		panic("Initializer must be set")
	}
	if optimizer.weeder == nil {
		panic("Weeder must be set")
	}
	if optimizer.selector == nil {
		panic("Selector must be set")
	}
	if optimizer.breeder == nil {
		panic("Breeder must be set")
	}
	if optimizer.mutator == nil {
		panic("Mutator must be set")
	}
	if optimizer.costFunction == nil {
		panic("CostFunction must be set")
	}
	if optimizer.stopCriterion == nil {
		panic("StopCriterion must be set")
	}
	if optimizer.statisticsConstructor == nil {
		panic("StatisticsConstructor must be set")
	}
	if optimizer.statisticsOptions == nil {
		panic("StatisticsOptions must be set")
	}
	if optimizer.popSize <= 0 {
		panic("PopSize must be positive value")
	}
	if optimizer.chromSize <= 0 {
		panic("ChromSize must be positive value")
	}
}