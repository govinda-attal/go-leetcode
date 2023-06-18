package centered_average

import "sort"

// Return the "centered" average of an array of ints,
// which we'll say is the mean average of the values,
// except ignoring the largest and smallest values in the array.
// If there are multiple copies of the smallest value, ignore just one copy, and likewise for the largest value.
// Use int division to produce the final average. You may assume that the array is length 3 or more.

func CenteredAverage(items []int) int {
	if len(items) <= 2 {
		return 0
	}
	sort.Ints(items)

	sum := 0
	size := len(items)
	for i := 1; i < size-1; i++ {
		sum += items[i]
	}
	return sum / (size - 1)
}
