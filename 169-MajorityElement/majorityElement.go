/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2
=====================

Runtime: 20 ms, faster than 44.17% of Go online submissions for Majority Element.
Memory Usage: 5.9 MB, less than 9.73% of Go online submissions for Majority Element.
https://leetcode.com/submissions/detail/234418185/
*/
package problem169

func majorityElement(nums []int) int {
	if len(nums) <= 2 {
		return nums[0]
	}

	// eventually, maxFreq > 1
	maxFreq, maxFreqNum := 1, -1
	m := make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 1
		} else {
			m[n]++
			if m[n] > maxFreq {
				maxFreq = m[n]
				maxFreqNum = n
			}
		}

	}
	return maxFreqNum
}
