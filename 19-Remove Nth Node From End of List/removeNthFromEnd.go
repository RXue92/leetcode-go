/*
	One pass, but need to store n pointers.
	https://leetcode.com/submissions/detail/221665986/
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
	testNode := ListNode{10, nil}
	out := removeNthFromEnd(&testNode, 1)
	fmt.Println(out)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var tmp= []*ListNode{head}
	next := head.Next
	for next != nil {
		tmp = append(tmp, next)
		if len(tmp) > n+1 {
			tmp = tmp[1:]
		}
		next = next.Next
	}
	fmt.Println(len(tmp))
	if len(tmp) == n {
		return head.Next
	}
	if n > 1 {
		tmp[0].Next = tmp[2]
	} else {
		tmp[0].Next = nil
	}

	return head
}