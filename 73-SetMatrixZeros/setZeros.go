/*
	Runtime: 20 ms, faster than 100.00% of Go online submissions for Set Matrix Zeroes.
	Memory Usage: 7.5 MB, less than 7.69% of Go online submissions for Set Matrix Zeroes.
	https://leetcode.com/submissions/detail/227113870/
	Space complexity is O(m+n)
	Time complexity is O(m*n*(m+n))
 */

package problem70


func setZeroes(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}

	zeroRow := make([]int, len(matrix))
	zeroCol := make([]int, len(matrix[0]))
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				zeroRow[i]++
				zeroCol[j]++
			}
		}
	}
	for row, i := range zeroRow {
		if i == 0 {
			continue
		}
		for j := 0; j < len(matrix[0]); j++ {
			matrix[row][j] = 0
		}
	}

	for col, j := range zeroCol {
		if j == 0 {
			continue
		}
		for i:=0; i<len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
}