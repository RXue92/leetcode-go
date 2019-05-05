package problem54

import (
	"testing"
)

type testCase struct {
	input [][]int
	expected []int
}

var (
	test1 = testCase{[][]int{{1,2,3}, {4,5,6}, {7,8,9}}, []int{1,2,3,6,9,8,7,4,5}}
	test2 = testCase{[][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}}, []int{1,2,3,4,8,12,11,10,9,5,6,7}}
	cases = []testCase{test2}
)

func TestSpiralOrder(t *testing.T) {
	for _, tc := range cases {
		output := spiralOrder(tc.input)
		if len(tc.expected) != len(output) {
			t.Errorf("expected: %d, output: %d", tc.expected, output)
		}
		for i:=0; i<len(tc.expected); i++ {
			if tc.expected[i] != output[i] {
				t.Errorf("expected: %d, output: %d", tc.expected, output)
				break
			}
		}
	}
}