/*
Runtime: 12 ms, faster than 99.83% of Go online submissions for Majority Element.
Memory Usage: 5.9 MB, less than 74.16% of Go online submissions for Majority Element.
https://leetcode.com/submissions/detail/234422810/

Boyer-Moore Voting Algorithm according to solution #6.
Intuition: Clear candidate when count == 0, and the final candidate must be the answer
since we cannot drop more minority than majority.
*/
package problem169

func betterSolutoin(nums []int) int {
	count, candidate := 0, 0

	for _, n := range nums {
		if count == 0 {
			candidate = n
		}

		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
