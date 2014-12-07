package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math/rand"
)

// Crossover for ordered chromosomes.
// Tends to preserve relative order.
//
// parent1:       A B C D E F
// parent2:       d f b e a c
//
// mask:          * * _ * _ _
// parent1:       A B C D E F
// The genes 'A B D' will be copied in child in order from parent1 and in paces from parent2
//
// child1 step1:  A _ B _ D _
//
// parent2:       d f b e a c
// filler block:  _ f _ e _ c
//
// child1:        A f B e D c
// Despite the fact that OrderBased behaves differently from Position crossover they are identical in expectation.
//
// Source: Modeling Simple Genetic Algorithms for Permutation Problems. Darrell Whitley , Nam-wook Yoo (1995)
// http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.18.3585
type OrderBasedCrossover struct {
	canProduceCopiesOfParents bool
}

func NewOrderBasedCrossover() *OrderBasedCrossover {
	crossover := new(OrderBasedCrossover)

	return crossover
}

func (crossover *OrderBasedCrossover) ParentsCount() int {
	return 2
}
func (crossover *OrderBasedCrossover) CanProduceCopiesOfParents(val bool) *OrderBasedCrossover {
	crossover.canProduceCopiesOfParents = val
	return crossover
}

func (crossover *OrderBasedCrossover) Crossover(parents Chromosomes) Chromosomes {
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

	mask := crossover.generateMask(genesLen)

	log.Tracef("Cross with %v", mask)

	c1, c2 := crossover.crossover(p1, p2, mask)

	return Chromosomes{c1, c2}
}
func (crossover *OrderBasedCrossover) generateMask(genesLen int) []int {
	mask := make([]int, 0, genesLen/2)
	for i := 0; i < genesLen; i++ {
		if rand.Intn(2) == 0 {
			mask = append(mask, i)
		}
	}
	return mask
}
func (crossover *OrderBasedCrossover) crossover(p1, p2 *OrderedChromosome, mask []int) (c1, c2 ChromosomeInterface) {
	p1genes := p1.OrderedGenes()
	p2genes := p2.OrderedGenes()

	genesLen := p1genes.Len()

	c1 = NewEmptyOrderedChromosome(genesLen)
	c1genes := c1.Genes().(OrderedGenes)

	c2 = NewEmptyOrderedChromosome(genesLen)
	c2genes := c2.Genes().(OrderedGenes)

	crossover.fillChild(c1genes, p1genes, p2genes, mask)
	crossover.fillChild(c2genes, p2genes, p1genes, mask)

	return
}
func (crossover *OrderBasedCrossover) fillChild(c, p1, p2 OrderedGenes, mask []int) {
	inMask := make(map[int]bool, len(mask))
	for i := 0; i < len(mask); i++ {
		ind := mask[i]
		inMask[p1[ind]] = true
	}

	maskInd := 0
	for i := 0; i < len(c); i++ {
		val := p2[i]
		if inMask[val] {
			c[i] = p1[mask[maskInd]]
			maskInd++
		} else {
			c[i] = val
		}
	}

	return
}
