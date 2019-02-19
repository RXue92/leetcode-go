package main

import "fmt"

var (
	nums1 = []int{1, 2, 3}
	nums2 = []int{1, 2, 2}
)

func main() {
	ret := findMedianSortedArrays(nums1, nums2)
	fmt.Println("median is: ", ret)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)

	if l1 < l2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	if l2 == 0 {
		if l1%2 == 1 {
			return float64(nums1[l1/2])
		}
		return 0.5 * float64(nums1[l1/2]+nums1[l1/2-1])
	}

	stop := (l1 + l2) / 2
	fmt.Println("stop=", stop)
	// out is a sorted list that keeps elements from both list
	out := []int{}
	// tmp = # elements from l1 in out
	// count = length of out
	var count, tmp int
	for i := 0; i < l2; i++ {
		idx := findPlace(nums2[i], nums1[tmp:])
		fmt.Println(idx)
		// If idx is nonzero, then append the elements no bigger than nums2[i] from l1 into out
		if idx > 0 {
			out = append(out, nums1[tmp:tmp+idx]...)
		}
		out = append(out, nums2[i])
		tmp += idx
		count += idx + 1
		fmt.Println("count:", count)
		if count > stop {
			fmt.Printf("count=%d>stop=%d\n", count, stop)
			break
		}
	}

	// If count <= stop is true, then we need more elements from l1 to calculate the median.
	if count <= stop {
		out = append(out, nums1[count-l2:]...)
	}

	fmt.Println(out)

	if (l1+l2)%2 == 0 {
		return 0.5 * float64(out[stop-1]+out[stop])
	}

	return float64(out[stop])
}

// Given an integer i and a sorted list, return the number of element no bigger than i in the list
func findPlace(i int, nums []int) int {
	if nums[0] > i {
		return 0
	}

	var index int
	for index < len(nums) && nums[index] <= i {
		index++
	}

	return index
}
