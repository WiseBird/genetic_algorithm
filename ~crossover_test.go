package genetic_algorithm

import (
	. "gopkg.in/check.v1"
	"reflect"
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

	// two point crossover tested as chooseTwoPointCrossSection
	// multi point crossover tested as chooseDifferentRandomNumbers
}
func (s *CrossoverSuite) TestMultiPointCrossover_threePoint(c *C) {
	parent1Genes := BinaryGenes{true, true}
	parent2Genes := BinaryGenes{false, false}

	expectedPart1Genes := BinaryGenes{true, false}
	expectedPart2Genes := BinaryGenes{false, true}

	parent1 := NewBinaryChromosome(parent1Genes)
	parent2 := NewBinaryChromosome(parent2Genes)

	children := NewMultiPointCrossover(NewEmptyBinaryChromosome, 3).Crossover([]ChromosomeInterface{parent1, parent2})

	compareTwoBinaryGenesWithoutOrder(c, children[0], children[1], expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestOrderCrossoverVer1_crossover(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 4, 5, 6, 7, 8, 9}
	parent2Genes := OrderedGenes{8, 4, 1, 5, 9, 3, 6, 2, 7}

	expectedPart1Genes := OrderedGenes{2, 7, 3, 4, 5, 6, 8, 1, 9}
	expectedPart2Genes := OrderedGenes{7, 8, 1, 5, 9, 3, 2, 4, 6}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewOrderCrossoverVer1().crossover(parent1, parent2, 2, 6)

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}
func (s *CrossoverSuite) TestOrderCrossoverVer1_crossoverOnEnd(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 4}
	parent2Genes := OrderedGenes{2, 4, 1, 3}

	expectedPart1Genes := OrderedGenes{2, 1, 3, 4}
	expectedPart2Genes := OrderedGenes{1, 2, 4, 3}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewOrderCrossoverVer1().crossover(parent1, parent2, 3, 4)

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}
func (s *CrossoverSuite) TestOrderCrossoverVer2_crossover(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 4, 5, 6, 7, 8, 9}
	parent2Genes := OrderedGenes{8, 4, 1, 5, 9, 3, 6, 2, 7}

	expectedPart1Genes := OrderedGenes{8, 1, 3, 4, 5, 6, 9, 2, 7}
	expectedPart2Genes := OrderedGenes{2, 4, 1, 5, 9, 3, 6, 7, 8}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewOrderCrossoverVer2().crossover(parent1, parent2, 2, 6)

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestPositionBasedCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 4, 5, 6, 7, 8, 9}
	parent2Genes := OrderedGenes{8, 4, 1, 5, 9, 3, 6, 2, 7}

	expectedPart1Genes := OrderedGenes{1, 8, 3, 4, 5, 6, 7, 9, 2}
	expectedPart2Genes := OrderedGenes{8, 2, 1, 4, 5, 3, 6, 7, 9}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewPositionBasedCrossover().crossover(parent1, parent2, []int{0, 2, 5, 6})

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestPartiallyMappedCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 4, 5, 6, 7}
	parent2Genes := OrderedGenes{3, 6, 5, 2, 1, 4, 7}

	expectedPart1Genes := OrderedGenes{1, 6, 3, 4, 5, 2, 7}
	expectedPart2Genes := OrderedGenes{3, 4, 5, 2, 1, 6, 7}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewPartiallyMappedCrossover().crossover(parent1, parent2, 2, 5)

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestRelativeOrderingCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 4, 5, 6, 7, 8}
	parent2Genes := OrderedGenes{5, 6, 2, 7, 3, 1, 8, 4}

	expectedPart1Genes := OrderedGenes{5, 2, 7, 3, 6, 1, 8, 4}
	expectedPart2Genes := OrderedGenes{3, 5, 6, 2, 1, 4, 7, 8}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewRelativeOrderingCrossover(4).crossover(parent1, parent2, []int{1, 2, 5, 7})

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestPrecedencePreservativeCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 4, 5, 6}
	parent2Genes := OrderedGenes{3, 1, 2, 6, 4, 5}

	expectedPart1Genes := OrderedGenes{1, 3, 2, 4, 6, 5}
	expectedPart2Genes := OrderedGenes{3, 1, 2, 6, 4, 5}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewPrecedencePreservativeCrossover().crossover(parent1, parent2, []int{1, 2, 1, 1, 2, 2})

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestCycleCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	parent2Genes := OrderedGenes{8, 4, 7, 3, 6, 2, 5, 1, 9, 0}

	expectedPart1Genes := OrderedGenes{0, 4, 7, 3, 6, 2, 5, 1, 8, 9}
	expectedPart2Genes := OrderedGenes{8, 1, 2, 3, 4, 5, 6, 7, 9, 0}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	children := NewCycleCrossover().Crossover([]ChromosomeInterface{parent1, parent2})

	compareTwoOrderedGenesWithoutOrder(c, children[0], children[1], expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestOrderBasedCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 4, 5, 6, 7, 8, 9}
	parent2Genes := OrderedGenes{5, 4, 6, 3, 1, 9, 2, 7, 8}

	expectedPart1Genes := OrderedGenes{2, 4, 5, 3, 1, 6, 9, 7, 8}
	expectedPart2Genes := OrderedGenes{4, 2, 3, 1, 5, 6, 7, 9, 8}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	c1, c2 := NewOrderBasedCrossover().crossover(parent1, parent2, []int{1, 4, 5, 8})

	compareTwoOrderedGenesWithoutOrder(c, c1, c2, expectedPart1Genes, expectedPart2Genes)
}

