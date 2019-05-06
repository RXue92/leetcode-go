/*
	Runtime: 32 ms, faster than 17.65% of Go online submissions for Set Matrix Zeroes.
	Memory Usage: 7.4 MB, less than 15.38% of Go online submissions for Set Matrix Zeroes.
	https://leetcode.com/submissions/detail/227128201/
	Space complexity is O(1)
	Could only use one extra variable
 */

package problem70

import "fmt"

func betterSolution(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	col0, row0 := false, false
	for i:=0; i<len(matrix); i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j:=0; j<len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				if i==0 {
					row0 = true
				}
			}
		}
	}

	fmt.Println(matrix)
	fmt.Println(row0, col0)

	if len(matrix) == 1 || len(matrix[0]) == 1 {
		if matrix[0][0] != 0 {
			return
		}
		for i:=0; i<len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
		return
	}

	for i:=1; i<len(matrix); i++ {
		for j:=1; j<len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	fmt.Println(matrix)


	if col0 {
		for i:=0; i<len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	if row0 {
		for j:=0; j<len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

}