package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math/rand"
	"sort"
)

type MultiPointCrossover struct {
	crossPointsCount          int
	chromConstr               EmptyChromosomeConstructor
	canProduceCopiesOfParents bool
}

func NewMultiPointCrossover(chromConstr EmptyChromosomeConstructor, crossPointsCount int) *MultiPointCrossover {
	if crossPointsCount <= 0 {
		panic("crossPointsCount must be positive")
	}

	crossover := new(MultiPointCrossover)

	crossover.chromConstr = chromConstr
	crossover.crossPointsCount = crossPointsCount

	return crossover
}
func NewOnePointCrossover(chromConstr EmptyChromosomeConstructor) *MultiPointCrossover {
	return NewMultiPointCrossover(chromConstr, 1)
}
func NewTwoPointCrossover(chromConstr EmptyChromosomeConstructor) *MultiPointCrossover {
	return NewMultiPointCrossover(chromConstr, 2)
}

func (crossover *MultiPointCrossover) ParentsCount() int {
	return 2
}
func (crossover *MultiPointCrossover) CanProduceCopiesOfParents(val bool) *MultiPointCrossover {
	crossover.canProduceCopiesOfParents = val
	return crossover
}
func (crossover *MultiPointCrossover) Crossover(parents Chromosomes) Chromosomes {
	if len(parents) != crossover.ParentsCount() {
		panic("Incorrect parents count")
	}

	p1 := parents[0]
	p2 := parents[1]

	genesLen := p1.Genes().Len()

	if genesLen != p2.Genes().Len() {
		panic("Crossover do not support different chromosome size")
	}

	crossover.checkGenesLen(genesLen)

	crossPointsList := crossover.chooseCrossPoints(genesLen)
	sort.Sort(sort.IntSlice(crossPointsList))

	log.Tracef("Cross on %v\n", crossPointsList)

	c1, c2 := crossover.crossover(p1, p2, crossPointsList)

	return Chromosomes{c1, c2}
}
func (crossover *MultiPointCrossover) checkGenesLen(genesLen int) {
	possibleCrossPoints := genesLen + 1
	if !crossover.canProduceCopiesOfParents && crossover.crossPointsCount <= 2 {
		if crossover.crossPointsCount == 1 {
			possibleCrossPoints -= 2
		} else {
			possibleCrossPoints--
		}
	}

	if possibleCrossPoints < crossover.crossPointsCount {
		panic("Chromosome too short")
	}
}
func (crossover *MultiPointCrossover) chooseCrossPoints(genesLen int) []int {
	if crossover.crossPointsCount == 1 {
		if crossover.canProduceCopiesOfParents {
			return []int{rand.Intn(genesLen + 1)}
		} else {
			return []int{rand.Intn(genesLen-1) + 1}
		}
	} else if crossover.crossPointsCount == 2 {
		p1, p2 := chooseTwoPointCrossSection(genesLen, crossover.canProduceCopiesOfParents)
		return []int{p1, p2}
	}
	return chooseDifferentRandomNumbers(crossover.crossPointsCount, genesLen+1)
}
func (crossover *MultiPointCrossover) crossover(p1, p2 ChromosomeInterface, crossPoints []int) (c1, c2 ChromosomeInterface) {
	p1genes := p1.Genes()
	p2genes := p2.Genes()

	genesLen := p1.Genes().Len()

	c1 = crossover.chromConstr(genesLen)
	c1genes, ok := c1.Genes().(CopyableGenesInterface)
	if !ok {
		panic("Chromosome's genes does not implement CopyableGenesInterface")
	}

	c2 = crossover.chromConstr(genesLen)
	c2genes, ok := c2.Genes().(CopyableGenesInterface)
	if !ok {
		panic("Chromosome's genes does not implement CopyableGenesInterface")
	}

	crossPoint := 0
	for i := 0; i < len(crossPoints); i++ {
		crossPoint = crossPoints[i]
		start := 0
		if i > 0 {
			start = crossPoints[i-1]
		}

		c1genes.Copy(p1genes, start, start, crossPoint)
		c2genes.Copy(p2genes, start, start, crossPoint)
		c1genes, c2genes = c2genes, c1genes
	}

	c1genes.Copy(p1genes, crossPoint, crossPoint, genesLen)
	c2genes.Copy(p2genes, crossPoint, crossPoint, genesLen)

	return
}
