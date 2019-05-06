/*
	Failed multiple times!!!
	Runtime: 16 ms, faster than 29.69% of Go online submissions for Merge Intervals.
	Memory Usage: 6.5 MB, less than 9.52% of Go online submissions for Merge Intervals.
	https://leetcode.com/submissions/detail/227066582/
 */

package problem56

import "fmt"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	var ret = intervals[:1]
	for i:=1; i<len(intervals); i++ {
		//fmt.Println(ret)
		ret = mergeInto(ret, intervals[i])
	}

	fmt.Println(ret)
	return removeDup(ret)
}

// dst[i][0]<= dst[i+1][0] && dst[i][1] <= dst[i+1][1]
func mergeInto(dst [][]int, src []int) [][]int {
	//fmt.Println(dst, src)

	for i := range dst {
		// too large
		if src[0] > dst[i][1] {
			if i == len(dst)-1 {
				dst = append(dst, src)
			}
			continue
		}

		// concat
		if src[0] == dst[i][1] {
			dst[i] = []int{dst[i][0], src[1]}
			break
		}

		if src[1] == dst[i][0] {
			dst[i] = []int{src[0], dst[i][1]}
			break
		}

		// not overlapped, insert
		if src[1] < dst[i][0] {
			if i==0 || src[0] > dst[i-1][1] {
				dst = append(dst[:i], append([][]int{src}, dst[i:]...)...)
			}
			break
		}

		// overlapped
		//head, tail := src[0], src[1]
		//if dst[i][0] < head {
		//	head = dst[i][0]
		//}
		//if dst[i][1] > tail {
		//	tail = dst[i][1]
		//}

		dst[i], _ = mergeTwo(src, dst[i])
		//break

	}

	return dst
}


func removeDup(target [][]int) [][]int {
	for i:=0; i<len(target)-1; i++ {
		//if target[i][0] == target[i+1][0] {
		//	target = append(target[:i], target[i+1:]...)
		//	return removeDup(target)
		//}
		//if target[i][1] == target[i+1][1] {
		//	target = append(target[:i+1], target[i+2:]...)
		//	return removeDup(target)
		//}
		//if target[i][1] == target[i+1][0] {
		//	target[i+1][0] = target[i][0]
		//	target = append(target[:i], target[i+1:]...)
		//	return removeDup(target)
		//}
		if out, overlapped := mergeTwo(target[i], target[i+1]); overlapped {
			target[i+1] = out
			target = append(target[:i], target[i+1:]...)
			return removeDup(target)
		}
	}
	return target
}

func mergeTwo(s1, s2 []int) ([]int, bool) {
	if (s1[0] <= s2[1] && s1[1] >= s2[0]) || (s2[0] <= s1[1] && s2[1] >= s1[0]) {
		head, tail := s1[0], s1[1]
		if s2[0] < head {
			head = s2[0]
		}
		if s2[1] > tail {
			tail = s2[1]
		}

		return []int{head, tail}, true
	}
	return []int{}, false
}