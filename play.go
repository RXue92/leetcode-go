package main

import "fmt"



func main() {
	s1 := "你好"
	for i:=0; i<len(s1); i++ {
		fmt.Println(s1[i])
		fmt.Println(s1[i]==228)
	}
	fmt.Println(s1[:3]=="你")

	s2 := `你好`
	for _, c:= range s2 {
		fmt.Println(c)
		fmt.Println(c=='你')
	}

	fmt.Println(-1%2==-1, "last")

	s3 := "123"
	fmt.Println(s3[1:])
}
