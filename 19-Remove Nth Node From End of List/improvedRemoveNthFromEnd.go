/*
	One pass two pointers method according to solution.
	https://leetcode.com/submissions/detail/221668945/
*/
package main

import (
	"fmt"
)


func main() {
	testNode := ListNode{10, nil}
	out := betterSolution(&testNode, 1)
	fmt.Println(out)
}

func betterSolution(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	var first, second = dummy, dummy
	for i:=0; i<n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}