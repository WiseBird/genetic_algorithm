package genetic_algorithm

import (
	"math/rand"
	log "github.com/cihub/seelog"
)

type OnePointBinaryBreeder struct {
}
func NewOnePointBinaryBreeder() *OnePointBinaryBreeder {
	var breeder *OnePointBinaryBreeder

	return breeder
}

func (breeder *OnePointBinaryBreeder) ParentsCount() int {
	return 2
}

func (breeder *OnePointBinaryBreeder) Crossover(parents Chromosomes) Chromosomes {
	if len(parents) != breeder.ParentsCount() {
		panic("Incorrect parents count")
	}

	bp1 := parents[0].(*BinaryChromosome)
	bp2 := parents[1].(*BinaryChromosome)

	if bp1.Size() != bp2.Size() {
		panic("Breeder do not support different chromosome size")
	}

	bitSize := bp1.Size()

	bitToCross := rand.Intn(bitSize - 1) + 1;
	log.Debugf("Cross on %v\n", bitToCross)

	c1Genes := make([]bool, bitSize)
	copy(c1Genes, bp1.Genes[:bitToCross])
	copy(c1Genes[bitToCross:], bp2.Genes[bitToCross:])

	c1 := NewBinaryChromosome(c1Genes)

	c2Genes := make([]bool, bitSize)
	copy(c2Genes, bp2.Genes[:bitToCross])
	copy(c2Genes[bitToCross:], bp1.Genes[bitToCross:])

	c2 := NewBinaryChromosome(c2Genes)

	return Chromosomes{c1, c2}
}