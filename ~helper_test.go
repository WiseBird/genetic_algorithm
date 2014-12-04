package genetic_algorithm

import (
	. "gopkg.in/check.v1"
	"math"
)

type HelperSuite struct{}

var _ = Suite(&HelperSuite{})

func (s *HelperSuite) Test_MeanInt64(c *C) {
	values := []int64{2, 3, 4}
	result := int64(3)

	c.Assert(meanInt64(values), Equals, result)
}
func (s *HelperSuite) Test_MeanFloat64Arr(c *C) {
	values := [][]float64{
		[]float64{2},
		[]float64{8, 6, 4},
		[]float64{20, 19, 15, 12, 9},
	}
	result := []float64{10, 9, 7, 6, 5}

	c.Assert(meanFloat64Arr(values), DeepEquals, result)
	c.Assert(meanFloat64ArrIter(len(values), func(i int) []float64 {
		return values[i]
	}), DeepEquals, result)
}
func (s *HelperSuite) Test_MeanFloat64Arr_BigNumbers(c *C) {
	values := [][]float64{
		[]float64{2e30, 1e8},
		[]float64{1e10, 1e8},
	}
	result := []float64{1e30, 1e8}

	c.Assert(meanFloat64Arr(values), DeepEquals, result)
}

func (s *HelperSuite) Test_chooseTwoPointCrossSection_firstCrossPoint(c *C) {
	p, _ := chooseTwoPointCrossSection(1, true)
	if p != 0 {
		c.Fatalf("Choose inappropriate point: %v", p)
	}
}
func (s *HelperSuite) Test_chooseTwoPointCrossSection_secondCrossPoint_notEqualsFirst(c *C) {
	_, p := chooseTwoPointCrossSection(1, true)
	if p != 1 {
		c.Fatalf("Choose inappropriate point: %v", p)
	}
}
func (s *HelperSuite) Test_chooseTwoPointCrossSection_secondCrossPoint_cantCopiesOfParent(c *C) {
	for i := 0; i < 10; i++ {
		p1, p2 := chooseTwoPointCrossSection(2, false)
		if ((p1 != 0) || (p2 != 1)) && ((p1 != 1) || (p2 != 2)) {
			c.Fatalf("Choose inappropriate points: %d:%d", p2, p2)
		}
	}
}

//Within Delta Custom Checker
type withinChecker struct {
	*CheckerInfo
}

var Within Checker = &withinChecker{
	&CheckerInfo{Name: "Within", Params: []string{"obtained", "delta", "expected"}},
}

func (c *withinChecker) Check(params []interface{}, names []string) (result bool, error string) {
	obtained, ok := params[0].(float64)
	if !ok {
		return false, "obtained must be a float64"
	}
	delta, ok := params[1].(float64)
	if !ok {
		return false, "delta must be a float64"
	}
	expected, ok := params[2].(float64)
	if !ok {
		return false, "expected must be a float64"
	}
	return math.Abs(obtained-expected) <= delta, ""
}
