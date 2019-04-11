/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
	Memory Usage: 2.5 MB, less than 87.50% of Go online submissions for Merge Two Sorted Lists.
	https://leetcode.com/submissions/detail/221724847/
 */
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	testNode1 := ListNode{10, nil}
	testNode2 := ListNode{11, nil}
	out := mergeTwoLists(&testNode1, &testNode2)
	fmt.Println(out)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil {
		return l1
	}
	l1 = &ListNode{0, l1}
	here := l1
	first, second := l1.Next, l2
	for here.Next != nil {
		if first.Val <= second.Val {
			here.Next = first
			first = first.Next
		} else {
			here.Next = second
			second = second.Next
		}
		here = here.Next
	}

	if first == nil {
		here.Next = second
	}

	if second == nil {
		here.Next = first
	}

	return l1.Next
}

/*
	Memory Usage: 2.5 MB, less than 54.46% of Go online submissions for Merge Two Sorted Lists.
	https://leetcode.com/submissions/detail/221727526/
 */
func mergeTwoListsRecursivly(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil {
		return l1
	}

	if l1 == nil {
		return l2
	}

	out := l1
	if l1.Val <= l2.Val {
		out = l1
		out.Next = mergeTwoListsRecursivly(l1.Next, l2)
	} else {
		out = l2
		out.Next = mergeTwoListsRecursivly(l1, l2.Next)
	}

	return out
}