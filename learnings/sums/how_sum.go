package sums

func HowSum(target int, nums []int) []int {
	combinations := make([][]int, target+1)
	combinations[0] = []int{}

	for i := 0; i <= target; i++ {
		if combinations[i] == nil {
			continue
		}
		for _, n := range nums {
			if i+n > target {
				continue
			}
			combinations[i+n] = append(combinations[i], n)
		}
	}
	return combinations[target]
}
