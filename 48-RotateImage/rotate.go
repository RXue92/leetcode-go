/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
	Memory Usage: 2.7 MB, less than 19.05% of Go online submissions for Rotate Image.
	https://leetcode.com/submissions/detail/226860431/
 */
package problem48

func rotate(matrix [][]int)  {
	if len(matrix) <= 1 {
		return
	}

	n := len(matrix)
	for i:=0; i<(n+1)/2; i++ {
		for j:=0; j<n/2; j++ {
			currentPos := []int{i, j}
			tmp := matrix[currentPos[0]][currentPos[1]]
			for k:=0; k<3; k++ {
				previousPos := []int{n-1-currentPos[1], currentPos[0]}
				matrix[currentPos[0]][currentPos[1]] = matrix[previousPos[0]][previousPos[1]]
				currentPos = previousPos
			}
			matrix[currentPos[0]][currentPos[1]] = tmp
		}
	}
}
