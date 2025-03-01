package main

import "fmt"

var (
	l     = 10
	input = [10]int{59, 28, 26, 40, 51, 24, 44, 23, 48, 86}
)

func main() {

	c := -1

forLoop:
	for {
		if c == 0 {
			break forLoop
		}

		c = 0
		for i := l - 1; i >= 1; i-- {
			if input[i] < input[i-1] {
				// tuple assignment
				input[i-1], input[i] = input[i], input[i-1]

				// // temp var
				// s := input[i-1]
				// input[i-1] = input[i]
				// input[i] = s

				c++
			}
		}
	}

	fmt.Println(input)
}
