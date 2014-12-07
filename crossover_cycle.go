package genetic_algorithm

import (
	log "github.com/cihub/seelog"
)

// Crossover for ordered chromosomes.
// Tends to preserve absolute order.
type CycleCrossover struct {
}

func NewCycleCrossover() *CycleCrossover {
	crossover := new(CycleCrossover)

	return crossover
}

func (crossover *CycleCrossover) ParentsCount() int {
	return 2
}

func (crossover *CycleCrossover) Crossover(parents Chromosomes) Chromosomes {
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

	mask := crossover.generateMask(p1, p2)

	log.Tracef("Cross with %v", mask)

	c1, c2 := crossover.crossover(p1, p2, mask)

	return Chromosomes{c1, c2}
}
func (crossover *CycleCrossover) generateMask(p1, p2 *OrderedChromosome) []int {
	p1genes := p1.OrderedGenes()
	p2genes := p2.OrderedGenes()

	genesLen := p1genes.Len()

	mask := make([]int, genesLen)

	ind := 0
	parentInd := 1
	for {
		startVal := p1genes[ind]
		for {
			mask[ind] = parentInd
			val := p2genes[ind]
			if val == startVal {
				break
			}

			ind = p1genes.Ind(val)
		}

		parentInd = (parentInd % 2) + 1
		ind = -1
		for i := 0; i < genesLen; i++ {
			if mask[i] == 0 {
				ind = i
			}
		}

		if ind == -1 {
			break
		}
	}

	return mask
}
func (crossover *CycleCrossover) crossover(p1, p2 *OrderedChromosome, mask []int) (c1, c2 ChromosomeInterface) {
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
func (crossover *CycleCrossover) fillChild(c, p1, p2 OrderedGenes, mask []int) {
	for i := 0; i < len(c); i++ {
		if mask[i] == 1 {
			c[i] = p1[i]
		} else {
			c[i] = p2[i]
		}
	}

	return
}
