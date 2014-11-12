package genetic_algorithm

import ( 
	. "gopkg.in/check.v1"
	"code.google.com/p/gomock/gomock"
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