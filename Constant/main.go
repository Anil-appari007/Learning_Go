package main

import "fmt"

func main() {
	const a = 1
	fmt.Println(a)
	// a = 133
	// Error:	Constant\main.go:8:2: cannot assign to a (neither addressable nor a map index expression)

	const (
		b = "b"
		c = 98.12
		d = false
	)
	fmt.Println(b, c, d)
}
