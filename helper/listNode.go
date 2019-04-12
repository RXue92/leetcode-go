/*
	Parse test cases (eg: [1,2,3]) in leetcode problems (eg: #2 and #19 in algorithms section) to list node
	so that you can test your code locally.
	TODO:1. Handle both int and float; 2. Less modify to code to be submitted when test locally; 3. Test file
 */

package helper

import (
	"fmt"
)

type ListNode struct {
	Val interface{}
	Next *ListNode
}

func Slice2ListNode(nums []interface{}) *ListNode {
	return parser(nums, nil)
}

func parser(s []interface{}, p *ListNode) *ListNode {
	l := len(s)
	if l == 0  {
		fmt.Println("invalid slice to convert")
		return nil
	}

	if l == 1 {
		return &ListNode{s[0], p}
	}

	return &ListNode{s[0], parser(s[1:], p)}

}

func ListNode2Slice(l *ListNode) (out []interface{}) {
	out = append(out, l.Val)
	temp := l.Next
	for temp != nil {
		out = append(out, temp.Val)
		temp = temp.Next
	}
	return out
}

