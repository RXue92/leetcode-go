/*
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
======================
Runtime: 32 ms, faster than 31.91% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
Memory Usage: 26.8 MB, less than 60.00% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
https://leetcode.com/submissions/detail/231063874/
Related to depth first search.
*/

package problem105

import (
	"fmt"
	. "github.com/rxue92/leetcode-go/helper"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	auxNode := &TreeNode{0, nil, nil}
	buildHelper(preorder, inorder, auxNode, true)
	fmt.Println(auxNode)
	return auxNode.Left

}

func buildHelper(preorder, inorder []int, parent *TreeNode, isLeft bool) {
	fmt.Println("p:", parent)
	if len(preorder) == 0 {
		if isLeft {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	}
	val := preorder[0]
	fmt.Println(val, isLeft)
	root := &TreeNode{val, nil, nil}
	if isLeft {
		parent.Left = root
	} else {
		parent.Right = root
	}
	if len(preorder) == 1 {
		return
	}

	idx := findIndex(val, inorder)
	buildHelper(preorder[1:idx+1], inorder[:idx], root, true)
	buildHelper(preorder[idx+1:], inorder[idx+1:], root, false)
}

// assume no duplicate
func findIndex(n int, nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == n {
			return i
		}
	}
	return -1
}
