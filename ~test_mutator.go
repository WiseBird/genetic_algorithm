package genetic_algorithm

import ( 
	. "gopkg.in/check.v1"
)

type MutatorSuite struct{}
var _ = Suite(&MutatorSuite{})

func (s *MutatorSuite) TestMutatorBase_Elitism(c *C) {
	trueGenes := BinaryGenes([]bool{true})
	falseGenes := BinaryGenes([]bool{false})

	pop := Chromosomes {
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