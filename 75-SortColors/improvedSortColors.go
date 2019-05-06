/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
	Memory Usage: 2.3 MB, less than 80.00% of Go online submissions for Sort Colors.
	https://leetcode.com/submissions/detail/227143657/
	One-pass implementation according to unofficial solution, related to
	https://en.wikipedia.org/wiki/Dutch_national_flag_problem
 */

package problem75


func betterSolution(nums []int) {
	// red:0, white:1, blue:2
	red, white, blue := 0, 0, len(nums)-1

	for white <= blue {
		if nums[white] == 0 {
			nums[white], nums[red] = nums[red], nums[white]
			white++
			red++
		} else if nums[white] == 1 {
			white++
		} else {
			nums[white], nums[blue] = nums[blue], nums[white]
			blue--
		}
	}
}