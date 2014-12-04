package genetic_algorithm

import (
	log "github.com/cihub/seelog"
)

// Crossover for ordered chromosomes.
// Tends to preserve relative order.
//
// Ordering genetic algorithms and deception. Hillol Kargupta, Ka. Lyanmoy Deb, David E. Goldberg (1992)
// http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.94.6805
type RelativeOrderingCrossover struct {
	preservedGenes int
}

func NewRelativeOrderingCrossover(preservedGenes int) *RelativeOrderingCrossover {
	if preservedGenes < 1 {
		panic("preservedGenes must be positive")
	}

	crossover := new(RelativeOrderingCrossover)

	crossover.preservedGenes = preservedGenes

	return crossover
}

func (crossover *RelativeOrderingCrossover) ParentsCount() int {
	return 2
}

func (crossover *RelativeOrderingCrossover) Crossover(parents Chromosomes) Chromosomes {
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

	if crossover.preservedGenes <= genesLen {
		log.Warnf("ROX will produce copies of parent because chromosome len lesser than preservedGenes")
	}

	indexes := chooseDifferentRandomNumbers(crossover.preservedGenes, genesLen)

	log.Tracef("Cross with %v", indexes)

	c1, c2 := crossover.crossover(p1, p2, indexes)

	return Chromosomes{c1, c2}
}
func (crossover *RelativeOrderingCrossover) crossover(p1, p2 *OrderedChromosome, indexes []int) (c1, c2 ChromosomeInterface) {
	p1genes := p1.OrderedGenes()
	p2genes := p2.OrderedGenes()

	genesLen := p1genes.Len()

	c1 = NewEmptyOrderedChromosome(genesLen)
	c1genes := c1.Genes().(OrderedGenes)

	c2 = NewEmptyOrderedChromosome(genesLen)
	c2genes := c2.Genes().(OrderedGenes)

	crossover.fillChild(c1genes, p1genes, p2genes, indexes)
	crossover.fillChild(c2genes, p2genes, p1genes, indexes)

	return
}
func (crossover *RelativeOrderingCrossover) fillChild(c, p1, p2 OrderedGenes, indexes []int) {
	waitingList := make([]int, 0)

	preserveList := make([]int, len(indexes))
	preserveMap := make(map[int]bool, len(indexes))
	for i := 0; i < len(indexes); i++ {
		preserveList[i] = p1[indexes[i]]
		preserveMap[preserveList[i]] = true
	}

	ind := 0
	for _, val := range p2 {
		if !preserveMap[val] {
			c[ind] = val
			ind++
		} else if preserveList[0] == val {
			for {
				c[ind] = preserveList[0]
				ind++

				preserveMap[preserveList[0]] = false
				preserveList = preserveList[1:]

				if len(preserveList) == 0 {
					break
				}

				inWaitingList := false
				for _, v := range waitingList {
					if v == preserveList[0] {
						inWaitingList = true
						break
					}
				}

				if !inWaitingList {
					break
				}
			}
		} else {
			waitingList = append(waitingList, val)
		}
	}

	return
}
