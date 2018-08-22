// beats 15%.94%
func removeDuplicates(nums []int) int {
	k := len(nums)
	if k < 2 {
		return k
	}

	flag := nums[0]
	count := 0
	loc := 1
	var temp int

	// for i:= 1; i < len(nums); i++ {
	//     temp = nums[i]
	//     if temp > flag {
	//         flag = temp
	//         continue
	//     }
	//     count++
	// }
	// return k -  count

	for loc+count < k {
		temp = nums[loc]
		if temp > flag {
			flag = temp
			loc++
			continue
		}
		count++
		if loc == k-1 {
			nums = nums[:loc]
		} else {
			nums = append(nums[:loc], nums[loc+1:]...)
		}

		//nums[k:] is nil slice
		//nums = nums[:loc+copy(nums[loc:], nums[loc+1:])]
	}
	return k - count

}
