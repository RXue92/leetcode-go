package main

import "fmt"


type student struct {
	Name string
	Age int
}

//func (s student) String() string{
//	return fmt.Sprintf("Name=%s, Age=%d", s.Name, s.Age)
//}

func problem1() {
	m := make(map[string]*student)
	stus := []student{
		{Name:"li", Age:20},
		{Name:"zhou", Age:21},
		{"wang", 22},
	}
	for _, stu := range stus{
		m[stu.Name] = &stu
	}

	fmt.Println(m)
	// map[li:0xc00000c060 wang:0xc00000c060 zhou:0xc00000c060]
	// the value of pointer to stu is the last element in slice stus
	fmt.Println(stus[0])
}

func calc(index string, a, b int) int {
	fmt.Println(index, a, b, a+b)
	return a+b
}

func problem2() {
	a := 0
	b := 1
	defer calc("10", a, calc("20", a, b))
	fmt.Println(123)
	a = 2
	defer calc("30", a, calc("40", a, b))
	b = 3
	fmt.Println(456)

	/*
	20 0 1 1
	123
	40 2 1 3
	456
	30 2 3 5
	10 0 1 1
	 */
}


type people struct {
	Name string
}

func (p *people) ShowA() {
	fmt.Println("a")
}

func (p *people) ShowB() {
	fmt.Println("b")
}

type teacher struct {
	people
	subject string
}

func (t * teacher) ShowA() {
	fmt.Println("teacher a")
}

func problem3() {
	// https://hackthology.com/object-oriented-inheritance-in-go.html
	t1 := teacher{people{"wang"}, "math"}
	t1.ShowA()
	// teacher a
}

func main() {
	problem1()
}