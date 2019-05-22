/*
Runtime: 4 ms, faster than 99.66% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.3 MB, less than 95.94% of Go online submissions for Validate Binary Search Tree.
https://leetcode.com/submissions/detail/230644315/
 */


package problem98


const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1


func betterSolution(root *TreeNode) bool {
	return isValidHelper(root, MinInt, MaxInt)
}

func isValidHelper(p *TreeNode, min, max int) bool {
	if p == nil {
		return true
	}

	if p.Val <= min {
		return false
	}

	if p.Val >= max {
		return false
	}

	return isValidHelper(p.Left, min, p.Val) && isValidHelper(p.Right, p.Val, max)
}