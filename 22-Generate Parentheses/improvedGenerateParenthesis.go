/*
	Runtime: 8 ms, faster than 64.19% of Go online submissions for Generate Parentheses.
	Memory Usage: 8.4 MB, less than 21.21% of Go online submissions for Generate Parentheses.
	https://leetcode.com/submissions/detail/225684697/
 */
package main

import (
	"fmt"
)

func main() {
	res := betterSolution(3)
	fmt.Println(res)
}

func betterSolution(n int) []string {
	var result []string
	_generate(n, "(", 1, 0, &result)
	return result
}

func _generate(total int, current string, left, right int, result *[]string) {
	if len(current) == 2*total {
		*result = append(*result, current)
	}

	if left < total {
		_generate(total, current+"(", left+1, right, result)
	}

	if right < left && right < total {
		_generate(total, current+")", left, right+1, result)
	}
}