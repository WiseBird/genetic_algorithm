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
func meanFloat64Iter(count int, value func(int) float64) float64 {
	if count == 0 {
		return 0
	}

	var sum float64
	for i := 0; i < count; i++ {
		sum += value(i)
	}
	return sum / float64(count)
}

// Expects at least one value in each array
func meanFloat64Arr(values [][]float64) []float64 {
	if len(values) == 0 {
		return []float64{}
	}

	length := 0
	for _, arr := range values {
		if length < len(arr) {
			length = len(arr)
		}
	}

	sum := make([]float64, length)
	for i := 0; i < length; i++ {
		for _, arr := range values {
			if len(arr) > i {
				sum[i] += arr[i]
			} else {
				sum[i] += arr[len(arr) - 1]
			}
		}
	}

	for i := 0; i < length; i++ {
		sum[i] /= float64(len(values))
	}

	return sum
}
// Expects at least one value in each array
func meanFloat64ArrIter(count int, value func(int) []float64) []float64 {
	values := make([][]float64, count)

	for i := 0; i < count; i++ {
		values[i] = value(i)
	}

	return meanFloat64Arr(values)
}

func meanInt64(values []int64) int64 {
	if len(values) == 0 {
		return 0
	}

	var sum int64
	for _, val := range values {
		sum += val
	}
	return sum / int64(len(values))
}
func meanInt64Iter(count int, value func(int) int64) int64 {
	if count == 0 {
		return 0
	}

	var sum int64
	for i := 0; i < count; i++ {
		sum += value(i)
	}
	return sum / int64(count)
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