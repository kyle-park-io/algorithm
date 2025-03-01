package main

import "fmt"

var (
	l     = 10
	input = [10]int{59, 28, 26, 40, 51, 24, 44, 23, 48, 86}
)

func main() {

	for i := 1; i < l; i++ {
		// swap
		for j := i; j >= 1; j-- {
			if input[j] < input[j-1] {
				// tuple
				input[j-1], input[j] = input[j], input[j-1]
			} else {
				// break routine
				break
			}
		}

		// shift
		value := input[i]
		j := i - 1
		for j >= 0 && input[j] > value {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = value
	}

	fmt.Println(input)
}
