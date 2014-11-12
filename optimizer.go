package genetic_algorithm

import (
	"sort"
	"math/rand"
	"time"
	log "github.com/cihub/seelog"
)

type CostFunction func(ChromosomeInterface) float64

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Optimizer struct {
	Initializer InitializerInterface
	Weeder WeederInterface
	Selector SelectorInterface
	Breeder BreederInterface
	Mutator MutatorInterface
	CostFunction CostFunction
	StatisticsConstructor StatisticsConstructor

	PopSize int
	ChromSize int

	population Chromosomes
}
func NewOptimizer(initializer InitializerInterface, weeder WeederInterface, selector SelectorInterface, breeder BreederInterface,
		mutator MutatorInterface, cost CostFunction, popSize, chromSize int) *Optimizer {

	optimizer := &Optimizer{}

	optimizer.Initializer = initializer
	optimizer.Weeder = weeder
	optimizer.Selector = selector
	optimizer.Breeder = breeder
	optimizer.Mutator = mutator
	optimizer.CostFunction = cost
	optimizer.StatisticsConstructor = NewStatisticsDefault

	optimizer.PopSize = popSize
	optimizer.ChromSize = chromSize

	return optimizer
}
func (optimizer *Optimizer) WithStatistics(constr StatisticsConstructor) *Optimizer {
	optimizer.StatisticsConstructor = constr
	return optimizer
}

func (optimizer *Optimizer) Optimize(stopCriterion StopCriterionInterface) (ChromosomeInterface, StatisticsInterface) {
	statistics := optimizer.StatisticsConstructor()
	stopCriterion.Setup(statistics)
	statistics.Start()

	optimizer.population = optimizer.Initializer.Init(optimizer.PopSize, optimizer.ChromSize)

	iter := 0
	for ;; {
		log.Infof("GENERATION %d", iter)

		optimizer.sort()
		statistics.OnGeneration(optimizer.population)

		if stopCriterion.ShouldStop(statistics) {
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

func (optimizer *Optimizer) sort() {
	optimizer.population.SetCost(optimizer.CostFunction)
	sort.Sort(optimizer.population)

	log.Infof("Best: %v", optimizer.population[0])
	log.Debugf("Population:\n%v\n\n", optimizer.population)
}
func (optimizer *Optimizer) weed() {
	optimizer.population = optimizer.Weeder.Weed(optimizer.population)

	log.Tracef("Weeded population:\n%v\n\n", optimizer.population)
}

func (optimizer *Optimizer) breed() {
	newPopulation := optimizer.population

	optimizer.Selector.Prepare(optimizer.population)

	for ;; {
		chromsToCross := optimizer.Selector.SelectMany(optimizer.Breeder.ParentsCount())
		log.Debugf("Parents:\n%v\n\n", chromsToCross)

		children := optimizer.Breeder.Crossover(chromsToCross)

		log.Debugf("Children\n%v\n", children)

		for i := 0; i < len(children); i++ {
			newPopulation = append(newPopulation, children[i])

			if len(newPopulation) == optimizer.PopSize {
				optimizer.population = newPopulation
				return
			}
		}
	}

}

func (optimizer *Optimizer) mutate() {
	optimizer.Mutator.Mutate(optimizer.population)
}