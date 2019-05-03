/*
	Runtime: 8 ms, faster than 22.18% of Go online submissions for Valid Sudoku.
	Memory Usage: 3.1 MB, less than 50.00% of Go online submissions for Valid Sudoku.
	https://leetcode.com/submissions/detail/226511203/
*/
package problem36

import "fmt"


func betterSolution(board [][]byte) bool {
	for i:=0; i<9; i++ {
		line := make([]int, 9)
		column := make([]int, 9)
		box := make([]int, 9)
		for j:=0; j<9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0' -1
				if line[num] == 1 {
					return false
				}
				line[num] = 1
			}


			if board[j][i] != '.' {
				num := board[j][i] - '0' -1
				if column[num] == 1 {
					return false
				}
				column[num] = 1
			}


			// only i decided upper left point of box
			x := (i/3)*3 + j/3
			y := (i%3)*3 + j%3

			if board[x][y] != '.' {
				fmt.Println([]int{x, y})
				num := board[x][y] - '0' -1
				if box[num] == 1 {
					return false
				}
				box[num] = 1
			}
		}
	}
	return true
}