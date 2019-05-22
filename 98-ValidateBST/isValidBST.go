/*
	Runtime: 16 ms, faster than 6.38% of Go online submissions for Validate Binary Search Tree.
	Memory Usage: 6.2 MB, less than 6.71% of Go online submissions for Validate Binary Search Tree.
	https://leetcode.com/submissions/detail/230639692/
	Inorder tree traversal must return an ascending slice.
	According to solution, it's not necessary to keep the whole slice if we valid BST through traversal.
 */


package problem98


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	inorderTree := inorderTraversal(root)

	for i:=0; i<len(inorderTree)-1; i++ {
		if inorderTree[i] >= inorderTree[i+1] {
			return false
		}
	}
	return true
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