/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
	Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
	https://leetcode.com/submissions/detail/228605954/
	Still a recursive style solution.
 */

package problem94

import "fmt"

func anotherSolution(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	var output []int
	anotherRecursive(root, stack, &output)

	return output
}


func anotherRecursive(p *TreeNode, s []*TreeNode, out *[]int) {
	fmt.Println(p.Val, len(s), out)
	if p.Left != nil {
		s = append(s, p)
		anotherRecursive(p.Left, s, out)
	}
	*out = append(*out, p.Val)
	s = s[:len(s)-1]

	if p.Right != nil {
		s = append(s, p.Right)
		anotherRecursive(p.Right, s, out)
	}
}