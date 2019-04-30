/*
	Runtime: 8 ms, faster than 98.51% of Go online submissions for Find First and Last Position of Element in Sorted Array.
	Memory Usage: 4 MB, less than 10.00% of Go online submissions for Find First and Last Position of Element in Sorted Array.
	https://leetcode.com/submissions/detail/225879803/
 */

package problem34

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	idx, aux := findOne(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}
	fmt.Println(aux)
	head, tail := idx, idx
	vsHead, vsTail := true, true
	for vsHead || vsTail {
		if head > 0 && nums[head-1] == target {
			head = head-1
		} else {
			vsHead = false
		}

		if tail < len(nums)-1 && nums[tail+1] == target {
			tail = tail+1
		} else {
			vsTail = false
		}
	}
	return []int{head, tail}
}

func findOne(nums []int , target int) (int, []int) {
	res, left, right := -1, 0, len(nums)-1
	aux := []int{left}
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			aux = append(aux, mid, len(nums)-1)
			return mid, aux
		} else if nums[mid] > target {
			right = mid - 1
			aux = append(aux, right)
		} else {
			left = mid + 1
			aux = append(aux, left)
		}
	}
	aux = append(aux, len(nums)-1)
	return res, aux
}