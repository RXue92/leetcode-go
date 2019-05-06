package helper

import (
	"strconv"
	"strings"
)

// [[1,2], [3,4]] to [][]int{{1,2}, {3,4}}
func StringTo2DArray(input string) [][]int {
	input = strings.TrimSpace(input)
	input = input[1 : len(input)-1]  //  remove outer []
	var ret [][]int
	var tmp []int
	var lastChar string
	for i := 0; i < len(input); i++ {
		char := lastChar + input[i : i+1]
		if char == "[" || char == "," || char == " " || char == "\n" {
			continue
		}
		if char == "]" {
			ret = append(ret, tmp)
			tmp = []int{}
			continue
		}
		_, err := strconv.Atoi(char)
		if err == nil {
			// handle multiple digits int
			lastChar = char
			continue
		}
		num, err := strconv.Atoi(lastChar)
		if err != nil {
			panic(err)
		}

		tmp = append(tmp, num)
		lastChar = ""
		if input[i:i+1] == "]" {
			ret = append(ret, tmp)
			tmp = []int{}
		}
	}
	return ret
}

// TODO: invariant to permutation of slices in 2DArray
func Compare2DArray(a1, a2 [][]int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i := 0; i < len(a1); i++ {
		if len(a1[i]) != len(a2[i]) {
			return false
		}
		for j := 0; j < len(a1[i]); j++ {
			if a1[i][j] != a2[i][j] {
				return false
			}
		}
	}
	return true
}
