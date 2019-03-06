/*
Based on "expend around center" method in solution section. Time complexity is O(n^2). Space complexity is constant.
Runtime: 4 ms, faster than 90.28% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Longest Palindromic Substring.
https://leetcode.com/submissions/detail/212765641/
*/

package main

import (
	"fmt"
)


var (
	testCase1 = "aaaaa"
)

// Expending from center
func betterSolution(s string) string {
	if len(s) < 1 {
		return ""
	}

	var start, end int
	for i:=0; i<len(s); i++ {
		l1 := expendAroundCenter(s, i, i)
		l2 := expendAroundCenter(s, i, i+1)
		if l1 < l2 {
			l1 = l2
		}
		if l1 > end - start {
			start = i - (l1+1)/2 + 1
			end = i + l1/2
		}
	}

	return s[start:end+1]
}

func expendAroundCenter(s string, i, j int) int {
	L, R := i, j

	for L>=0 && R<len(s) && s[L]==s[R] {
		L--
		R++
	}
	return R-L-1
}

func main() {
	fmt.Println(betterSolution(testCase1))
}