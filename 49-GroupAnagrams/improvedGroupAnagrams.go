/*
	According to solution, use sorted string or equivalent as key of map.
	Runtime: 284 ms, faster than 74.83% of Go online submissions for Group Anagrams.
	Memory Usage: 340.6 MB, less than 7.69% of Go online submissions for Group Anagrams.
	https://leetcode.com/submissions/detail/226887541/
 */

package problem49

func betterSolution(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	if len(strs) == 1 {
		return [][]string{strs}
	}

	var ret = make([][]string, 0)
	var m = make(map[string]int)
	for _, str := range strs {
		// TODO: var index [26]byte, and skip making keys
		index := make([]int, 26)
		for i:=0; i<len(str); i++ {
			index[str[i] - 'a']++
		}

		// key = sort(string)
		var key string
		for i:=0; i<26; i++ {
			for j:=0; j<index[i]; j++ {
				key+= string(i+'a')
			}
		}

		if v, ok := m[key]; ok {
			ret[v] = append(ret[v], str)
		} else {
			m[key] = len(ret)
			ret= append(ret, []string{str})
		}
	}
	return ret
}