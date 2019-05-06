package problem56


import (
	"testing"
	"github.com/rxue92/leetcode-go/helper"
)


type testCase struct {
	input string
	expect string
}

var (
	tests = []testCase{
		{"[[2,3],[5,7],[0,1],[4,6],[5,5],[4,6],[5,6],[1,3],[4,4],[0,0],[3,5],[2,2]]", "[[0,7]]"},
		{"[[1,3],[2,6],[8,10],[15,18]]", "[[1,6],[8,10],[15,18]]"},
		{"[[0,0],[1,2],[5,5],[2,4],[3,3],[5,6],[5,6],[4,6],[0,0],[1,2],[0,2],[4,5]]", "[[0,6]]"},
		{"[[2,3],[4,6],[5,7],[3,4]]", "[[2,7]]"},
		{"[[2,3],[4,5],[6,7],[8,9],[1,10]]", "[[1,10]]"},
		{"[[2,3],[2,2],[3,3],[1,3],[5,7],[2,2],[4,6]]", "[[1,3],[4,7]]"},
		{"[[2,3],[5,5],[2,2],[3,4],[3,4]]", "[[2,4],[5,5]]"},
		{"[[1,4],[0,2]]", "[[0,4]]"},
	}
)


func TestMerge(t *testing.T) {
	for i := range tests {
		output := merge(helper.StringTo2DArray(tests[i].input))
		expected := helper.StringTo2DArray(tests[i].expect)
		if !helper.Compare2DArray(output, expected) {
			t.Errorf("test #%d: expected: %d, output: %d", i, expected, output)
		}
	}
}


func TestBetterSolution(t *testing.T) {
	for i := range tests {
		output := betterSolution(helper.StringTo2DArray(tests[i].input))
		expected := helper.StringTo2DArray(tests[i].expect)
		if !helper.Compare2DArray(output, expected) {
			t.Errorf("test #%d: expected: %d, output: %d", i, expected, output)
		}
	}
}