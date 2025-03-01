package main

import (
	"fmt"

	"template/utils"
)

var (
	a = [10]int{}
	b = []int{}

	c = [10]int{1, 2, 3}
	d = []int{1, 2, 3}

	m = make(map[int]bool)
)

const (
	num = 100
	l   = 10
)

func main() {

	arr := make([]int, 0)
	c := 0

forLoop:
	for {
		if c == l {
			break forLoop
		}

		n := utils.GetRandom(num)

		if !m[n] {
			arr = append(arr, n)
			c++
		}
	}

	fmt.Println(arr)
}
