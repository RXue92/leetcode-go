/*
According to solution, this is the so called "Brute Force" method. Time complexity is O(n^3).
Runtime: 84 ms, faster than 36.73% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Longest Palindromic Substring.
https://leetcode.com/submissions/detail/212506159/
 */

package main

import (
	"fmt"
)


var (
	testCase = "abcda"
)

func longestPalindrome(s string) string {
	if isPalindrome(s) {
		return s
	}
	// Start from longest possible substring and test if it's palindrome. Ignore all substring with length = 1.
	for i:=len(s)-1; i>1; i-- {
		for j:=0; j<=len(s)-i; j++ {
			if isPalindrome(s[j:i+j]) {
				return s[j:j+i]
			}
		}
	}
	return s[0:1]
}

func isPalindrome(s string) bool {
	for i:=0; i<len(s)/2; i++ {
		if s[i]==s[len(s)-1-i] {
			continue
		}
		return false
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome(testCase))
}