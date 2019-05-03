/*
	Runtime: 20 ms, faster than 6.45% of Go online submissions for Valid Sudoku.
	Memory Usage: 6 MB, less than 12.50% of Go online submissions for Valid Sudoku.
	https://leetcode.com/submissions/detail/226491211/
*/
package problem36

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var holders = make([]holder, 18)
	var boxes = make(map[[2]int]*holder)
	for i:=0; i<18; i++ {
		// 0-8: each row; 9-17: each column; 18-26: each sub-box
		holders[i] = holder{make(map[byte]interface{}), true}
	}
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			boxes[[2]int{i,j}] = &holder{make(map[byte]interface{}), true}
		}
	}
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			if string(board[i][j]) == "." {
				continue
			}
			holders[i].insert(board[i][j])
			holders[j+9].insert(board[i][j])
			box := boxes[[2]int{i/3, j/3}]
			box.insert(board[i][j])
		}
	}

	for _, h := range holders {
		if !h.valid {
			return false
		}
	}
	fmt.Println("lines ok")

	for k, box := range boxes {
		fmt.Println(k, box)
		if !box.valid {
			return false
		}
	}

	return true
}

type holder struct {
	data map[byte]interface{}
	valid bool
}

func (h *holder) insert(b byte) {
	if string(b) == "."  {
		return
	}
	_, ok := h.data[b]
	if ok {
		h.valid = false
		return
	}
	h.data[b] = b
}