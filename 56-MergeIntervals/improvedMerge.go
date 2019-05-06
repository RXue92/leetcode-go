/*
	Intuition: sort int slice by starting value, then merge them from left to right (small to big)
	Runtime: 12 ms, faster than 52.49% of Go online submissions for Merge Intervals.
	Memory Usage: 4.9 MB, less than 100.00% of Go online submissions for Merge Intervals.
	https://leetcode.com/submissions/detail/227094263/
*/

package problem56


import "sort"

func betterSolution(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {return intervals[i][0]<=intervals[j][0]})
	var ret = [][]int{intervals[0]}
	for i:=1; i<len(intervals); i++ {
		if ret[len(ret)-1][1] < intervals[i][0] {
			ret = append(ret, intervals[i])
			continue
		}
		if ret[len(ret)-1][1] < intervals[i][1] {
			ret[len(ret)-1][1] = intervals[i][1]
		}
	}

	return ret
}