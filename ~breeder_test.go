package genetic_algorithm

import ( 
	. "gopkg.in/check.v1"
	"reflect"
)

type BreederSuite struct{}
var _ = Suite(&BreederSuite{})

func (s *BreederSuite) TestOnePointBreeder_CannotOnEdge_crossover(c *C) {
	parent1Genes := BinaryGenes { true, true }
	parent2Genes := BinaryGenes { false, false }

	expectedPart1Genes := BinaryGenes { true, false }
	expectedPart2Genes := BinaryGenes { false, true }

	parent1 := NewBinaryChromosome(parent1Genes)
	parent2 := NewBinaryChromosome(parent2Genes)

	children := NewOnePointBreeder(NewEmptyBinaryChromosome).Crossover([]ChromosomeInterface {parent1, parent2})
	c1 := children[0]
	c2 := children[1]

	var expC2 BinaryGenes
	if reflect.DeepEqual(c1.Genes(), expectedPart1Genes) {
		expC2 = expectedPart2Genes
	} else if reflect.DeepEqual(c1.Genes(), expectedPart2Genes) {
		expC2 = expectedPart1Genes
	} else {
		c.Fatalf("Unexpected child1 genes: %v", c1.Genes())
	}

	if !reflect.DeepEqual(c2.Genes(), expC2) {
		c.Fatalf("Unexpected child2 genes. Exp: [%v]. Got: [%v]", expC2, c2.Genes())	
	}
}
func (s *BreederSuite) TestOnePointBreeder_CanOnEdge_crossover(c *C) {
	parent1Genes := BinaryGenes { true }
	parent2Genes := BinaryGenes { false }

	parent1 := NewBinaryChromosome(parent1Genes)
	parent2 := NewBinaryChromosome(parent2Genes)

	children := NewOnePointBreeder(NewEmptyBinaryChromosome).CanCrossOnEdge(true).Crossover([]ChromosomeInterface {parent1, parent2})
	c1 := children[0]
	c2 := children[1]

	var expC2 BinaryGenes
	if reflect.DeepEqual(c1.Genes(), parent1Genes) {
		expC2 = parent2Genes
	} else if reflect.DeepEqual(c1.Genes(), parent2Genes) {
		expC2 = parent1Genes
	} else {
		c.Fatalf("Unexpected child1 genes: %v", c1.Genes())
	}

	if !reflect.DeepEqual(c2.Genes(), expC2) {
		c.Fatalf("Unexpected child2 genes. Exp: [%v]. Got: [%v]", expC2, c2.Genes())	
	}
}
func (s *BreederSuite) TestTwoPointBreeder_crossover(c *C) {
	parent1Genes := BinaryGenes { true, true, true }
	parent2Genes := BinaryGenes { false, false, false }

	expectedPart1Genes := BinaryGenes { true, false, true }
	expectedPart2Genes := BinaryGenes { false, true, false }

	parent1 := NewBinaryChromosome(parent1Genes)
	parent2 := NewBinaryChromosome(parent2Genes)

	children := NewTwoPointBreeder(NewEmptyBinaryChromosome).Crossover([]ChromosomeInterface {parent1, parent2})
	c1 := children[0]
	c2 := children[1]

	var expC2 BinaryGenes
	if reflect.DeepEqual(c1.Genes(), expectedPart1Genes) {
		expC2 = expectedPart2Genes
	} else if reflect.DeepEqual(c1.Genes(), expectedPart2Genes) {
		expC2 = expectedPart1Genes
	} else {
		c.Fatalf("Unexpected child1 genes: %v", c1.Genes())
	}

	if !reflect.DeepEqual(c2.Genes(), expC2) {
		c.Fatalf("Unexpected child2 genes. Exp: [%v]. Got: [%v]", expC2, c2.Genes())	
	}
}
