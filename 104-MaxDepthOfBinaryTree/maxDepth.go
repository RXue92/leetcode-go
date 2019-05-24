/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.

========================

Runtime: 4 ms, faster than 99.19% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.7 MB, less than 5.27% of Go online submissions for Maximum Depth of Binary Tree.
https://leetcode.com/submissions/detail/231022636/

*/
package problem104

import . "github.com/rxue92/leetcode-go/helper"

func maxDepth(root *TreeNode) int {
	depth := 0
	tmp := make([]int, 0)

	digHelper(root, &tmp, depth)

	return len(tmp)
}

func digHelper(node *TreeNode, keeper *[]int, depth int) {
	if node == nil {
		return
	}

	if len(*keeper) < depth+1 {
		*keeper = append(*keeper, 0)
	}

	(*keeper)[depth] += 1
	digHelper(node.Left, keeper, depth+1)
	digHelper(node.Right, keeper, depth+1)
}
