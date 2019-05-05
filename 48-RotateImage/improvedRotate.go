/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
	Memory Usage: 2.6 MB, less than 23.81% of Go online submissions for Rotate Image.
	https://leetcode.com/submissions/detail/226864358/
*/

package problem48

func betterSolution(matrix [][]int)  {
	if len(matrix) <= 1 {
		return
	}

	n := len(matrix)
	for i:=0; i<(n+1)/2; i++ {
		for j:=0; j<n/2; j++ {
			matrix[i][j], matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i] =
				matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i], matrix[i][j]
		}
	}
}
