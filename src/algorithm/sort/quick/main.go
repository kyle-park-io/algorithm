package main

import (
	"fmt"

	"quick/method"
)

var (
	l     = 10
	input = []int{100, 28, 59, 26, 51, 24, 44, 23, 48, 86}
)

func main() {
	method.QuickSortHoare(input, 0, l-1)
	fmt.Println(input)

	method.QuickSortLomuto(input, 0, l-1)
	fmt.Println(input)
}
