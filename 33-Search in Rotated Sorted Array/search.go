/*
	Runtime: 4 ms, faster than 42.97% of Go online submissions for Search in Rotated Sorted Array.
	Memory Usage: 3 MB, less than 7.69% of Go online submissions for Search in Rotated Sorted Array.
	https://leetcode.com/submissions/detail/225804336/
 */

package main

import (
	"fmt"
)

func main() {
	test := []int{4}
	res := search(test, 4)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	minIndex := findMinIndex(nums, 0, len(nums)-1)
	res := bisec(append(nums[minIndex:], nums[:minIndex]...), target, 0, len(nums)-1)
	if res == -1 {
		return res
	}
	return (res + minIndex) % len(nums)
}


func findMinIndex(nums []int, head, tail int) int {
	fmt.Println(head, tail)
	if nums[head] <= nums[tail] {
		return head
	}

	if len(nums[head:tail+1]) == 2 {
		return tail
	}

	mid := (head+tail)/2
	if nums[tail] > nums[mid] {
		return findMinIndex(nums, head, mid)
	}

	return findMinIndex(nums, mid, tail)
}

func bisec(nums []int, target int, left, right int) int {
	fmt.Println(nums[left:right+1])
	if nums[left] == target {
		return left
	}

	if nums[right] == target {
		return right
	}

	if len(nums[left:right+1]) == 1 {
		return -1
	}

	if nums[(left+right)/2] >= target {
		return bisec(nums, target, left, (left+right)/2)
	}

	return bisec(nums, target, (left+right+1)/2, right)
}


// depreciated
func searchHelper(nums []int, target int, head, tail int) int {
	mid := (head+tail)/2
	fmt.Println(head, mid, tail)
	if target == nums[head] {
		return head
	}

	if len(nums[head:tail]) <= 1 {
		return -1
	}

	vsHead := target < nums[head]
	vsMid := target < nums[(head+tail)/2]

	if (vsHead && vsMid) || (!vsHead && !vsMid){
		return searchHelper(nums, target, (head+tail)/2, tail)
	}

	return searchHelper(nums, target, head, (head+tail)/2)
}