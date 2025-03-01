package main

import (
	"merge/method"

	"fmt"
)

var (
	l     = 10
	input = []int{59, 28, 26, 40, 51, 24, 44, 23, 48, 86}
)

// Divide & Conquer
func main() {
	method.MergeSortRecursive(input, 0, len(input)-1)
	fmt.Println(input)

	method.MergeSortIterative(input)
	fmt.Println(input)
}
