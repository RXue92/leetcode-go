package problem70


import "testing"


type testCase struct {
	input int
	expect int
}

var (
	tests = []testCase{
		{0,1},
		{1,1},
		{2,2},
		{3,3},
		{50, 20365011074},
		{30, 1346269},
	}
)

func TestClimbStairs(t *testing.T) {
	for i:= range tests {
		output := climbStairs(tests[i].input)
		if tests[i].expect != output {
			t.Errorf("test #%d: expected: %d, output: %d", i, tests[i].expect, output)
		}
	}
}