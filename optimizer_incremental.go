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

func NewIncrementalOptimizer(
		initializer InitializerInterface, 
		weeder WeederInterface, 
		selector SelectorInterface, 
		breeder BreederInterface,
		mutator MutatorInterface, 
		cost CostFunction, 
		stopCriterion StopCriterionInterface,
		popSize, 
		chromSize int) *IncrementalOptimizer {

	optimizer := NewEmptyIncrementalOptimizer()

	optimizer.initializer = initializer
	optimizer.weeder = weeder
	optimizer.selector = selector
	optimizer.breeder = breeder
	optimizer.mutator = mutator

	optimizer.costFunction = cost
	optimizer.stopCriterion = stopCriterion

	optimizer.popSize = popSize
	optimizer.chromSize = chromSize

	return optimizer
}
func NewEmptyIncrementalOptimizer() *IncrementalOptimizer {
	optimizer := &IncrementalOptimizer{}

	optimizer.statisticsConstructor = NewStatisticsDefault
	optimizer.statisticsOptions = NewStatisticsDefaultOptions()

	return optimizer
}

func (optimizer *IncrementalOptimizer) WithInitializer(initializer InitializerInterface) *IncrementalOptimizer {
	optimizer.initializer = initializer
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithWeeder(weeder WeederInterface) *IncrementalOptimizer {
	optimizer.weeder = weeder
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithSelector(selector SelectorInterface) *IncrementalOptimizer {
	optimizer.selector = selector
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithBreeder(breeder BreederInterface) *IncrementalOptimizer {
	optimizer.breeder = breeder
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithMutator(mutator MutatorInterface) *IncrementalOptimizer {
	optimizer.mutator = mutator
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithCostFunction(cost CostFunction) *IncrementalOptimizer {
	optimizer.costFunction = cost
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithStopCriterion(stopCriterion StopCriterionInterface) *IncrementalOptimizer {
	optimizer.stopCriterion = stopCriterion
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithStatistics(constr StatisticsConstructor) *IncrementalOptimizer {
	optimizer.statisticsConstructor = constr
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithStatisticsOptions(statisticsOptions StatisticsOptionsInterface) *IncrementalOptimizer {
	optimizer.statisticsOptions = statisticsOptions
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithPopSize(popSize int) *IncrementalOptimizer {
	optimizer.popSize = popSize
	return optimizer
}
func (optimizer *IncrementalOptimizer) WithChromSize(chromSize int) *IncrementalOptimizer {
	optimizer.chromSize = chromSize
	return optimizer
}

func (optimizer *IncrementalOptimizer) Optimize() (ChromosomeInterface, StatisticsInterface) {
	statistics := optimizer.statisticsConstructor()
	optimizer.stopCriterion.Setup(optimizer.statisticsOptions)
	statistics.SetOptions(optimizer.statisticsOptions)
	statistics.Start()

	optimizer.population = optimizer.initializer.Init(optimizer.popSize, optimizer.chromSize)

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