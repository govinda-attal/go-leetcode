package mergesortedinplace

// [4,5,6,0,0,0] && [1,2,3]
// [4,5,6,4,5,6] && [1,2,3]
// [1,2,3,4,5,6] && [1,2,3]
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for ; j >= 0; k-- {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
