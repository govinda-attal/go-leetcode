package three_sum

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) (rs [][]int) {
	sort.Ints(nums)
	uniques := map[string]struct{}{}
	targetSeen := map[int]struct{}{}
	for i := 0; i < len(nums)-2; i++ {
		target := -(nums[i])
		if _, ok := targetSeen[target]; ok {
			continue
		}
		targetSeen[target] = struct{}{}

		twos := twoSum(nums, i+1, target)
		if len(twos) == 0 {
			continue
		}
		for _, t := range twos {
			set := append([]int{nums[i]}, t...)
			k := fmt.Sprintf("%v", set)
			if _, ok := uniques[k]; ok {
				continue
			}
			rs = append(rs, set)
			uniques[k] = struct{}{}
		}
	}
	return rs
}

func twoSum(nums []int, s int, target int) (rs [][]int) {
	rem := map[int]int{}
	for i := s; i < len(nums); i++ {
		r := target - nums[i]
		_, ok := rem[r]
		if ok {
			rs = append(rs, []int{r, nums[i]})
			continue
		}
		rem[nums[i]] = i
	}
	return rs
}
