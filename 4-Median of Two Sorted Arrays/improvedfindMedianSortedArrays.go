/*
Based on reference https://leetcode.com/articles/median-of-two-sorted-arrays/
 */

package main

import (
	"fmt"
	"math"
)


var (
	nums01 = []int{1, 2}
	nums02 = []int{-1, 3}
)

func main() {
	ret := betterSolution(nums01, nums02)
	fmt.Println("median is: ", ret)
}

func betterSolution(nums1, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		return betterSolution(nums2, nums1)
	}

	// Find i in [0, l1] s.t. i+j=(l1+l2)/2 approximately, nums1[i-1] < nums2[j] and nums2[j-1] < nums1[i]
	iMin, iMax, j := 0, l1, 0
	// Start with i=l1/2, decrease or increase i according to following inequations
	i := (iMin + iMax) / 2

	for iMin <= iMax {
		j = (l1+l2+1)/2 - i
		if i < l1 && nums2[j-1] > nums1[i] {
			i++
		} else if i > 0 && nums1[i-1] > nums2[j] {
			i--
		} else {
			break
		}
	}
	fmt.Println(i, j)

	var leftMax, rightMin float64
	if j == 0 {
		leftMax = float64(nums1[i-1])
	} else if i == 0 {
		leftMax = float64(nums2[j-1])
	} else {
		leftMax = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
	}

	if (l1+l2)%2 == 1 {
		return leftMax
	}

	if i == l1 {
		rightMin = float64(nums2[j])
	} else if j == l2 {
		rightMin = float64(nums1[i])
	} else {
		rightMin = math.Min(float64(nums1[i]), float64(nums2[j]))
	}

	return (leftMax+rightMin)/2
}
