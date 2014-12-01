package genetic_algorithm

import (
	"math/rand"
	log "github.com/cihub/seelog"
	"fmt"
	"sort"
)

type MultiPointCrossover struct {
	crossPointsCount int
	chromConstr EmptyChromosomeConstructor
	canCrossOnEdge bool
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
// If true cross point can be selected on edge of chromosome
// For example in case of one point crossover it will produce two copies of parents
func (crossover *MultiPointCrossover) CanCrossOnEdge(val bool) *MultiPointCrossover {
	crossover.canCrossOnEdge = val
	return crossover
}
func (crossover *MultiPointCrossover) Crossover(parents Chromosomes) Chromosomes {
	if len(parents) != crossover.ParentsCount() {
		panic("Incorrect parents count")
	}

	p1 := parents[0]
	p2 := parents[1]

	if p1.Genes().Len() != p2.Genes().Len() {
		panic("Crossover do not support different chromosome size")
	}

	possibleCrossPoints := p1.Genes().Len() + 1
	if !crossover.canCrossOnEdge {
		possibleCrossPoints -= 2
	}

	if possibleCrossPoints < crossover.crossPointsCount {
		panic(fmt.Sprintf("Crossover can't split gene on %d parts", crossover.crossPointsCount + 1))
	}

	firstPossibleCrossPoint := 0
	if !crossover.canCrossOnEdge {
		firstPossibleCrossPoint = 1
	}

	crossPointsMap := make(map[int]bool, crossover.crossPointsCount)
	crossPointsList := make([]int, 0, crossover.crossPointsCount)
	for i := 0; i < crossover.crossPointsCount; i++ {
		for ;; {
			crossPoint := rand.Intn(possibleCrossPoints) + firstPossibleCrossPoint;
			if !crossPointsMap[crossPoint] {
				crossPointsMap[crossPoint] = true
				crossPointsList = append(crossPointsList, crossPoint)
				break
			}
		}
	}

	sort.Sort(sort.IntSlice(crossPointsList))

	log.Tracef("Cross on %v\n", crossPointsList)

	c1, c2 := crossover.crossover(p1, p2, crossPointsList)

	return Chromosomes{c1, c2}
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