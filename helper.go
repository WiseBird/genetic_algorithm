package genetic_algorithm

import (
	"math"
)

// https://hg.python.org/cpython/file/4480506137ed/Lib/statistics.py#l453

func meanFloat64(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	var sum float64
	for _, val := range values {
		sum += val
	}
	return sum / float64(len(values))
}
// Return sum of square deviations
func ssFloat64(values []float64) float64 {
	mean := meanFloat64(values)

	var sum float64
	var dsum float64
	for _, val := range values {
		dsum += val - mean
		sum += math.Pow(val - mean, 2)
	}

	// Rounding error compensation. Ideally dsum equals zero.
	sum -= math.Pow(dsum, 2) / float64(len(values))

	return sum
}
// Sample variance
func varianceFloat64(values []float64) float64 {
	return ssFloat64(values) / (float64(len(values)) - 1)
}
// Population variance
func pvarianceFloat64(values []float64) float64 {
	return ssFloat64(values) / float64(len(values))
}