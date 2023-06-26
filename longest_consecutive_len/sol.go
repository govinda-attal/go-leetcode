package longest_consecutive_len

func LongestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for _, n := range nums {
		set[n] = struct{}{}
	}

	maxSqLen := 0

	for n := range set {
		if _, ok := set[n-1]; ok {
			continue
		}
		sqLen := 1
		for {
			if _, ok := set[n+sqLen]; ok {
				sqLen++
				continue
			}
			maxSqLen = max(sqLen, maxSqLen)
			break
		}
	}
	return maxSqLen
}
