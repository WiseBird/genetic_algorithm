package genetic_algorithm

import (
	"math"
	. "gopkg.in/check.v1"
)

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