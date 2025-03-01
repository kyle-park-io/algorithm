package main

import "fmt"

var (
	l     = 10
	input = [10]int{59, 28, 26, 40, 51, 24, 44, 23, 48, 86}
)

func main() {

	gap := l / 2

	for gap > 0 {
		for i := gap; i < l; i++ {
			// shift
			value := input[i]
			j := i
			for j-gap >= 0 && input[j-gap] > value {
				input[j] = input[j-gap]
				j -= gap
			}
			input[j] = value
		}

		gap /= 2
	}

	fmt.Println(input)
}
