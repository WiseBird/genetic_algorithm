package genetic_algorithm

import (
	"math/rand"
)

type OrderedRandomInitializer struct {
}

func NewOrderedRandomInitializer() *OrderedRandomInitializer {
	initializer := new(OrderedRandomInitializer)

	return initializer
}
func (initializer *OrderedRandomInitializer) Init(count, chromSize int) Chromosomes {
	result := make([]ChromosomeInterface, count)

	for chromeInd := 0; chromeInd < count; chromeInd++ {

		genes := make(OrderedGenes, chromSize)
		// http://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
		for geneInd := 0; geneInd < chromSize; geneInd++ {
			randInd := rand.Intn(geneInd + 1)

			genes[geneInd] = genes[randInd]
			genes[randInd] = geneInd
		}

		result[chromeInd] = NewOrderedChromosome(genes)
	}

	return result
}
