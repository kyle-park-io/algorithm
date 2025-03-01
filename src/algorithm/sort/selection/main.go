package main

import "fmt"

var (
	l      = 10
	input  = [10]int{59, 28, 26, 40, 51, 24, 44, 23, 48, 86}
	answer = []int{}
)

// improvement
func sort2() {

	index := -1
	min := 10000
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if input[j] < min {
				min = input[j]
				index = j
			}
		}

		// tuple assignment
		input[i], input[index] = input[index], input[i]

		index = -1
		min = 10000
	}

	fmt.Println(input)
}

func sort() {
	index := 0
	value := 10000

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if input[j] != -1 && input[j] < value {
				index = j
				value = input[j]
			}
		}
		answer = append(answer, value)
		input[index] = -1

		index = 0
		value = 10000
	}

	fmt.Println(answer)
}

func main() {
	// sort()
	sort2()
}
