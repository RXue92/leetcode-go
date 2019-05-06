/*
	https://leetcode.com/submissions/detail/227139800/
	Two-pass implementation.
 */

package problem75


func sortColors(nums []int)  {
	tmp := make([]int, 3)
	for _, i := range nums {
		tmp[i]++
	}

	i, j := 0, 0
	for i < len(nums) {
		if tmp[j] == 0 {
			j++
			continue
		}

		nums[i] = j
		tmp[j]--
		i++
	}
}