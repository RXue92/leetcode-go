/*
	Rewrite find min index and bisection search using for loop
	Runtime: 4 ms, faster than 42.97% of Go online submissions for Search in Rotated Sorted Array.
	Memory Usage: 2.6 MB, less than 57.69% of Go online submissions for Search in Rotated Sorted Array.
	https://leetcode.com/submissions/detail/225827360/
 */

package main

import (
	"fmt"
)

func main() {
	test := []int{4,5,6,7,0,1,2}
	res := betterSolution(test, 9)
	fmt.Println(res)
}

func betterSolution(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	minIndex := findOffset(nums)
	fmt.Println(minIndex)
	sorted := append(nums[minIndex:], nums[:minIndex]...)
	left, right := 0, len(sorted)-1
	for left <= right {
		fmt.Println(left, right)
		if sorted[(left+right)/2] > target {
			right = (left+right)/2 - 1
		} else if sorted[(left+right)/2] < target{
			left = (right+left)/2 + 1
		} else {
			return ((right+left)/2 + minIndex) % len(sorted)
		}
	}
	return -1
}

func findOffset(nums []int) int {
	res, left, right := 0, 0, len(nums)-1
	for nums[left] > nums[right] {
		if right - left == 1 {
			res = right
			break
		}
		if nums[(left+right)/2] > nums[right] {
			left = (left+right)/2
		} else {
			right = (left+right)/2
		}
	}

	return res
}



