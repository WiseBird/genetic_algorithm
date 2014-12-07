package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math/rand"
)

// Crossover for ordered chromosomes.
// Tends to preserve absolute order.
//
// Source: Introduction to Genetic Algorithms. S.N. Sivanandam, S. N. Deepa (2008)
// http://www.amazon.com/Introduction-Genetic-Algorithms-S-N-Sivanandam/dp/354073189X/
type PrecedencePreservativeCrossover struct {
}

func NewPrecedencePreservativeCrossover() *PrecedencePreservativeCrossover {
	crossover := new(PrecedencePreservativeCrossover)

	return crossover
}

func (crossover *PrecedencePreservativeCrossover) ParentsCount() int {
	return 2
}

func (crossover *PrecedencePreservativeCrossover) Crossover(parents Chromosomes) Chromosomes {
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

	mask := crossover.generateMask(genesLen)

	log.Tracef("Cross with %v", mask)

	c1, c2 := crossover.crossover(p1, p2, mask)

	return Chromosomes{c1, c2}
}
func (crossover *PrecedencePreservativeCrossover) generateMask(genesLen int) []int {
	mask := make([]int, genesLen)
	for i := 0; i < genesLen; i++ {
		if rand.Intn(2) == 0 {
			mask[i] = 1
		} else {
			mask[i] = 2
		}
	}
	return mask
}
func (crossover *PrecedencePreservativeCrossover) crossover(p1, p2 *OrderedChromosome, mask []int) (c1, c2 ChromosomeInterface) {
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
func (crossover *PrecedencePreservativeCrossover) fillChild(c, p1, p2 OrderedGenes, mask []int) {
	alreadyInChild := make(map[int]bool, len(c))

	p1Ind := 0
	p2Ind := 0
	for i := 0; i < len(c); i++ {
		if mask[i] == 1 {
			for alreadyInChild[p1[p1Ind]] {
				p1Ind++
			}

			c[i] = p1[p1Ind]
			alreadyInChild[p1[p1Ind]] = true
			p1Ind++
		} else {
			for alreadyInChild[p2[p2Ind]] {
				p2Ind++
			}

			c[i] = p2[p2Ind]
			alreadyInChild[p2[p2Ind]] = true
			p2Ind++
		}
	}
}
