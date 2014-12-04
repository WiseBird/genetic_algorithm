package genetic_algorithm

import ( 
	. "gopkg.in/check.v1"
	"reflect"
	"sort"
)

type CrossoverSuite struct{}
var _ = Suite(&CrossoverSuite{})

func (s *CrossoverSuite) TestMultiPointCrossover_chooseCrossPoints(c *C) {
	onePointCrossover := NewOnePointCrossover(NewEmptyBinaryChromosome)
	ps := onePointCrossover.chooseCrossPoints(2)
	if ps[0] != 1 {
		c.Fatalf("One point crossover didn't choose middle point. choose %d", ps[0])
	}

	onePointCrossover.CanProduceCopiesOfParents(true)
	ps = onePointCrossover.chooseCrossPoints(1)
	if ps[0] != 0 && ps[0] != 1 {
		c.Fatalf("One point crossover choose wrong point %d", ps[0])
	}

	// two point crossover uses chooseTwoPointCrossSection

	multiPonintCrossover := NewMultiPointCrossover(NewEmptyBinaryChromosome, 3)
	ps = multiPonintCrossover.chooseCrossPoints(2)
	sort.Sort(sort.IntSlice(ps))
	c.Assert(ps, DeepEquals, []int{0,1,2})
}

func (s *CrossoverSuite) TestOrderCrossoverVer1_crossover(c *C) {
	parent1Genes := OrderedGenes { 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	parent2Genes := OrderedGenes { 8, 4, 1, 5, 9, 3, 6, 2, 7 }

	expectedPart1Genes := OrderedGenes { 2, 7, 3, 4, 5, 6, 8, 1, 9 }
	expectedPart2Genes := OrderedGenes { 7, 8, 1, 5, 9, 3, 2, 4, 6 }

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewOrderCrossoverVer1().crossover(parent1, parent2, 2, 6)

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}
func (s *CrossoverSuite) TestOrderCrossoverVer1_crossoverOnEnd(c *C) {
	parent1Genes := OrderedGenes { 1, 2, 3, 4 }
	parent2Genes := OrderedGenes { 2, 4, 1, 3 }

	expectedPart1Genes := OrderedGenes { 2, 1, 3, 4 }
	expectedPart2Genes := OrderedGenes { 1, 2, 4, 3 }

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewOrderCrossoverVer1().crossover(parent1, parent2, 3, 4)

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}
func (s *CrossoverSuite) TestOrderCrossoverVer2_crossover(c *C) {
	parent1Genes := OrderedGenes { 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	parent2Genes := OrderedGenes { 8, 4, 1, 5, 9, 3, 6, 2, 7 }

	expectedPart1Genes := OrderedGenes { 8, 1, 3, 4, 5, 6, 9, 2, 7 }
	expectedPart2Genes := OrderedGenes { 2, 4, 1, 5, 9, 3, 6, 7, 8 }

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewOrderCrossoverVer2().crossover(parent1, parent2, 2, 6)

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestPositionCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes { 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	parent2Genes := OrderedGenes { 8, 4, 1, 5, 9, 3, 6, 2, 7 }

	expectedPart1Genes := OrderedGenes { 1, 8, 3, 4, 5, 6, 7, 9, 2 }
	expectedPart2Genes := OrderedGenes { 8, 2, 1, 4, 5, 3, 6, 7, 9 }

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewPositionCrossover().crossover(parent1, parent2, []int { 0, 2, 5, 6 })

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestPartiallyMappedCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes { 1, 2, 3, 4, 5, 6, 7 }
	parent2Genes := OrderedGenes { 3, 6, 5, 2, 1, 4, 7 }

	expectedPart1Genes := OrderedGenes { 1, 6, 3, 4, 5, 2, 7 }
	expectedPart2Genes := OrderedGenes { 3, 4, 5, 2, 1, 6, 7 }

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewPartiallyMappedCrossover().crossover(parent1, parent2, 2, 5)

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func compareTwoBinaryGenesWithoutOrder(c *C, c1, c2 ChromosomeInterface, ec1, ec2 BinaryGenes) {
	var expC2 BinaryGenes
	if reflect.DeepEqual(c1.Genes(), ec1) {
		expC2 = ec2
	} else if reflect.DeepEqual(c1.Genes(), ec2) {
		expC2 = ec1
	} else {
		c.Fatalf("Unexpected child1 genes: %v", c1.Genes())
	}

	if !reflect.DeepEqual(c2.Genes(), expC2) {
		c.Fatalf("Unexpected child2 genes. Exp: [%v]. Got: [%v]", expC2, c2.Genes())	
	}
}
func compareTwoOrderedGenesWithoutOrder(c *C, c1, c2 ChromosomeInterface, ec1, ec2 OrderedGenes) {
	var expC2 OrderedGenes
	if reflect.DeepEqual(c1.Genes(), ec1) {
		expC2 = ec2
	} else if reflect.DeepEqual(c1.Genes(), ec2) {
		expC2 = ec1
	} else {
		c.Fatalf("Unexpected child1 genes: %v", c1.Genes())
	}

	if !reflect.DeepEqual(c2.Genes(), expC2) {
		c.Fatalf("Unexpected child2 genes. Exp: [%v]. Got: [%v]", expC2, c2.Genes())	
	}
}