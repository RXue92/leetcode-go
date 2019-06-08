/*
Runtime: 52 ms, faster than 96.57% of Go online submissions for Rotate Array.
Memory Usage: 7.6 MB, less than 60.14% of Go online submissions for Rotate Array.
https://leetcode.com/submissions/detail/234429591/
Reverse method according to solution #4.
*/

package problem189

func rotate3(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		end--
		start++
	}
}
