package sort_by_parity

// Problem Statement: https://leetcode.com/problems/sort-array-by-parity/

func SortArrayByParity(nums []int) []int {

	// j := 0
	// for i, n := range nums {
	// 	if j < i {
	// 		j = i
	// 	}
	// 	if n%2 != 0 {
	// 		j++
	// 		for ; j < len(nums); j++ {
	// 			if nums[j]%2 != 0 {
	// 				continue
	// 			}
	// 			nums[i], nums[j] = nums[j], nums[i]
	// 			break
	// 		}
	// 		if j == len(nums) {
	// 			return nums
	// 		}
	// 	}

	// }
	// return nums

	j := 0

	for i, num := range nums {
		if j < i {
			j = i
		}
		if num%2 != 0 {

			for ; j < len(nums); j++ {
				if nums[j]%2 == 0 {
					break
				}
			}
			if j == len(nums) {
				return nums
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
