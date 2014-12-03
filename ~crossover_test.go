package genetic_algorithm

import ( 
	. "gopkg.in/check.v1"
	"reflect"
)

type CrossoverSuite struct{}
var _ = Suite(&CrossoverSuite{})

func (s *CrossoverSuite) TestOnePointCrossover_CannotOnEdge_crossover(c *C) {
	parent1Genes := BinaryGenes { true, true }
	parent2Genes := BinaryGenes { false, false }

	expectedPart1Genes := BinaryGenes { true, false }
	expectedPart2Genes := BinaryGenes { false, true }

	parent1 := NewBinaryChromosome(parent1Genes)
	parent2 := NewBinaryChromosome(parent2Genes)

	children := NewOnePointCrossover(NewEmptyBinaryChromosome).Crossover([]ChromosomeInterface {parent1, parent2})

	compareTwoBinaryGenesWithoutOrder(c, children[0], children[1], expectedPart1Genes, expectedPart2Genes)
}
func (s *CrossoverSuite) TestOnePointCrossover_CanOnEdge_crossover(c *C) {
	parent1Genes := BinaryGenes { true }
	parent2Genes := BinaryGenes { false }

	parent1 := NewBinaryChromosome(parent1Genes)
	parent2 := NewBinaryChromosome(parent2Genes)

	children := NewOnePointCrossover(NewEmptyBinaryChromosome).CanCrossOnEdge(true).Crossover([]ChromosomeInterface {parent1, parent2})

	compareTwoBinaryGenesWithoutOrder(c, children[0], children[1], parent1Genes, parent2Genes)
}
func (s *CrossoverSuite) TestTwoPointCrossover_crossover(c *C) {
	parent1Genes := BinaryGenes { true, true, true }
	parent2Genes := BinaryGenes { false, false, false }

	expectedPart1Genes := BinaryGenes { true, false, true }
	expectedPart2Genes := BinaryGenes { false, true, false }

	parent1 := NewBinaryChromosome(parent1Genes)
	parent2 := NewBinaryChromosome(parent2Genes)

	children := NewTwoPointCrossover(NewEmptyBinaryChromosome).Crossover([]ChromosomeInterface {parent1, parent2})

	compareTwoBinaryGenesWithoutOrder(c, children[0], children[1], expectedPart1Genes, expectedPart2Genes)
}


func (s *CrossoverSuite) TestOrderCrossover_firstCrossPoint(c *C) {
	crossover := newOrderCrossover(nil)

	for i := 0; i < 10; i++ {
		p := crossover.chooseFirstCrossPoint(1) 
		if p != 0 {
			c.Fatalf("Choose inappropriate point: %v", p)
		}
	}
}
func (s *CrossoverSuite) TestOrderCrossover_secondCrossPoint_notEqualsFirst(c *C) {
	crossover := newOrderCrossover(nil).CanProduceCopiesOfParents(true)

	for i := 0; i < 10; i++ {
		p := crossover.chooseSecondCrossPoint(1, 0) 
		if p != 1 {
			c.Fatalf("Choose inappropriate point: %v", p)
		}
	}
}
func (s *CrossoverSuite) TestOrderCrossover_secondCrossPoint_cantCopiesOfParent(c *C) {
	crossover := newOrderCrossover(nil).CanProduceCopiesOfParents(false)

	for i := 0; i < 10; i++ {
		p := crossover.chooseSecondCrossPoint(2, 0) 
		if p != 1 {
			c.Fatalf("Choose inappropriate point: %v", p)
		}
	}
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