package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math/rand"
)

// Crossover for ordered chromosomes.
// Generalization of OrderCrossoverVer2.
// Tends to preserve relative order.
//
// parent1:       A B C D E
// parent2:       d b e a c
//
// mask:          _ * _ * _
// parent1:       A B C D E
//
// child1 step1:  _ B _ D _
//
// parent2:       d b e a c
// filler block:  _ _ e a c
//
// child1:        e B a D c
//
// Source: Modeling Simple Genetic Algorithms for Permutation Problems. Darrell Whitley, Nam-wook Yoo (1995)
// http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.18.3585
type PositionBasedCrossover struct {
	canProduceCopiesOfParents bool
}

func NewPositionBasedCrossover() *PositionBasedCrossover {
	crossover := new(PositionBasedCrossover)

	return crossover
}

func (crossover *PositionBasedCrossover) ParentsCount() int {
	return 2
}
func (crossover *PositionBasedCrossover) CanProduceCopiesOfParents(val bool) *PositionBasedCrossover {
	crossover.canProduceCopiesOfParents = val
	return crossover
}

func (crossover *PositionBasedCrossover) Crossover(parents Chromosomes) Chromosomes {
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
func (crossover *PositionBasedCrossover) generateMask(genesLen int) []int {
	mask := make([]int, 0, genesLen/2)
	for i := 0; i < genesLen; i++ {
		if rand.Intn(2) == 0 {
			mask = append(mask, i)
		}
	}
	return mask
}
func (crossover *PositionBasedCrossover) crossover(p1, p2 *OrderedChromosome, mask []int) (c1, c2 ChromosomeInterface) {
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
func (crossover *PositionBasedCrossover) fillChild(c, p1, p2 OrderedGenes, mask []int) {
	alreadyInChild := make(map[int]bool, len(mask))

	for i := 0; i < len(mask); i++ {
		ind := mask[i]

		c[ind] = p1[ind]
		alreadyInChild[p1[ind]] = true
	}

	p2Ind := 0
	for i := 0; i < len(c); i++ {
		if c[i] != -1 {
			continue
		}

		val := 0
		for {
			val = p2[p2Ind]
			p2Ind++

			if !alreadyInChild[val] {
				break
			}
		}

		c[i] = val
	}

	return
}
