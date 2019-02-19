package main

import (
	"fmt"
)

var (
	testCase1=[]string{"flower", "flow", "flight"}
	testCase2=[]string{}
	)

// Same as the reference solution.
func main() {
	prefix := longestCommonPrefix(testCase1)
	fmt.Println("output:", prefix)
	if prefix=="" {
		fmt.Println("Explanation: There is no common prefix among the input strings.")
	}
}

func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) == 0 {
		return prefix
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix = strs[0]
	for i:=1; i<len(strs); i++ {
		if prefix == "" {
			break
		}
		prefix = longestCommonPrefix42(prefix, strs[i])
	}


	return prefix
}

func longestCommonPrefix42(str1, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if l1 > l2 {
		l1 = l2
	}
	var prefix []byte
	for i:=0; i<l1;i++ {
		if str1[i] == str2[i] {
			prefix = append(prefix, str1[i])
		} else {
			break
		}
	}

	return string(prefix)
}