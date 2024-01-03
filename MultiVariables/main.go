package main

import "fmt"

func main() {
	a, b, c, d := "string-a", 1, 10.87, false

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	var (
		e = 2
		f = "f"
		g = 98.9
		h = true
	)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
}
