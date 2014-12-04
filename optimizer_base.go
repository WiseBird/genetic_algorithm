package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"sort"
)

type OptimizerBase struct {
	OptimizerBaseVirtualMInterface

	initializer InitializerInterface
	weeder      WeederInterface
	selector    SelectorInterface
	crossover   CrossoverInterface
	mutator     MutatorInterface

	costFunction          CostFunction
	statisticsConstructor StatisticsConstructor
	statisticsOptions     StatisticsOptionsInterface
	stopCriterion         StopCriterionInterface

	popSize   int
	chromSize int

	population Chromosomes
}

// MutatorBase's virtual methods
type OptimizerBaseVirtualMInterface interface {
	optimizeInner()
	check()
}

func NewOptimizerBase(virtual OptimizerBaseVirtualMInterface) *OptimizerBase {
	optimizer := new(OptimizerBase)

	optimizer.OptimizerBaseVirtualMInterface = virtual
	optimizer.statisticsConstructor = NewStatisticsDefault
	optimizer.statisticsOptions = NewStatisticsDefaultOptions()

	return optimizer
}

func (optimizer *OptimizerBase) Initializer(initializer InitializerInterface) *OptimizerBase {
	optimizer.initializer = initializer
	return optimizer
}
func (optimizer *OptimizerBase) Selector(selector SelectorInterface) *OptimizerBase {
	optimizer.selector = selector
	return optimizer
}
func (optimizer *OptimizerBase) Crossover(crossover CrossoverInterface) *OptimizerBase {
	optimizer.crossover = crossover
	return optimizer
}
func (optimizer *OptimizerBase) Mutator(mutator MutatorInterface) *OptimizerBase {
	optimizer.mutator = mutator
	return optimizer
}
func (optimizer *OptimizerBase) CostFunction(cost CostFunction) *OptimizerBase {
	optimizer.costFunction = cost
	return optimizer
}
func (optimizer *OptimizerBase) StopCriterion(stopCriterion StopCriterionInterface) *OptimizerBase {
	optimizer.stopCriterion = stopCriterion
	return optimizer
}
func (optimizer *OptimizerBase) StatisticsConstructor(constr StatisticsConstructor) *OptimizerBase {
	optimizer.statisticsConstructor = constr
	return optimizer
}
func (optimizer *OptimizerBase) StatisticsOptions(statisticsOptions StatisticsOptionsInterface) *OptimizerBase {
	optimizer.statisticsOptions = statisticsOptions
	return optimizer
}
func (optimizer *OptimizerBase) PopSize(popSize int) *OptimizerBase {
	optimizer.popSize = popSize
	return optimizer
}
func (optimizer *OptimizerBase) ChromSize(chromSize int) *OptimizerBase {
	optimizer.chromSize = chromSize
	return optimizer
}
func (optimizer *OptimizerBase) check() {
	if optimizer.initializer == nil {
		panic("Initializer must be set")
	}
	if optimizer.selector == nil {
		panic("Selector must be set")
	}
	if optimizer.crossover == nil {
		panic("Crossover must be set")
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

	optimizer.OptimizerBaseVirtualMInterface.check()
}

func (optimizer *OptimizerBase) Optimize() (ChromosomeInterface, StatisticsInterface) {
	optimizer.check()
	optimizer.stopCriterion.Setup(optimizer.statisticsOptions)

	statistics := optimizer.statisticsConstructor(optimizer.statisticsOptions)
	statistics.Start()

	optimizer.population = optimizer.initializer.Init(optimizer.popSize, optimizer.chromSize)
	if len(optimizer.population) == 0 {
		panic("Init population is empty")
	}

	iter := 0
	for {
		log.Infof("GENERATION %d", iter)

		optimizer.sort()
		statistics.OnGeneration(optimizer.population)

		if optimizer.stopCriterion.ShouldStop(statistics) {
			break
		}

		optimizer.OptimizerBaseVirtualMInterface.optimizeInner()

		iter++
	}

	statistics.End()

	return optimizer.population[0], statistics
}
func (optimizer *OptimizerBase) sort() {
	optimizer.population.SetCost(optimizer.costFunction)
	sort.Sort(optimizer.population)

	log.Infof("Best: %v", optimizer.population[0])
	log.Debugf("Population:\n%v\n", optimizer.population)
}

func (optimizer *OptimizerBase) SetupStatisticsOptions() StatisticsOptionsInterface {
	return optimizer.statisticsOptions
}
