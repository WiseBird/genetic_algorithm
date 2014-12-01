package genetic_algorithm

import (
	"math/rand"
)

type BinaryRandomInitializer struct {
}
func NewBinaryRandomInitializer() *BinaryRandomInitializer {
	initializer := new(BinaryRandomInitializer)	

	return initializer
}
func (initializer *BinaryRandomInitializer) Init(count, chromSize int) Chromosomes {
	result := make([]ChromosomeInterface, count)

	for chromeInd := 0; chromeInd < count; chromeInd++ {

		genes := make(BinaryGenes, chromSize)
		for geneInd := 0; geneInd < chromSize; geneInd++ {
			x := rand.Intn(2)
			genes[geneInd] = x == 1
		}

		result[chromeInd] = NewBinaryChromosome(genes)
	}

	return result
}