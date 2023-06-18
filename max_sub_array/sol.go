package max_subarray

import (
	"math"
)

// Problem Statement: https://leetcode.com/problems/4sum/

func MaxSubArray(nums []int) int {
	return maxSub(nums, 0, len(nums)-1)
}

func maxSub(nn []int, s, e int) int {
	if s > e {
		return math.MinInt
	}
	if s == e {
		return nn[s]
	}
	m := (s + e) / 2
	l := maxSub(nn, s, m-1)
	r := maxSub(nn, m+1, e)
	mc := maxCrossing(nn, s, m, e)

	return max(l, r, mc)
}

func maxCrossing(nn []int, s, m, e int) int {

	sum := 0
	leftSum := math.MinInt

	for i := m; i >= s; i-- {
		sum = sum + nn[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	sum = 0
	rightSum := math.MinInt
	for i := m; i <= e; i++ {
		sum = sum + nn[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return max(leftSum+rightSum-nn[m], leftSum, rightSum)
}

func max(items ...int) (rs int) {
	rs = math.MinInt
	for _, item := range items {
		if item > rs {
			rs = item
		}
	}
	return rs
}

func MaxSubArrayKA(nums []int) int {
	gmax := nums[0]
	curMax := nums[0]

	for i := 1; i < len(nums); i++ {
		curMax = max(nums[i], curMax+nums[i])
		if curMax > gmax {
			gmax = curMax
		}
	}
	return gmax
}
