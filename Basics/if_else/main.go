package main

import "fmt"

func main() {
	var a int = 2
	var b int = 20
	if a > b {
		fmt.Println("a > b")
	}

	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a<b")
	}

	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("else condition running")
	}

	if a == 2 && b == 20 {
		fmt.Println("Both are true")
	}
	if a == 2 || b == 22 {
		fmt.Println("only one is true")
	}
}
