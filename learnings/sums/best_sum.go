package sums

func BestSum(target int, nums []int) []int {
	combs := make([][]int, target+1)
	combs[0] = []int{}

	for i := 0; i <= target; i++ {
		if combs[i] == nil {
			continue
		}
		for _, num := range nums {
			if i+num > target {
				continue
			}
			existing := combs[i+num]
			newComb := append(combs[i], num)
			if existing == nil || len(newComb) < len(existing) {
				combs[i+num] = newComb
			}

		}
	}
	return combs[target]
}
