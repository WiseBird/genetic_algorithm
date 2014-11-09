package genetic_algorithm

import (
	"math/rand"
	log "github.com/cihub/seelog"
)

type OnePointBreeder struct {
	chromConstr EmptyChromosomeConstructor
}
func NewOnePointBreeder(chromConstr EmptyChromosomeConstructor) *OnePointBreeder {
	breeder := new(OnePointBreeder)

	breeder.chromConstr = chromConstr

	return breeder
}

func (breeder *OnePointBreeder) ParentsCount() int {
	return 2
}

func (breeder *OnePointBreeder) Crossover(parents Chromosomes) Chromosomes {
	if len(parents) != breeder.ParentsCount() {
		panic("Incorrect parents count")
	}

	p1 := parents[0]
	p2 := parents[1]

	if p1.Genes().Len() != p2.Genes().Len() {
		panic("Breeder do not support different chromosome size")
	}

	genesLen := p1.Genes().Len()

	bitToCross := rand.Intn(genesLen - 1) + 1;
	log.Debugf("Cross on %v\n", bitToCross)

	c1, c2 := breeder.crossover(p1, p2, bitToCross)

	return Chromosomes{c1, c2}
}
func (breeder *OnePointBreeder) crossover(p1, p2 ChromosomeInterface, bitToCross int) (c1, c2 ChromosomeInterface) {
	genesLen := p1.Genes().Len()

	c1 = breeder.chromConstr(genesLen)
	c1.Genes().Copy(p1.Genes(), 0, 0, bitToCross)
	c1.Genes().Copy(p2.Genes(), bitToCross, bitToCross, genesLen)

	c2 = breeder.chromConstr(genesLen)
	c2.Genes().Copy(p2.Genes(), 0, 0, bitToCross)
	c2.Genes().Copy(p1.Genes(), bitToCross, bitToCross, genesLen)

	return
}