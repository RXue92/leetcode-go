package main

import (
	"fmt"
	_ "sort"
)

func main() {
	height := []int{9,6,8,8,5,6,3}
	fmt.Println("height", height)
	fmt.Println(trap(height))
}

func trap(height []int) int {

	volume := 0
	l := len(height)
	if l < 3 {
		return volume
	}

	height = append(height, 0)

	i, j := 0, 0
	localmax := []int{-1}

	for i < l-2 {
		//container := []int{}
		//find a container

		for height[i] <= height[i+1] {
			i++
		}
		j = i+2
		if k:= len(localmax); localmax[k-1] != i {
			localmax = append(localmax, i)
		}

		trend := -1 //downwards
		for i < l-2 && j <= l{
			if trend < 0 {
				if height[j] > height[j-1]{
					trend = 1
				}
				j++

			} else {
				if height[j] <= height[j-1]{
					localmax = append(localmax, j-1)
					break
				}
				j++
			}
		}
		i = j - 1
	}
	localmax = localmax[1:]
	fmt.Println("localmax", localmax)


	//calculate volume
	volume = vol(height, localmax)

	return volume
}

func vol(height []int, localmax []int) int {
	l := len(localmax)
	if l == 1 {
		return 0
	}

	volume := 0
	i := 0
	container := []int{}
	for j := 1; j < l; j++ {
		if height[localmax[i]] > height[localmax[j]] && !isbiggest(j, localmax, height) {
			continue
		} else {
			container = height[localmax[i]:localmax[j]+1]
			fmt.Println("one container:", container)
			volume =  volume + volofcontainer(container)
			i = j
		}
	}

	return volume
}

//index (slice) is localmax of data (slice), and j is an index of index
//judging if j correspond to globalamx of data since index[j]
func isbiggest(j int, index []int, data []int) bool {
	l := len(index)
	if j == l-1 {
		return true
	}
	max := data[index[j]]
	for _, m := range(index[j+1:]) {
		if max < data[m] {
			return false
		}
	}

	return true
}

func volofcontainer(container []int) int {
	l := len(container)

	top := container[l-1]
	if container[0] <= container[l-1] {
		top = container[0]
	}
	volume := 0
	for _, h := range container[1 : l-1] {
		if h < top {
			volume = volume + (top - h)
		}
	}
	return volume
}
