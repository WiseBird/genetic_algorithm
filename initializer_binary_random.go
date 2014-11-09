package genetic_algorithm

import (
	"math/rand"
)

var (
	BinaryRandomInitializerInstance = new(BinaryRandomInitializer)
)

type BinaryRandomInitializer struct {
}

func (izer *BinaryRandomInitializer) Init(count, chromSize int) Chromosomes {
	result := make([]ChromosomeInterface, count)

	for chromeInd := 0; chromeInd < count; chromeInd++ {

		genes := make([]bool, chromSize)
		for geneInd := 0; geneInd < chromSize; geneInd++ {
			x := rand.Intn(2)
			genes[geneInd] = x == 1
		}

		result[chromeInd] = NewBinaryChromosome(genes)
	}

	return result
}