/*
	https://leetcode.com/submissions/detail/227101758/
	Time Complexity is O(n)
	Return may overflow if n is too large, need to change return value type.
	TODO: use power of matrix to achieve O(log(n))
 */

package problem70


func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	rem := [3]int{1,1,2}
	count := 2
	for count < n {
		rem[0], rem[1], rem[2] = rem[1], rem[2], rem[1]+rem[2]
		count++
	}
	return rem[2]
}