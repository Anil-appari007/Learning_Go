package main

import "fmt"

func main() {
	a := 3
	b := 5

	if a == b {
		fmt.Println("a equals b")
	} else {
		fmt.Println("A not equals b")
	}

	c := 9

	// Multiple conditions with &&
	if (a > b) && (a > c) {
		fmt.Println("a is greater than both b & c")
	} else if (a < b) && (a < c) {
		fmt.Println("a is less than both b & c")
	}

	// nested if

	if a < c {
		fmt.Println("a is lt c")
		if a > b {
			fmt.Println("a is gt b")
		}
	}
}
