/*
Runtime: 4 ms, faster than 93.78% of Go online submissions for Trapping Rain Water.
Memory Usage: 2.8 MB, less than 38.68% of Go online submissions for Trapping Rain Water.
https://leetcode.com/submissions/detail/234023494/
"""
So, we can say that if there is a larger bar at one end (say right), we are assured that
the water trapped would be dependant on height of bar in current direction (from left to right).
As soon as we find the bar at other end (right) is smaller, we start iterating in opposite
direction (from right to left).
"""
 */

package problem42

func betterSolution(height []int) int {
	if len(height) < 1 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	var   area int
	for left < right {
		if height[left] >= height[right] {
			if height[right] > rightMax {
				rightMax = height[right]
			}
			area += rightMax - height[right]
			right--
		} else {
			if height[left] > leftMax {
				leftMax = height[left]
			}
			area += leftMax - height[left]
			left++
		}
	}
	return area
}