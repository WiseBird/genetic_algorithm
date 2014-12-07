package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math/rand"
)

// The edge recombination operator (ERO) is an operator that creates a path that is similar to a set of existing paths (parents) by looking at the edges rather than the vertices.
//
// http://en.wikipedia.org/wiki/Edge_recombination_operator
type EdgeRecombinationCrossover struct {
}

func NewEdgeRecombinationCrossover() *EdgeRecombinationCrossover {
	crossover := new(EdgeRecombinationCrossover)

	return crossover
}

func (crossover *EdgeRecombinationCrossover) ParentsCount() int {
	return 2
}

func (crossover *EdgeRecombinationCrossover) Crossover(parents Chromosomes) Chromosomes {
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

	matrix := crossover.generateMatrix(p1, p2)

	log.Tracef("Cross with %v", matrix)

	c1 := crossover.crossover(p1, p2, matrix)

	return Chromosomes{c1}
}
func (crossover *EdgeRecombinationCrossover) generateMatrix(p1, p2 *OrderedChromosome) map[int]map[int]bool {
	p1genes := p1.OrderedGenes()
	p2genes := p2.OrderedGenes()

	genesLen := p1genes.Len()

	matrix := make(map[int]map[int]bool, genesLen)
	for i := 0; i < genesLen; i++ {
		matrix[p1genes[i]] = make(map[int]bool, 4)
	}

	for i := 0; i < genesLen; i++ {
		ind1 := (i - 1 + genesLen) % genesLen
		ind2 := (i + 1) % genesLen

		matrix[p1genes[i]][p1genes[ind1]] = true
		matrix[p1genes[i]][p1genes[ind2]] = true

		matrix[p2genes[i]][p2genes[ind1]] = true
		matrix[p2genes[i]][p2genes[ind2]] = true
	}

	return matrix
}
func (crossover *EdgeRecombinationCrossover) crossover(p1, p2 *OrderedChromosome, matrix map[int]map[int]bool) (c1 ChromosomeInterface) {
	p1genes := p1.OrderedGenes()
	p2genes := p2.OrderedGenes()

	genesLen := p1genes.Len()

	c1 = NewEmptyOrderedChromosome(genesLen)
	c1genes := c1.Genes().(OrderedGenes)

	crossover.fillChild(c1genes, p1genes, p2genes, matrix)

	return
}
func (crossover *EdgeRecombinationCrossover) fillChild(c, p1, p2 OrderedGenes, matrix map[int]map[int]bool) {
	nextNs := make([]int, 0, 4)

	var n int
	if rand.Intn(2) == 0 {
		n = p1[0]
	} else {
		n = p2[0]
	}
	for i := 0; i < len(c); i++ {
		c[i] = n

		nextNs = nextNs[0:0]
		minLen := 0
		for k, _ := range matrix[n] {
			delete(matrix[k], n)

			if len(nextNs) == 0 || len(matrix[k]) == minLen {
				nextNs = append(nextNs, k)
				minLen = len(matrix[k])
			} else if len(matrix[k]) < minLen {
				nextNs = nextNs[0:0]
				nextNs = append(nextNs, k)
				minLen = len(matrix[k])
			}
		}

		delete(matrix, n)

		if len(nextNs) == 0 {
			if len(matrix) == 0 {
				break
			}
			for k, _ := range matrix {
				n = k
			}
		} else {
			n = nextNs[rand.Intn(len(nextNs))]
		}
	}
}
