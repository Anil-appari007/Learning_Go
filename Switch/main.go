package main

import "fmt"

func main() {
	fmt.Println("SWitch")

	a := 1

	switch {
	case a == 0:
		fmt.Println("a is 0")
	case a == 9:
		fmt.Println("a is 9")
	case a == 1:
		fmt.Println("a is 1")
	default:
		fmt.Println("a is default")
	}

	switch a {
	case 0, 2:
		fmt.Println("Macthes 0 & 2")
	case 1, 4:
		fmt.Println("Matches 1 &4")
	}

}
