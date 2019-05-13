/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
	Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
	https://leetcode.com/submissions/detail/227936282/
	Recursive style
 */


package problem94

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var out []int
	inorder(root, &out)
	return out
}

// visiting order: left subtree, node, right subtree
func inorder(p *TreeNode, out *[]int) {
	if p.Left != nil {
		inorder(p.Left, out)
	}
	*out = append(*out, p.Val)

	if p.Right != nil {
		inorder(p.Right, out)
	}
}
