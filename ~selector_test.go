package genetic_algorithm

import (
	"code.google.com/p/gomock/gomock"
	. "gopkg.in/check.v1"
)

type SelectorSuite struct{}

var _ = Suite(&SelectorSuite{})

func (s *SelectorSuite) Test_SelectorBase_SelectMany_Should_Panic_OnNegativeCount(c *C) {
	selector := NewRouletteWheelCostWeightingSelector()

	c.Assert(func() { selector.SelectMany(-10) }, PanicMatches, `Count.*`)
}

func (s *SelectorSuite) Test_SelectorBase_SelectMany_Should_Panic_WhenUniqueCountBiggerPop(c *C) {
	pop := make(Chromosomes, 2)

	ctrl := gomock.NewController(c)
	defer ctrl.Finish()

	vm := NewMockSelectorBaseVirtualMInterface(ctrl)

	selector := NewSelectorBase(vm).SelectManyAreUnique(true)

	selector.Prepare(pop)
	c.Assert(func() { selector.SelectMany(3) }, PanicMatches, `.*unique.*`)
}

func (s *SelectorSuite) Test_SelectorTournamentProbabilities(c *C) {
	selector := NewTournamentSelector(0.5, 1)
	c.Assert(0.5, Within, 0.0001, selector.ithProbability(0))
	c.Assert(0.25, Within, 0.0001, selector.ithProbability(1))
	c.Assert(0.125, Within, 0.0001, selector.ithProbability(2))

	selector = NewTournamentSelector(0.8, 1)
	c.Assert(0.8, Within, 0.0001, selector.ithProbability(0))
	c.Assert(0.16, Within, 0.0001, selector.ithProbability(1))
	c.Assert(0.032, Within, 0.0001, selector.ithProbability(2))
}

func (s *SelectorSuite) Test_SelectorSimpleTournament_ReturnsTheOnePossibleValue(c *C) {
	pop := Chromosomes{NewEmptyBinaryChromosome(1)}

	selector := NewSimpleTournamentSelector(3)
	selector.Prepare(pop)

	c.Assert(selector.SelectInd(), Equals, 0)
}

func (s *SelectorSuite) Test_SelectorTournament_ReturnsTheOnePossibleValue(c *C) {
	pop := Chromosomes{NewEmptyBinaryChromosome(1)}

	selector := NewTournamentSelector(0.5, 3)
	selector.Prepare(pop)

	c.Assert(selector.SelectInd(), Equals, 0)
}

func (s *SelectorSuite) Test_SelectorWheelRank_Weighting(c *C) {
	selector := NewRouletteWheelRankWeightingSelector()

	selector.Prepare(make(Chromosomes, 1))
	c.Assert(selector.weights, DeepEquals, []float64{1})

	selector.Prepare(make(Chromosomes, 4))
	c.Assert(selector.weights, DeepEquals, []float64{0.4, 0.3, 0.2, 0.1})
}