func (s *CrossoverSuite) TestEdgeRecombinationCrossover_generateMatrix(c *C) {
	parent1Genes := OrderedGenes{1, 2, 3, 5, 6, 4}
	parent2Genes := OrderedGenes{3, 1, 2, 4, 5, 6}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	matrix := NewEdgeRecombinationCrossover().generateMatrix(parent1, parent2)

	result := map[int]map[int]bool{
		1: map[int]bool{2: true, 3: true, 4: true},
		2: map[int]bool{1: true, 3: true, 4: true},
		3: map[int]bool{1: true, 2: true, 5: true, 6: true},
		4: map[int]bool{1: true, 2: true, 5: true, 6: true},
		5: map[int]bool{3: true, 4: true, 6: true},
		6: map[int]bool{3: true, 4: true, 5: true},
	}

	c.Assert(matrix, DeepEquals, result)
}
func (s *CrossoverSuite) TestEdgeRecombinationCrossover_fillChild(c *C) {
	matrix := map[int]map[int]bool{
		1: map[int]bool{4: true, 5: true, 6: true},
		2: map[int]bool{3: true, 4: true, 6: true},
		3: map[int]bool{2: true, 5: true},
		4: map[int]bool{1: true, 2: true},
		5: map[int]bool{1: true, 3: true, 6: true},
		6: map[int]bool{1: true, 2: true, 5: true},
	}

	c1genes := make(OrderedGenes, 6)

	NewEdgeRecombinationCrossover().fillChild(c1genes, OrderedGenes{1}, OrderedGenes{1}, matrix)

	result1 := OrderedGenes{1, 4, 2, 3, 5, 6}
	result2 := OrderedGenes{1, 4, 2, 6, 5, 3}

	if !reflect.DeepEqual(c1genes, result1) && !reflect.DeepEqual(c1genes, result2) {
		c.Fatalf("Unexpected genes: %v", c1genes)
	}
}
func (s *CrossoverSuite) TestEdgeRecombinationCrossover_crossover(c *C) {
	parent1Genes := OrderedGenes{1, 0, 7, 4, 3, 2, 6, 5}
	parent2Genes := OrderedGenes{5, 3, 4, 2, 6, 7, 1, 0}

	parent1 := NewOrderedChromosome(parent1Genes)
	parent2 := NewOrderedChromosome(parent2Genes)

	for j := 0; j < 10; j++ {
		children := NewEdgeRecombinationCrossover().Crossover([]ChromosomeInterface{parent1, parent2})
		childGenes := children[0].Genes().(OrderedGenes)
		for i := 0; i < len(childGenes); i++ {
			if childGenes[i] == -1 {
				c.Fatalf("Unexpected genes: %v", childGenes)
			}
		}
	}
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
