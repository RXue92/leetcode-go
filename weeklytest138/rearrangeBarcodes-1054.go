/*
In a warehouse, there is a row of barcodes, where the i-th barcode is barcodes[i].

Rearrange the barcodes so that no two adjacent barcodes are equal.  You may return any answer, and it is guaranteed an answer exists.



Example 1:

Input: [1,1,1,2,2,2]
Output: [2,1,2,1,2,1]
Example 2:

Input: [1,1,1,1,2,2,3,3]
Output: [1,3,1,3,2,1,2,1]


Note:

1 <= barcodes.length <= 10000
1 <= barcodes[i] <= 10000

===========================
https://leetcode.com/submissions/detail/232004676/
TODO: use [][]int to replace Counter and fill out in one pass (skip 2).
*/

package weeklytest138

import (
	"fmt"
	"sort"
)

type Counter struct {
	num   int
	count int
}

type ByCount []Counter

func (a ByCount) Less(i, j int) bool { return a[i].count > a[j].count } // descending
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Len() int           { return len(a) }

func rearrangeBarcodes(barcodes []int) []int {
	m := make(map[int]int)
	for _, c := range barcodes {
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	var CodeCounts []Counter
	for k, v := range m {
		CodeCounts = append(CodeCounts, Counter{k, v})
	}

	sort.Sort(ByCount(CodeCounts))
	fmt.Println(CodeCounts)

	var out = make([]int, 0)
	//c := 0
	//flag := false
	//var last = CodeCounts[1].num
	//for c < len(barcodes) {
	//	fmt.Println(out)
	//	if last != CodeCounts[0].num {
	//		out = append(out, CodeCounts[0].num)
	//		last = CodeCounts[0].num
	//		CodeCounts[0].count--
	//	} else {
	//		out = append(out, CodeCounts[1].num)
	//		last = CodeCounts[1].num
	//		CodeCounts[1].count--
	//	}
	//	if CodeCounts[1].count == 0 && len(CodeCounts) > 2{
	//		CodeCounts = append(CodeCounts[:1], CodeCounts[2:]...)
	//	}
	//	if CodeCounts[0].count == 0 && len(CodeCounts) > 2 {
	//		CodeCounts[0], CodeCounts[2] = CodeCounts[2], CodeCounts[0]
	//		CodeCounts = append(CodeCounts[:2], CodeCounts[3:]...)
	//	}
	//	c++
	//}

	j := 0
	for i := 0; i < len(barcodes); i++ {
		if i%2 == 1 {
			out = append(out, 0)
		} else {
			if CodeCounts[j].count == 0 {
				j++
			}
			out = append(out, CodeCounts[j].num)
			CodeCounts[j].count--
		}
	}

	for i := 0; i < len(barcodes); i++ {
		if CodeCounts[j].count == 0 {
			j++
		}
		if i%2 == 1 {
			out[i] = CodeCounts[j].num
			CodeCounts[j].count--
		}
	}

	return out
}
