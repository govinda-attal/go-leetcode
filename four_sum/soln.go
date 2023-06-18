package four_sum

import (
	"fmt"
	"sort"
)

// Problem Statement: https://leetcode.com/problems/4sum/

func FourSum(nums []int, target int) (rs [][]int) {

	sort.Ints(nums)

	if len(nums) < 4 {
		return
	}

	unique := map[string]struct{}{}

	total := len(nums)

	for i := 0; i < total-3; i++ {
		jjs := map[int]struct{}{}
		for j := i + 1; j < total-2; j++ {
			if _, ok := jjs[nums[j]]; ok {
				continue
			} else {
				jjs[nums[j]] = struct{}{}
			}

			exp := target - (nums[i] + nums[j])

			l := j + 1
			r := total - 1

			ss := twoSum(nums, l, r, exp)
			if ss == nil {
				continue
			}
			for _, css := range ss {
				comb := []int{nums[i], nums[j], css[0], css[1]}
				combKey := fmt.Sprintf("%v", comb)
				if _, ok := unique[combKey]; ok {
					continue
				}
				rs = append(rs, comb)
				unique[combKey] = struct{}{}
			}

		}
	}

	return rs
}

func twoSum(nums []int, l, r, target int) (ts [][]int) {
	type index = int
	compl := map[int]index{}

	for i := l; i <= r; i++ {
		bal := target - nums[i]
		if _, ok := compl[bal]; ok {
			ts = append(ts, []int{bal, nums[i]})
		}
		compl[nums[i]] = i
	}
	return ts
}
