/*
Based on fastest go solution.
 */
package main

import (
	"fmt"
)

var (
	testCase2="((()))"
)

func main() {
	fmt.Println(betterSolution2(testCase2))
}

func bestSolution(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}

		if c == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		if c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		if c == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func betterSolution2(s string) bool {
	allParentheses := "()[]{}"
	stack := make([]byte, 0)

	for i:=0; i<len(s); i++ {
		if s[i] == allParentheses[0] || s[i] == allParentheses[2] || s[i] == allParentheses[4] {
			stack = append(stack, s[i])
			continue
		}

		if s[i] == allParentheses[1] {
			if len(stack) == 0 || stack[len(stack)-1] != allParentheses[0] {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		if s[i] == allParentheses[3] {
			if len(stack) == 0 || stack[len(stack)-1] != allParentheses[2] {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		if s[i] == allParentheses[5] {
			if len(stack) == 0 || stack[len(stack)-1] != allParentheses[4] {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}