package main

import "fmt"

func Varidic(a int, b ...int) {
	fmt.Println(a)
	fmt.Println(b)
	for _, each := range b {
		fmt.Println(each)
	}
}

func main() {
	Varidic(1, 3)
	Varidic(1, 3, 4, 5)
}
