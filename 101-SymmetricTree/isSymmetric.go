/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Symmetric Tree.
Memory Usage: 2.9 MB, less than 62.25% of Go online submissions for Symmetric Tree.
https://leetcode.com/submissions/detail/230654349/
Recursive solution.
TODO: iterative solution using stack
 */

package problem101


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}