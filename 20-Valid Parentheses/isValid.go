package main

import "fmt"

var (
	testCase="{[]}"
)

func main() {
	fmt.Println(isValid(testCase))
}

func isValid(s string) bool {
	fmt.Println(s)
	if len(s)%2 == 1 {
		return false
	}

	if s == "" {
		return true
	}


	for i:=0; i<len(s)-1; i++ {
		//fmt.Println(s[i:i+2])
		if s[i:i+2] == "()" || s[i:i+2] == "[]" || s[i:i+2] == "{}" {
			//temp := s[:i]+s[i+1:]
			//fmt.Println(temp)
			return isValid(s[:i]+s[i+2:])
		}
	}

	return false
}