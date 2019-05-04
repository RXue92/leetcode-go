/*
	Runtime: 4 ms, faster than 100.00% of Go online submissions for Permutations.
	Memory Usage: 7.2 MB, less than 11.11% of Go online submissions for Permutations.
	https://leetcode.com/submissions/detail/226657946/
 */

package problem46

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	tmp := permute(nums[1:])
	l := len(tmp[0])
	var ret [][]int
	for _, s := range tmp {
		for i:=0; i<= l; i++ {
			// s = [1,2], ss = [3,1,2], [1,3,2], [1,2,3]
			ss := make([]int, l+1)
			copy(ss[:i], s[:i])
			ss[i] = nums[0]
			copy(ss[i+1:], s[i:])
			ret = append(ret, ss)
		}
	}

	return ret
}

