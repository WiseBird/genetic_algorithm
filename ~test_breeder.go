package genetic_algorithm

import ( 
	. "gopkg.in/check.v1"
	"reflect"
)

type BreederSuite struct{}
var _ = Suite(&BreederSuite{})

func (s *BreederSuite) TestOnePointBreeder_crossover(c *C) {
	pointToCross := 2
	parent1Genes := BinaryGenes { true, true, true, true }
	parent2Genes := BinaryGenes { false, false, false, false }

	expectedPart1Genes := BinaryGenes { true, true, false, false }
	expectedPart2Genes := BinaryGenes { false, false, true, true }

	parent1 := NewBinaryChromosome(parent1Genes)
	parent2 := NewBinaryChromosome(parent2Genes)

	c1, c2 := NewOnePointBreeder(NewEmptyBinaryChromosome).crossover(parent1, parent2, pointToCross)

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