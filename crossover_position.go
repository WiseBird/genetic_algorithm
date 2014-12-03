package genetic_algorithm

import (
	"math/rand"
	log "github.com/cihub/seelog"
)
// Crossover for ordered chromosomes.
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
// Source: Modeling Simple Genetic Algorithms for Permutation Problems. Darrell Whitley , Nam-wook Yoo (1995)
// http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.18.3585
type PositionCrossover struct {
	canProduceCopiesOfParents bool
}
func NewPositionCrossover() *PositionCrossover {
	crossover := new(PositionCrossover)

	return crossover
}
// Despite the fact that UOX behaves differently from Position crossover they are identical in expectation.
//
// Source: Modeling Simple Genetic Algorithms for Permutation Problems. Darrell Whitley , Nam-wook Yoo (1995)
// http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.18.3585
func NewUniformOrderingCrossover() *PositionCrossover {
	return NewPositionCrossover()
}

func (crossover *PositionCrossover) ParentsCount() int {
	return 2
}
func (crossover *PositionCrossover) CanProduceCopiesOfParents(val bool) *PositionCrossover {
	crossover.canProduceCopiesOfParents = val
	return crossover
}

func (crossover *PositionCrossover) Crossover(parents Chromosomes) Chromosomes {
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

	mask := make([]int, genesLen / 2)
	for i := 0; i < genesLen; i++ {
		if rand.Intn(2) == 0 {
			mask = append(mask, i)
		}
	}

	log.Tracef("Cross with %v", mask)

	c1, c2 := crossover.crossover(p1, p2, mask)

	return Chromosomes{c1, c2}
}
func (crossover *PositionCrossover) crossover(p1, p2 *OrderedChromosome, mask []int) (c1, c2 ChromosomeInterface) {
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
func (crossover *PositionCrossover) fillChild(c, p1, p2 OrderedGenes, mask []int) {
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
		for ;; {
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
