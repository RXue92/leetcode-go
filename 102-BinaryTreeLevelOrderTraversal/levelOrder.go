/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]

=================================

Runtime: 4 ms, faster than 40.83% of Go online submissions for Binary Tree Level Order Traversal.
Memory Usage: 5.9 MB, less than 58.47% of Go online submissions for Binary Tree Level Order Traversal.
https://leetcode.com/submissions/detail/230975176/
Iterative solution.
*/
package problem102


import (
	. "github.com/rxue92/leetcode-go/helper"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	out := [][]int{{root.Val}}
	current := []*TreeNode{root}
	for  {
		var  next []*TreeNode
		allNil := true
		var nextVal []int
		for i:=0; i<len(current); i++ {
			if current[i].Left != nil {
				next = append(next, current[i].Left)
				nextVal = append(nextVal, current[i].Left.Val)
				allNil = false
			}
			if current[i].Right != nil {
				next = append(next, current[i].Right)
				nextVal = append(nextVal, current[i].Right.Val)
				allNil = false
			}
		}
		if allNil {
			break
		}
		out = append(out, nextVal)
		current = next  // reset next
	}

	return out
}