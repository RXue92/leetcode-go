/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
==============================

Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
Memory Usage: 2.3 MB, less than 42.14% of Go online submissions for Pascal's Triangle.
https://leetcode.com/submissions/detail/231193847/
 */


package problem118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	out := make([][]int, 0)
	depth := 0
	for depth < numRows {
		depth++
		row := make([]int, 0)
		for i:=0; i<depth; i++ {
			if i == 0 || i == depth-1 {
				row = append(row, 1)
			} else {
				row = append(row, out[depth-2][i-1]+out[depth-2][i])
			}
		}
		out = append(out, row)
	}

	return out
}