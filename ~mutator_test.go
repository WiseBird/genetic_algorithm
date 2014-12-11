package genetic_algorithm

import (
	. "gopkg.in/check.v1"
)

type MutatorSuite struct{}

var _ = Suite(&MutatorSuite{})

func (s *MutatorSuite) TestMutatorGenesBase_Elitism(c *C) {
	trueGenes := BinaryGenes([]bool{true})
	falseGenes := BinaryGenes([]bool{false})

	pop := Chromosomes{
		NewBinaryChromosome(falseGenes),
		NewBinaryChromosome(falseGenes),
		NewBinaryChromosome(falseGenes),
	}

	mutator := NewBinaryMutator(1).WithoutElitism()

	mutator.Mutate(pop)
	c.Assert(pop[0].Genes(), DeepEquals, trueGenes)
	c.Assert(pop[1].Genes(), DeepEquals, trueGenes)
	c.Assert(pop[2].Genes(), DeepEquals, trueGenes)

	mutator.WithElitism(1)
	mutator.Mutate(pop)
	c.Assert(pop[0].Genes(), DeepEquals, trueGenes)
	c.Assert(pop[1].Genes(), DeepEquals, falseGenes)
	c.Assert(pop[2].Genes(), DeepEquals, falseGenes)

	mutator.WithElitism(3)
	mutator.Mutate(pop)
	c.Assert(pop[0].Genes(), DeepEquals, trueGenes)
	c.Assert(pop[1].Genes(), DeepEquals, falseGenes)
	c.Assert(pop[2].Genes(), DeepEquals, falseGenes)
}

func (s *MutatorSuite) TestMutatorInvert_getIntervalLen(c *C) {
	var mutator *InvertMutator

	mutator = NewInvertExactMutator(1, nil, 5)
	c.Assert(mutator.getIntervalLen(10), Equals, 5)

	fromExact := 5
	toExact := 7
	mutator = NewInvertExactIntervalMutator(1, nil, fromExact, toExact)
	res := mutator.getIntervalLen(10)
	if res < fromExact || res > toExact {
		c.Fatalf("Unexpected interval len. Exp: [%d:%d]. Got: [%d]", fromExact, toExact, res)
	}

	mutator = NewInvertPercentageMutator(1, nil, .5)
	c.Assert(mutator.getIntervalLen(10), Equals, 5)

	fromPercentage := .5
	toPercentage := .7
	genesLen := 10
	mutator = NewInvertPercentageIntervalMutator(1, nil, fromPercentage, toPercentage)
	resf := float64(mutator.getIntervalLen(genesLen))
	if resf < fromPercentage*float64(genesLen) || resf > toPercentage*float64(genesLen) {
		c.Fatalf("Unexpected interval len. Exp: [%f:%f]. Got: [%f]", fromPercentage*float64(genesLen), toPercentage*float64(genesLen), resf)
	}
}
func (s *MutatorSuite) TestMutatorInvert_getInterval(c *C) {
	mutator := &InvertMutator{}

	from, to := mutator.getInterval(1, 1)
	c.Assert(from, Equals, 0)
	c.Assert(to, Equals, 1)

	from, to = mutator.getInterval(3, 3)
	c.Assert(from, Equals, 0)
	c.Assert(to, Equals, 3)

	intervalLen := 5
	from, to = mutator.getInterval(10, intervalLen)
	c.Assert(to-from, Equals, intervalLen)
}
func (s *MutatorSuite) TestMutatorInvert_mutate(c *C) {
	beforeGenes := OrderedGenes{1, 2, 3, 4, 5, 6}
	afterGenes := OrderedGenes{1, 2, 5, 4, 3, 6}

	mutator := newInvertMutator(1, NewEmptyOrderedChromosome)
	mutator.mutate(beforeGenes, 2, 5)

	c.Assert(beforeGenes, DeepEquals, afterGenes)
}
