/*
Runtime: 8 ms, faster than 98.10% of Go online submissions for Single Number.
Memory Usage: 4.7 MB, less than 72.52% of Go online submissions for Single Number.
https://leetcode.com/submissions/detail/234404320/

XOR: a ^ b, in a or b but not in both (bitwise)
property:
1) 0^N = N
2) N ^ N  = 0
3) communicative and associative
In golang, bitwise operations and shit operation apply to integers only.
See more bitwise operation in https://homerl.github.io/2016/03/29/golang-bitwise-operators/
*/
package problem136

func betterSolution(nums []int) int {
	ret := 0
	for _, n := range nums {
		ret ^= n // XOR
	}

	return ret
}
