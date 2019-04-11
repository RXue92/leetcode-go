/*
	Backtracking with goroutine
	TODO: solve deadlock
 */
package main

import (
	"fmt"
	"sync"
)

var c = make(chan string)

func main() {
	digits := "023"
	out := betterSolution(digits)
	fmt.Println(out)
}

var dict = map[string][]string{
	"2":[]string{"a","b","c"},
	"3":[]string{"d","e","f"},
	"4":[]string{"g","h","i"},
	"5":[]string{"j","k","l"},
	"6":[]string{"m","n","o"},
	"7":[]string{"p","q","r","s"},
	"8":[]string{"t","u","v"},
	"9":[]string{"w","x","y","z"},
}

func betterSolution(digits string) []string {
	var tmp []string


	var wg sync.WaitGroup
	backtrack("", digits, &wg)
	wg.Add(1)
	go func() {
		for k := range c {
			tmp = append(tmp, k)
		}
	}()
	wg.Wait()
	return tmp
}

func backtrack(combinations string, lastDigits string, rwg *sync.WaitGroup) {
	if len(lastDigits) == 0 {
		c <- combinations
		return
	}

	digit := lastDigits[0:1]
	if digit == "0" || digit == "1" {
		backtrack(combinations, lastDigits[1:], rwg)
		return
	}

	for _, letter := range dict[digit] {
		rwg.Add(1)
		go func() {
			backtrack(combinations+letter, lastDigits[1:], rwg)
		}()
	}
	rwg.Wait()
}