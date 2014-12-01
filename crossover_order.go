package genetic_algorithm

import (
	"math/rand"
	log "github.com/cihub/seelog"
)

type OrderCrossover struct {
	canProduceCopiesOfParents bool
}
func NewOrderCrossover() *OrderCrossover {
	crossover := new(OrderCrossover)

	return crossover
}

func (crossover *OrderCrossover) ParentsCount() int {
	return 2
}
func (crossover *OrderCrossover) CanProduceCopiesOfParents(val bool) *OrderCrossover {
	crossover.canProduceCopiesOfParents = val
	return crossover
}

func (crossover *OrderCrossover) Crossover(parents Chromosomes) Chromosomes {
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

	crossPoint1 := crossover.chooseFirstCrossPoint(genesLen)
	crossPoint2 := crossover.chooseSecondCrossPoint(genesLen, crossPoint1)

	log.Tracef("Cross on %d:%d\n", crossPoint1, crossPoint2)

	c1, c2 := crossover.crossover(p1, p2, crossPoint1, crossPoint2)

	return Chromosomes{c1, c2}
}
func (crossover *OrderCrossover) chooseFirstCrossPoint(genesLen int) int {
	return rand.Intn(genesLen)
}
func (crossover *OrderCrossover) chooseSecondCrossPoint(genesLen, crossPoint1 int) int {
	if !crossover.canProduceCopiesOfParents && crossPoint1 == 0 {
		return rand.Intn(genesLen - 1) + 1
	} else {
		return rand.Intn(genesLen - crossPoint1) + 1 + crossPoint1
	}
}
func (crossover *OrderCrossover) crossover(p1, p2 *OrderedChromosome, crossPoint1, crossPoint2 int) (c1, c2 ChromosomeInterface) {
	p1genes := p1.OrderedGenes()
	p2genes := p2.OrderedGenes()

	genesLen := p1genes.Len()

	c1 = NewEmptyOrderedChromosome(genesLen)
	c1genes := c1.Genes().(OrderedGenes)

	c2 = NewEmptyOrderedChromosome(genesLen)
	c2genes := c2.Genes().(OrderedGenes)

	copy(c1genes[crossPoint1:], p1genes[crossPoint1:crossPoint2])
	copy(c2genes[crossPoint1:], p2genes[crossPoint1:crossPoint2])

	crossover.copyFromFillerString(c1genes, p2genes, crossPoint1, crossPoint2)
	crossover.copyFromFillerString(c2genes, p1genes, crossPoint1, crossPoint2)

	return
}
func (crossover *OrderCrossover) copyFromFillerString(child, filler OrderedGenes, crossPoint1, crossPoint2 int){
	genesLen := len(filler)

	cutSection := make(map[int]bool, crossPoint2 - crossPoint1)
	for i := crossPoint1; i < crossPoint2; i++ {
		cutSection[child[i]] = true
	}

	ind := crossPoint2 % genesLen
	for i := 0; i < genesLen; i++ {
		if cutSection[filler[i]] {
			continue
		}

		child[ind] = filler[i]
		ind = (ind + 1) % genesLen
	}
}