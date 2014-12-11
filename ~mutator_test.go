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

func (s *MutatorSuite) TestMutatorIntervalBase_getIntervalLen(c *C) {
	var mutator *MutatorIntervalBase

	mutator = NewMutatorIntervalBase(nil, 1, nil).ExactInterval(5, 5)
	c.Assert(mutator.getIntervalLen(10), Equals, 5)

	fromExact := 5
	toExact := 7
	mutator = NewMutatorIntervalBase(nil, 1, nil).ExactInterval(fromExact, toExact)
	res := mutator.getIntervalLen(10)
	if res < fromExact || res > toExact {
		c.Fatalf("Unexpected interval len. Exp: [%d:%d]. Got: [%d]", fromExact, toExact, res)
	}

	mutator = NewMutatorIntervalBase(nil, 1, nil).PercentageInterval(.5, .5)
	c.Assert(mutator.getIntervalLen(10), Equals, 5)

	fromPercentage := .5
	toPercentage := .7
	genesLen := 10
	mutator = NewMutatorIntervalBase(nil, 1, nil).PercentageInterval(fromPercentage, toPercentage)
	resf := float64(mutator.getIntervalLen(genesLen))
	if resf < fromPercentage*float64(genesLen) || resf > toPercentage*float64(genesLen) {
		c.Fatalf("Unexpected interval len. Exp: [%f:%f]. Got: [%f]", fromPercentage*float64(genesLen), toPercentage*float64(genesLen), resf)
	}
}
func (s *MutatorSuite) TestMutatorIntervalBase_getInterval(c *C) {
	mutator := &MutatorIntervalBase{}

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

func (s *MutatorSuite) TestInvertMutator_mutate(c *C) {
	beforeGenes := OrderedGenes{1, 2, 3, 4, 5, 6}
	afterGenes := OrderedGenes{1, 2, 5, 4, 3, 6}

	mutator := NewInvertMutator(1, NewEmptyOrderedChromosome)
	mutator.MutateGenes(beforeGenes, 2, 5)

	c.Assert(beforeGenes, DeepEquals, afterGenes)
}

func (s *MutatorSuite) TestDisplacementMutator_insert(c *C) {
	beforeGenes := OrderedGenes{1, 2, 3, 4, 5, 6}
	genes := make(OrderedGenes, len(beforeGenes))
	mutator := NewDisplacementMutator(1, NewEmptyOrderedChromosome)

	genes.Copy(beforeGenes, 0, 0, len(beforeGenes))
	mutator.insert(genes, 2, 5, 0)
	c.Assert(genes, DeepEquals, OrderedGenes{3, 4, 5, 1, 2, 6})

	genes.Copy(beforeGenes, 0, 0, len(beforeGenes))
	mutator.insert(genes, 2, 5, 1)
	c.Assert(genes, DeepEquals, OrderedGenes{1, 3, 4, 5, 2, 6})

	genes.Copy(beforeGenes, 0, 0, len(beforeGenes))
	mutator.insert(genes, 2, 5, 6)
	c.Assert(genes, DeepEquals, OrderedGenes{1, 2, 6, 3, 4, 5})
}
func (s *MutatorSuite) TestDisplacementMutator_chooseInsertPoint(c *C) {
	mutator := NewDisplacementMutator(1, NewEmptyOrderedChromosome)

	point := mutator.chooseInsertPoint(2, 0, 1)
	c.Assert(point, Equals, 2)

	point = mutator.chooseInsertPoint(2, 1, 2)
	c.Assert(point, Equals, 0)

	point = mutator.chooseInsertPoint(6, 2, 5)
	if point != 0 && point != 1 && point != 6 {
		c.Fatalf("Unexpected insert point. Exp: [0|1|6]. Got: [%d]", point)
	}
}
