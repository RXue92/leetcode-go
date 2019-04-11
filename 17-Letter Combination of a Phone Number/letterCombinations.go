/*
	Memory Usage: 2.7 MB, less than 29.41% of Go online submissions for Letter Combinations of a Phone Number.
	https://leetcode.com/submissions/detail/221632651/
 */
package main

import (
	"strings"
	"fmt"
	"strconv"
)


func main() {
	digits := "023"
	out := letterCombinations(digits)
	fmt.Println(out)
}

func makeBoard() map[string][]string {
	var oneBoard = map[int]string {
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

    board := make(map[string][]string)
	for k, v := range oneBoard {
		board[strconv.Itoa(k)] = strings.Split(v, "")
	}

	return board
}

var finalBoard = makeBoard()

func letterCombinations(digits string) []string {
	var tmp []string
	//fmt.Println(finalBoard)
	for i:=0; i<len(digits); i++ {
		digit := digits[i:i+1]
		if digit == "0" || digit == "1" {
			continue
		}
		tmp = sliceMultiply(tmp, finalBoard[digit])
	}
	return tmp
}


// return all possible combination of strings, and return ss1 if ss2 is empty
func sliceMultiply(ss1, ss2 []string) (out []string) {
	if len(ss1) < len(ss2) {
		return sliceMultiply(ss2, ss1)
	}

	if len(ss2) == 0 {
		return ss1
	}

	for _, c1 := range ss1 {
		for _, c2 := range ss2 {
			out = append(out, c1+c2)
		}
	}

	return out
}
