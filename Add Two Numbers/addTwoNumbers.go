package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}


//this method need to handle over flow issues

func main() {
	//s1 := []int{2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9}
	//s2 := []int{1,2,3,5,6,4,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9,9,9,9}
	s1 := []int{1,2,3,4,5,1,2,3,4,5,1,2,3,4,5,1,2,3,4,5,1,2,3,4,5,1,2,3,4,5,1,2,3,4,5,1,2,3,4,5}
	s2 := []int{6,7,8,6,7,8}
	out := addTwoNumbers(parser(s1), parser(s2))
	fmt.Println("get:", out)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var out *ListNode
	s1 := getdigits(l1)
	s2 := getdigits(l2)
	s := adder(s1, s2)
	out = parser(s)
	return out
}

//[2,1,4] to ListNode
func parser(s []int) *ListNode {
	var head ListNode
	l := len(s)
	head.Val = s[l-1]
	if l <= 1 {
		return &head
	} else {
		for i := 2; i <= l; i++ {
			tail := head
			head.Val = s[l-i]
			head.Next = &tail
		}
		return &head
	}
}

//214 -> [4,1,2]
func adder(s1 []int, s2 []int) []int {
	n1 := toNum(s1)
	fmt.Println(n1)
	n2 := toNum(s2)
	fmt.Println(n2)
	n := n1 + n2
	fmt.Println(n)
	out := []int{}
	for n > 0 {
		out = append(out, n%10)
		n = n / 10
	}
	if len(out) == 0 {
		out = append(out, 0)
	}
	return out
}
 //overflow
func toNum(s []int) int {
	base := 1
	var sum int
	for _, v := range s {
		sum = sum + v*base
		base = base * 10
	}
	return sum
}

func getdigits(l *ListNode) []int {
	out := []int{}
	out = append(out, l.Val)
	temp := l.Next
	for temp != nil {
		out = append(out, temp.Val)
		temp = temp.Next
	}
	return out
}
