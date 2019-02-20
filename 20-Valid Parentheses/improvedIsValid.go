/*
Based on solution https://leetcode.com/problems/valid-parentheses/solution/
Slightly better than first solution.
 */
package main

import (
	"fmt"
)

var (
	testCase1="((()))"
)

func main() {
	fmt.Println(betterSolution(testCase1))
}

func betterSolution(s string) bool {
	l := len(s)
	if l%2 == 1 {
		return false
	}

	if s == "" {
		return true
	}

	parenthesesMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	var tempStack []string

	for i:=0; i<l; i++ {
		key, ok := parenthesesMap[s[i:i+1]]
		if !ok {
			tempStack = append([]string{s[i:i+1]}, tempStack...)
		} else if len(tempStack) > 0 && key == tempStack[0]{
			tempStack = tempStack[1:]
		} else {
			return false
		}
	}

	if len(tempStack) > 0 {
		return false
	}

	return true
}