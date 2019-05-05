package problem48

import (
	"testing"
)

type testCase struct {
	input [][]int
	expect [][]int
}

var (
	test1 = testCase{[][]int{{1,2,3}, {4,5,6}, {7,8,9}}, [][]int{{7,4,1}, {8,5,2}, {9,6,3}}}
	test2 = testCase{[][]int{{5,1,9,11}, {2,4,8,10}, {13,3,6,7}, {15,14,12,16}},
	[][]int{{15,13,2,5}, {14,3,4,1}, {12,6,8,9}, {16,7,10,11}}}
	test3 = testCase{[][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}, {13,14,15,16}},
	[][]int{{13,9,5,1}, {14,10,6,2}, {15,11,7,3}, {16,12,8,4}}}
	cases = []testCase{test2, test1, test3}
)

func TestRotate(t *testing.T) {
	for _, tc := range cases {
		rotate(tc.input)
		if !SameMatrix(tc.input, tc.expect) {
			t.Errorf("matrix not equal: after rotation = %d, expect: %d\n", tc.input, tc.expect)
		}
	}

}


func TestBetterSolution(t *testing.T) {
	for _, tc := range cases {
		betterSolution(tc.input)
		if !SameMatrix(tc.input, tc.expect) {
			t.Errorf("matrix not equal: after rotation = %d, expect: %d\n", tc.input, tc.expect)
		}
	}
}

func SameMatrix(target, source [][]int) bool {
	for i:=0; i<len(target); i++ {
		for j:=0; j<len(target); j++ {
			if target[i][j] != source[i][j] {
				//fmt.Println(i, j)
				return false
			}
		}
	}
	return true
}