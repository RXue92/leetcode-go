/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
Memory Usage: 6.1 MB, less than 28.87% of Go online submissions for Binary Tree Level Order Traversal.
https://leetcode.com/submissions/detail/230984349/
Recursive solution according to user solution.
 */
package problem102


func betterSolution(root *TreeNode) [][]int {
	depth := 0
	out := make([][]int, 0)
	helper(root, &out, depth)
	return out
}

func helper(node *TreeNode, res *[][]int, depth int) {
	if node == nil {
		return
	}

	if len(*res) < depth+1 {
		*res = append(*res, make([]int, 0))
	}

	(*res)[depth] = append((*res)[depth], node.Val)

	// descendant of left child always prevail descendant of right child at same level
	helper(node.Left, res, depth+1)
	helper(node.Right, res, depth+1)
}
