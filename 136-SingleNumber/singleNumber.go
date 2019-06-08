/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
========================
Runtime: 16 ms, faster than 18.10% of Go online submissions for Single Number.
Memory Usage: 5.8 MB, less than 22.28% of Go online submissions for Single Number.
https://leetcode.com/submissions/detail/234402892/
*/
package problem136

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
		} else {
			m[n] = 1
		}
	}

	for k := range m {
		return k
	}
	return -1
}
