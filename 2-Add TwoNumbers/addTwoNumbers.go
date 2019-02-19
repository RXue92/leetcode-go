package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	s1 = []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	s2 = []int{6, 7, 8, 6, 7, 8}
	s3 = []int{8, 9, 9, 9}
	s4 = []int{2}
)

func main() {
	out := addTwoNumbers(parser(s3, nil), parser(s4, nil))
	fmt.Println(revParser(parser(s3, nil)))
	fmt.Println("get:", revParser(out))
}

//[2,1,4] to ListNode
func parser(s []int, p *ListNode) *ListNode {
	l := len(s)
	if l == 0 || (l == 1 && s[0] == 0) {
		fmt.Println("invalid slice to convert")
		return nil
	}

	if l == 1 {
		return &ListNode{s[0], p}
	}

	return &ListNode{s[0], parser(s[1:], p)}

}

func revParser(l *ListNode) []int {
	out := []int{}
	out = append(out, l.Val)
	temp := l.Next
	for temp != nil {
		out = append(out, temp.Val)
		temp = temp.Next
	}
	return out
}

// beats 4.55%
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersAux(l1, l2, 0)
}

func addTwoNumbersAux(l1 *ListNode, l2 *ListNode, commit int) *ListNode {
	if l1 == nil && l2 == nil {
		if commit < 1 {
			return nil
		}
		return &ListNode{1, nil}
	}

	// if l1 is null node, use l2 as the first parameter
	if l1 == nil {
		return addTwoNumbersAux(l2, nil, commit)
	}

	if l2 == nil {
		if l1.Val < 10-commit {
			l1.Val += commit
			return l1
		}
		// l1.Val = 0
		if l1.Next == nil {
			return &ListNode{0, &ListNode{1, nil}}
		}

		fmt.Println(l1)
		return &ListNode{0, addTwoNumbersAux(l1.Next, nil, 1)}
	}

	val := l1.Val + l2.Val + commit
	nextCommit := val / 10
	fmt.Println(val, nextCommit)
	return &ListNode{val % 10, addTwoNumbersAux(l1.Next, l2.Next, nextCommit)}

}
