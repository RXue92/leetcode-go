/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Spiral Matrix.
	Memory Usage: 2 MB, less than 100.00% of Go online submissions for Spiral Matrix.
	https://leetcode.com/submissions/detail/226914183/
 */

package problem54

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	const (
		right = iota
		down
		left
		up
	)

	m, n := len(matrix), len(matrix[0])
	fmt.Println(m, n)
	count := m*n
	ret := []int{}
	direction := right
	var i, j, c int
	leftMost, upMost := -1, 0

	for c < count {
		fmt.Println(i, j, direction, c)
		switch direction {
		case down:
			if i < m {
				ret = append(ret, matrix[i][j])
				c++
				i++
				continue
			}
			i--
			j--
			direction++
			m--
		case right:
			if j < n {
				ret = append(ret, matrix[i][j])
				c++
				j++
				continue
			}
			i++
			j--
			direction++
			n--
		case up:
			if i > upMost {
				ret = append(ret, matrix[i][j])
				c++
				i--
				continue
			}
			direction++
			i++
			j++
			upMost++
		case left:
			if j > leftMost {
				ret = append(ret, matrix[i][j])
				c++
				j--
				continue
			}
			direction++
			j++
			i--
			leftMost++
		default:  // direction==4
			direction = direction - 4
		}

	}
	return ret
}