/*
Runtime: 4 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 2.3 MB, less than 90.48% of Go online submissions for Binary Tree Inorder Traversal.
https://leetcode.com/submissions/detail/228610949/
Iterative solution using stack to store nodes.
 */
package problem94

import "fmt"

func improveSolution(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var out []int
	var stack = []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		fmt.Println(current.Val, len(stack), out)

		if current.Left != nil {
			fmt.Println(current)  // &{2 0xc00000a160 <nil>}
			stack = append(stack, current.Left)
			fmt.Println(stack)  // [0xc00000a140 0xc00000a160]
			current.Left = nil  // set current.Left = nil, but will not affect pointer located at 0xc00000a160
			fmt.Println(current, stack)  // &{2 <nil> <nil>} [0xc00000a140 0xc00000a160]
			continue
		}

		out = append(out, current.Val)
		stack = stack[:len(stack)-1]

		if current.Right != nil {
			fmt.Println(current)
			stack = append(stack, current.Right)
			fmt.Println(stack)
			current.Right = nil
			fmt.Println(current, stack)
		}

	}

	return out
}