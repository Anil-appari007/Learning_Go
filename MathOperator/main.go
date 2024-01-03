package main

import "fmt"

func main() {
	a := 2
	b := 4

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	// fmt.Println(a * *b)
	fmt.Println(a % b)

	a++
	fmt.Println(a)

	a--
	fmt.Println(a)

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)

}
