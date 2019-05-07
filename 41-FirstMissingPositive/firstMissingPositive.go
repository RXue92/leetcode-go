package problem41


func firstMissingPositive(nums []int) int {
	lb, ub, max := 1, 1, nums[0]
	localMax := nums[0]
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}

		if nums[i] == lb {
			if lb < ub {
				lb++
			} else {
				lb = localMax+1
				ub = lb
				localMax++
				continue
			}
		}
		if nums[i] <= ub {
			ub = nums[i] - 1
		} else {
			localMax = nums[i]
		}
	}

	return lb
}