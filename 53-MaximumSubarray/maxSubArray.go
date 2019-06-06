/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
===========================

Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Subarray.
Memory Usage: 3.3 MB, less than 92.14% of Go online submissions for Maximum Subarray.
https://leetcode.com/submissions/detail/234070008/
*/
package problem53

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	var ret, p int
	for p < len(nums) {
		ret += nums[p]
		if ret > maxSum {
			maxSum = ret
		}
		if ret < 0 {
			ret = 0
		}
		p++
	}
	return maxSum
}
