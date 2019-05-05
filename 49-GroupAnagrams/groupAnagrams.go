/*
	Runtime: 2824 ms, faster than 5.03% of Go online submissions for Group Anagrams.
	Memory Usage: 341 MB, less than 7.69% of Go online submissions for Group Anagrams.
	https://leetcode.com/submissions/detail/226876931/
	too slow
 */
package problem49


import "fmt"

func groupAnagrams(strs []string) [][]string {
	var ret [][]string
	var bigMap = make(map[string]map[int32]int)
	var flag = false
	for _, str := range strs {
		strMap := convert2Map(str)
		for k, group := range ret {
			if compareMap(strMap, bigMap[group[0]]) {
				flag = true
				ret[k] = append(ret[k], str)
				break
			}
		}
		fmt.Println(str, flag)
		if !flag {
			bigMap[str] = strMap
			ret = append(ret, []string{str})
		}
		flag = false
	}
	return ret
}

func convert2Map(str string) map[int32]int {
	ret := make(map[int32]int)
	for _, s := range str {
		if _, ok := ret[s]; ok {
			ret[s] = ret[s]+1
		} else {
			ret[s] = 1
		}
	}
	return ret
}

func compareMap(m1, m2 map[int32]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k:= range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}

