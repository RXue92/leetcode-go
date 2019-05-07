package problem41


import "testing"


type testCase struct {
	input []int
	expect int
}

var (
	tests = []testCase{
		{[]int{1,2,0}, 3},
		{[]int{3,4,-1,1}, 2},
		{[]int{7,8,9,11,12}, 1},
		{[]int{100,1,4,3,2}, 5},
		{[]int{100,2,4,3,2,1}, 5},
	}
)


func TestFirstMissingPositive(t *testing.T) {
	for i:= range tests {
		output := firstMissingPositive(tests[i].input)
		if tests[i].expect != output {
			t.Errorf("test #%d: expected: %d, output: %d", i, tests[i].expect, output)
		}
	}
}
