package statistics

import (
	"slices"
)

// for this we use a slice floats because we want to calculate the mean, median, and mode of a set of numbers. The slice allows us to store a variable number of float64 values, which can then be processed to compute the desired statistical measures.

func Mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func Median(data []float64) float64 {
	n := len(data)

	if n == 0 {
		return 0
	}

	// clone the slice
	sortedNumbers := slices.Clone(data)
	// sort the cloned slice
	slices.Sort(sortedNumbers)

	// apply median rules based on odd/even  count
	if n%2 != 0 {
	
		// odd count, return the middle value
		return sortedNumbers[n/2]
	}
	// even count, return the average of the two middle values
	mid1 := sortedNumbers[n/2-1]
	mid2 := sortedNumbers[n/2]
	return (mid1 + mid2) / 2



}