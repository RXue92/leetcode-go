package problem34

import "testing"

var testCase = []struct {
	input []int
	target int
	expect []int
}{
	{[]int{5,7,7,8,8,10}, 8, []int{3,4}},
	{[]int{5,7,7,8,8,10}, 6, []int{-1,-1}},
	{[]int{1,2,3,4,5}, 4, []int{3,3}},
	{[]int{1,2,3,4,5}, 6, []int{-1,-1}},
	{[]int{2,2,2,2}, 2, []int{0,3}},
	{[]int{2}, 3, []int{-1,-1}},

}

func TestSearchRange(t *testing.T) {
	for _, tc := range testCase {
		answer := searchRange(tc.input, tc.target)
		asExpected := len(answer) == len(tc.expect)
		for i := range answer {
			if answer[i] != tc.expect[i] {
				asExpected = false
				break
			}
		}
		if !asExpected {
			t.Errorf("input %d, expected %d, actual %d", tc.input, tc.expect, answer)
		}
	}
}

func TestBetterSolution(t *testing.T) {
	for _, tc := range testCase {
		answer := betterSolution(tc.input, tc.target)
		asExpected := len(answer) == len(tc.expect)
		for i := range answer {
			if answer[i] != tc.expect[i] {
				asExpected = false
				break
			}
		}
		if !asExpected {
			t.Errorf("input %d, expected %d, actual %d", tc.input, tc.expect, answer)
		}
	}
}
