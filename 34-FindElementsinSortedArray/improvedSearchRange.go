/*
	Implement binary search method according solution
	Runtime: 8 ms, faster than 98.51% of Go online submissions for Find First and Last Position of Element in Sorted Array.
	Memory Usage: 4.1 MB, less than 5.00% of Go online submissions for Find First and Last Position of Element in Sorted Array.
	https://leetcode.com/submissions/detail/225890768/
 */
package problem34

func betterSolution(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftmostIdx := findEdge(nums, target, true)
	//fmt.Println(leftmostIdx)
	if leftmostIdx == len(nums) || nums[leftmostIdx] != target {
		return []int{-1, -1}
	}

	rightmostIdx := findEdge(nums, target, false)
	return []int{leftmostIdx, rightmostIdx-1}   // subtract 1
}

// Return index leftmost (rightmost, if left==false) which equals to target
func findEdge(nums []int, target int, left bool) int {
	head, tail := 0, len(nums)
	for head < tail {
		mid := (head+tail)/2
		if nums[mid] > target || (left && nums[mid] == target) { // move pointer on the right of index
			tail = mid
		} else {
			head = mid+1  // guarantee conversion
		}
	}
	return head
}