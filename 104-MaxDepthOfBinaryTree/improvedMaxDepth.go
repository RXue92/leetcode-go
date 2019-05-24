/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.4 MB, less than 93.77% of Go online submissions for Maximum Depth of Binary Tree.
https://leetcode.com/submissions/detail/231023932/
A more concise solution according to
https://leetcode.com/problems/maximum-depth-of-binary-tree/discuss/34216/Simple-solution-using-Java
 */
package problem104

import . "github.com/rxue92/leetcode-go/helper"

func betterSolution(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := betterSolution(root.Left)
	rightDepth := betterSolution(root.Right)

	if leftDepth > rightDepth {
		return leftDepth+1
	}
	return rightDepth+1
}
