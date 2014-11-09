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
	Weeder WeederInterface
	Selector SelectorInterface
	Breeder BreederInterface
	Mutator MutatorInterface
	CostFunction CostFunction

	Population Chromosomes
	popSize int
}
func NewOptimizer(weeder WeederInterface, selector SelectorInterface, breeder BreederInterface, mutator MutatorInterface,
		cost CostFunction) *Optimizer {

	optimizer := &Optimizer{}

	optimizer.Weeder = weeder
	optimizer.Selector = selector
	optimizer.Breeder = breeder
	optimizer.Mutator = mutator
	optimizer.CostFunction = cost

	return optimizer
}

func (optimizer *Optimizer) Optimize(population Chromosomes, iterations int) {
	optimizer.Population = population
	optimizer.popSize = len(optimizer.Population)

	iter := 0
	for ;; {
		log.Infof("ITERATION %d\n\n", iter)

		optimizer.sort()

		if iter >= iterations {
			break
		}

		optimizer.weed()
		optimizer.breed()
		optimizer.mutate()

		iter++
	}
}

func (optimizer *Optimizer) sort() {
	optimizer.Population.SetCost(optimizer.CostFunction)
	sort.Sort(optimizer.Population)

	log.Infof("Best: %v\n", optimizer.Population[0])
	log.Debugf("Population:\n%v\n\n", optimizer.Population)
	log.Flush()
}
func (optimizer *Optimizer) weed() {
	optimizer.Population = optimizer.Weeder.Weed(optimizer.Population)

	log.Tracef("Weeded population:\n%v\n\n", optimizer.Population)
	log.Flush()
}

func (optimizer *Optimizer) breed() {
	newPopulation := optimizer.Population

	optimizer.Selector.Prepare(optimizer.Population)

	for ;; {
		chromsToCross := optimizer.Selector.SelectMany(optimizer.Breeder.ParentsCount())
		log.Debugf("Parents:\n%v\n\n", chromsToCross)

		children := optimizer.Breeder.Crossover(chromsToCross)

		log.Debugf("Children\n%v\n", children)

		for i := 0; i < len(children); i++ {
			newPopulation = append(newPopulation, children[i])

			if len(newPopulation) == optimizer.popSize {
				log.Flush()
				optimizer.Population = newPopulation
				return
			}
		}
	}

}

func (optimizer *Optimizer) mutate() {
	optimizer.Mutator.Mutate(optimizer.Population)
}