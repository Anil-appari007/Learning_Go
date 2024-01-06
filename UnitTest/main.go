package main

import "fmt"

func addTwoNum(a int, b int) (res int) {
	res = a + b
	return res
}

func subTwoNum(a int, b int) (res int) {
	res = a - b
	return res
}

func main() {
	a1 := addTwoNum(1, 3)
	fmt.Println(a1)

	a2 := subTwoNum(9, 5)
	fmt.Println(a2)
}
