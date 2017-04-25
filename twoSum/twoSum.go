package main


import "fmt"



// beat 1.9%
//func twoSum(nums []int, target int) []int {
//	flag := 0
//	out := []int{}
//	for i, val1 := range(nums) {
//		for j, val2 := range(nums) {
//			if val1 + val2 == target && i != j {
//				flag = 1
//				out = []int{i, j}
//				break
//			}
//		}
//		if flag == 1 {
//			break
//		}
//	}
//
//	return out
//}


// beat 29.66%
func twoSum(nums []int, target int) []int{
	lap := 1
	out := []int{}
	for i, val := range(nums) {
		temp := nums[i+1:]
		if jd := contains(temp, target - val); jd == -1 {
			if len(temp) == 1 {
				break
			}
			lap ++
		} else {
			out = []int{i, jd+lap}
			return out
		}
	}

	return out
}

func contains(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}




func main() {
	//fmt.Print("hello")
	nums := []int{3, 2, 4, 1, 3}
	target := 6
	fmt.Print(twoSum(nums, target))
}

