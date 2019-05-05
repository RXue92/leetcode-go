package problem49

import "testing"

type testCase struct {
	input []string
	expect [][]string
}

var (
	test1 = testCase{[]string{"eate","teat","tan","ate","nat","bat"},
	[][]string{{"tan", "nat"}, {"eate"}, }}
	cases = []testCase{test1}
)
func TestGroupAnagrams(t *testing.T) {

}