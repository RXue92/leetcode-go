/*
	Runtime: 8 ms, faster than 64.19% of Go online submissions for Generate Parentheses.
	Memory Usage: 8.4 MB, less than 21.21% of Go online submissions for Generate Parentheses.
	https://leetcode.com/submissions/detail/225684697/
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	res := generateParenthesis(4)
	fmt.Println(res)
	fmt.Println(isValidPathorDone("((", 2))
}

func generateParenthesis(n int) []string {
	return backTrackSearch([]string{"("}, n)
}

func backTrackSearch(current []string, total int) (result []string) {
	for _, substr := range current {
		//fmt.Println(substr)
		isValid, isDone := isValidPathorDone(substr, total)
		if !isValid {
			continue
		} else if isDone {
			result = append(result, substr)
			continue
		}

		result = append(result, backTrackSearch([]string{substr+"("}, total)...)
		result = append(result, backTrackSearch([]string{substr+")"}, total)...)
	}

	//if len(result) == 0 && len(current[0]) == 2*total {
	//	result = current
	//}

	return result
}

func isValidPathorDone(str string, total int) (valid, done bool) {
	ones := strings.Count(str, "(")
	if len(str) < total*2 {
		if len(str) > 2*ones || ones > total{
			return false, true
		}

		return true, false
	}

	if total < ones {
		return false, true
	}

	return true, true
}