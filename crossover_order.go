package genetic_algorithm

import (
	log "github.com/cihub/seelog"
)

type orderCrossover struct {
	virtualMethods            orderCrossoverVirtualMInterface
	canProduceCopiesOfParents bool
}
type orderCrossoverVirtualMInterface interface {
	copyFromFillerString(child, filler OrderedGenes, crossPoint1, crossPoint2 int)
}

func newOrderCrossover(vm orderCrossoverVirtualMInterface) *orderCrossover {
	crossover := new(orderCrossover)

	crossover.virtualMethods = vm

	return crossover
}

func (crossover *orderCrossover) ParentsCount() int {
	return 2
}
func (crossover *orderCrossover) CanProduceCopiesOfParents(val bool) *orderCrossover {
	crossover.canProduceCopiesOfParents = val
	return crossover
}

func (crossover *orderCrossover) Crossover(parents Chromosomes) Chromosomes {
	if len(parents) != crossover.ParentsCount() {
		panic("Incorrect parents count")
	}

	p1, ok := parents[0].(*OrderedChromosome)
	if !ok {
		panic("Expects OrderedChromosome")
	}
	p2, ok := parents[1].(*OrderedChromosome)
	if !ok {
		panic("Expects OrderedChromosome")
	}

	genesLen := p1.Genes().Len()

	if genesLen != p2.Genes().Len() {
		panic("Crossover do not support different chromosome size")
	}

	if !crossover.canProduceCopiesOfParents && genesLen < 2 {
		panic("Crossover can only produce copies of parents if genesLen < 2")
	}

	crossPoint1, crossPoint2 := chooseTwoPointCrossSection(genesLen, crossover.canProduceCopiesOfParents)

	log.Tracef("Cross on %d:%d", crossPoint1, crossPoint2)

	c1, c2 := crossover.crossover(p1, p2, crossPoint1, crossPoint2)

	return Chromosomes{c1, c2}
}
func (crossover *orderCrossover) crossover(p1, p2 *OrderedChromosome, crossPoint1, crossPoint2 int) (c1, c2 ChromosomeInterface) {
	p1genes := p1.OrderedGenes()
	p2genes := p2.OrderedGenes()

	genesLen := p1genes.Len()

	c1 = NewEmptyOrderedChromosome(genesLen)
	c1genes := c1.Genes().(OrderedGenes)

	c2 = NewEmptyOrderedChromosome(genesLen)
	c2genes := c2.Genes().(OrderedGenes)

	copy(c1genes[crossPoint1:], p1genes[crossPoint1:crossPoint2])
	copy(c2genes[crossPoint1:], p2genes[crossPoint1:crossPoint2])

	crossover.virtualMethods.copyFromFillerString(c1genes, p2genes, crossPoint1, crossPoint2)
	crossover.virtualMethods.copyFromFillerString(c2genes, p1genes, crossPoint1, crossPoint2)

	return
}

// Crossover for ordered chromosomes.
// Tends to preserve relative order.
//
// parent1:       A B C D E
// parent2:       d b e a c
//
// cross section: _ * * _ _
//
// child1 step1:  _ B C _ _
//
// parent2:       d b e a c
// filler block:  d _ e a _
// The first element of filler block is added at the end of the cross setion.
//
// child1:        a B C d e
//
// Source: Modeling Simple Genetic Algorithms for Permutation Problems. Darrell Whitley , Nam-wook Yoo (1995)
// http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.18.3585
type OrderCrossoverVer1 struct{}

func NewOrderCrossoverVer1() *orderCrossover {
	return newOrderCrossover(new(OrderCrossoverVer1))
}
func (crossover *OrderCrossoverVer1) copyFromFillerString(child, filler OrderedGenes, crossPoint1, crossPoint2 int) {
	genesLen := len(filler)

	crossSection := make(map[int]bool, crossPoint2-crossPoint1)
	for i := crossPoint1; i < crossPoint2; i++ {
		crossSection[child[i]] = true
	}

	ind := crossPoint2 % genesLen
	for i := 0; i < genesLen; i++ {
		if crossSection[filler[i]] {
			continue
		}

		child[ind] = filler[i]
		ind = (ind + 1) % genesLen
	}
}

// Crossover for ordered chromosomes.
// Tends to preserve relative order.
//
// parent1:       A B C D E
// parent2:       d b e a c
//
// cross section: _ * * _ _
//
// child1 step1:  _ B C _ _
//
// parent2:       d b e a c
// filler block:  d _ e a _
// The first element of filler block is added at the start of the child.
//
// child1:        d B C e a
//
// Source: On Genetic Crossover Operators for Relative Order Preservation
// http://www.dmi.unict.it/mpavone/nc-cs/materiale/moscato89.pdf
type OrderCrossoverVer2 struct{}

func NewOrderCrossoverVer2() *orderCrossover {
	return newOrderCrossover(new(OrderCrossoverVer2))
}
func (crossover *OrderCrossoverVer2) copyFromFillerString(child, filler OrderedGenes, crossPoint1, crossPoint2 int) {
	genesLen := len(filler)

	crossSection := make(map[int]bool, crossPoint2-crossPoint1)
	for i := crossPoint1; i < crossPoint2; i++ {
		crossSection[child[i]] = true
	}

	ind := 0
	for i := 0; i < genesLen; i++ {
		if crossSection[filler[i]] {
			continue
		}
		if ind == crossPoint1 {
			ind = crossPoint2
		}

		child[ind] = filler[i]
		ind++
	}
}
