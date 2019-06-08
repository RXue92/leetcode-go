/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
Note:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?
*/
package problem189

/*
Runtime: 172 ms, faster than 18.95% of Go online submissions for Rotate Array.
Memory Usage: 7.6 MB, less than 60.14% of Go online submissions for Rotate Array.
https://leetcode.com/submissions/detail/234425391/
*/
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	tmp := nums[len(nums)-1]
	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i-1]
	}

	nums[0] = tmp
	rotate(nums, k-1)
}

/*
Runtime: 52 ms, faster than 96.57% of Go online submissions for Rotate Array.
Memory Usage: 7.6 MB, less than 66.89% of Go online submissions for Rotate Array.
https://leetcode.com/submissions/detail/234428189/
*/
func rotate2(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	n := len(nums) - k
	/*
		nums = append(nums, nums[:n]...)
		nums = nums[n:]
		nums is changed but the backing array is extended, so leetcode will not let you pass
	*/

	tmp := make([]int, k)
	copy(tmp, nums[n:])
	copy(nums[k:], nums[:n])
	//fmt.Println(nums)
	copy(nums[:k], tmp)
	//fmt.Println(nums)
}
