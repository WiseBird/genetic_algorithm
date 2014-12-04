package genetic_algorithm

import (
	. "gopkg.in/check.v1"
)

type WeederSuite struct{}

var _ = Suite(&WeederSuite{})

func (s *WeederSuite) TestSimpleWeeder(c *C) {
	popLen := 100

	pop := make(Chromosomes, popLen)

	weeder := NewSimpleWeeder(50)
	c.Assert(len(weeder.Weed(pop)), Equals, 50)

	weeder = NewSimpleWeeder(80)
	c.Assert(len(weeder.Weed(pop)), Equals, 20)

	weeder = NewSimpleWeeder(25)
	c.Assert(len(weeder.Weed(pop)), Equals, 75)
}
