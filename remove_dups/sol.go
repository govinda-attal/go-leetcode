package remove_dups

func RemoveDuplicates(nums []int) int {
	ix := 0
	unique := len(nums)
	set := map[int]struct{}{}
	for i := range nums {
		if _, ok := set[nums[i]]; ok {
			unique--
			continue
		}
		nums[ix] = nums[i]
		ix++
		set[nums[i]] = struct{}{}
	}
	return unique
}
