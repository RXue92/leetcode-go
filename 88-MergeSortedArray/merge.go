/*
	Runtime: 36 ms, faster than 18.83% of Go online submissions for Merge Sorted Array.
	Memory Usage: 6.1 MB, less than 5.56% of Go online submissions for Merge Sorted Array.
	https://leetcode.com/submissions/detail/227180784/
 */

package problem88

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j, c := 0, 0, 0  // c/j: # nums1/nums2 element visited
	for c < m && j < n {
		fmt.Println(i, j, nums1)
		if nums1[i] > nums2[j] {
			//nums1[i+j+1:m+j+1] = nums1[i+j:m+j]
			//nums1[i+j] = nums2[j]
			nums1 = append(nums1[:i], append([]int{nums2[j]}, nums1[i:len(nums1)-1]...)...)
			j++
		} else {
			c++
		}
		i++
	}

	// if all nums2 elements have been visited, it's done; otherwise copy unvisited nums2 element to nums1
	if c == m {
		copy(nums1[m+j:m+n], nums2[j:n])
	}
}