package genetic_algorithm

import (
	log "github.com/cihub/seelog"
)

// Crossover for ordered chromosomes.
// Tends to preserve absolute order.
//
// Source: Modeling Simple Genetic Algorithms for Permutation Problems. Darrell Whitley , Nam-wook Yoo (1995)
// http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.18.3585
type PartiallyMappedCrossover struct {
	canProduceCopiesOfParents bool
}

func NewPartiallyMappedCrossover() *PartiallyMappedCrossover {
	crossover := new(PartiallyMappedCrossover)

	return crossover
}

func (crossover *PartiallyMappedCrossover) ParentsCount() int {
	return 2
}
func (crossover *PartiallyMappedCrossover) CanProduceCopiesOfParents(val bool) *PartiallyMappedCrossover {
	crossover.canProduceCopiesOfParents = val
	return crossover
}

func (crossover *PartiallyMappedCrossover) Crossover(parents Chromosomes) Chromosomes {
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
func (crossover *PartiallyMappedCrossover) crossover(p1, p2 *OrderedChromosome, crossPoint1, crossPoint2 int) (c1, c2 ChromosomeInterface) {
	p1genes := p1.OrderedGenes()
	p2genes := p2.OrderedGenes()

	genesLen := p1genes.Len()

	c1 = NewEmptyOrderedChromosome(genesLen)
	c1genes := c1.Genes().(OrderedGenes)

	c2 = NewEmptyOrderedChromosome(genesLen)
	c2genes := c2.Genes().(OrderedGenes)

	copy(c1genes[crossPoint1:], p1genes[crossPoint1:crossPoint2])
	copy(c2genes[crossPoint1:], p2genes[crossPoint1:crossPoint2])

	crossover.fillChild(c1genes, p1genes, p2genes, crossPoint1, crossPoint2)
	crossover.fillChild(c2genes, p2genes, p1genes, crossPoint1, crossPoint2)

	return
}
func (crossover *PartiallyMappedCrossover) fillChild(c, p1, p2 OrderedGenes, crossPoint1, crossPoint2 int) {
	alreadyInChild := make(map[int]bool, crossPoint2-crossPoint1)

	for i := crossPoint1; i < crossPoint2; i++ {
		alreadyInChild[c[i]] = true
	}

	for i := crossPoint1; i < crossPoint2; i++ {
		val := p2[i]
		if alreadyInChild[val] {
			continue
		}

		ind := i
		for ind >= crossPoint1 && ind < crossPoint2 {
			ind = p2.Ind(p1[ind])
		}

		c[ind] = val
	}

	for i := 0; i < len(c); i++ {
		if c[i] == -1 {
			c[i] = p2[i]
		}
	}
}
