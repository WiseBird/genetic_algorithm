package genetic_algorithm

import ( 
	. "gopkg.in/check.v1"
	"time"
)

type StatisticsSuite struct{}
var _ = Suite(&StatisticsSuite{})

func (s *StatisticsSuite) TestStatisticsDefault_TrackTime_When_Running(c *C) {
	stat := NewStatisticsDefault(NewStatisticsDefaultOptions()).(*StatisticsDefault)

	stat.Start()
	time.Sleep(1e7)

	d1 := stat.Duration()
	time.Sleep(1e7)
	d2 := stat.Duration()

	c.Assert(d2 > d1, Equals, true)
}
func (s *StatisticsSuite) TestStatisticsDefault_StopTrackTime_After_Stop(c *C) {
	stat := NewStatisticsDefault(NewStatisticsDefaultOptions()).(*StatisticsDefault)

	stat.Start()
	time.Sleep(1e7)
	stat.End()

	d1 := stat.Duration()
	time.Sleep(1e7)
	d2 := stat.Duration()

	c.Assert(d2 == d1, Equals, true)
}