/*
	https://leetcode.com/submissions/detail/227379400/
 */
package problem88

// fill nums1 from end to toe
func similarSolution(nums1 []int, m int, nums2 []int, n int)  {
	idx := m+n-1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[idx] = nums1[m-1]
			m--
		} else {
			nums1[idx] = nums2[n-1]
			n--
		}
		idx--
	}

	if m == 0 {
		copy(nums1[:n], nums2[:n])
	}
}
