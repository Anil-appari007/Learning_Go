package main

import "fmt"

const pi = 3.14

func main() {
	fmt.Println("Value of Pi:", pi)

	// pi = 31.5	// Basics\Constants\main.go:10:2: cannot assign to pi (neither addressable nor a map index expression)
	fmt.Println("Value of Pi:", pi)

	// Multiple constans declaration
	const (
		g     = 9.81
		onekm = "1000m"
	)

	// Naming Convention
	// constants are usually written in uppercase letters
	const GRAVITY float32 = 9.8
}
