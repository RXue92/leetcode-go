/*
	Runtime: 8 ms, faster than 100.00% of Go online submissions for Jump Game.
	Memory Usage: 4.2 MB, less than 84.00% of Go online submissions for Jump Game.
	https://leetcode.com/submissions/detail/226926676/
 */

package problem55


func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	target := len(nums) - 1
	for i:=len(nums)-2; i>0; i-- {
		if nums[i]+i>=target {
			target = i
		}
	}
	if nums[0]<target {
		return false
	}
	return true
}