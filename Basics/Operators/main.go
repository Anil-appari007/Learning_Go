package main

import "fmt"

func main() {
	a := 1
	b := 2

	var c = a + b
	fmt.Println("Sum is :", c)
	var d = a - b
	fmt.Println("diff is :", d)

	var e = a * b
	fmt.Println("multiply is :", e)

	var f = a / b
	fmt.Println("div is :", f)

	var i = a % b
	fmt.Printf("Modulus is %v\n", i)

	// increment
	var j = 23
	j++
	fmt.Printf("+ve increment: %v\n", j)

	var k = 44
	k--
	fmt.Printf("decrement: %v\n", k)

	// Comparison operators
	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	fmt.Println(1 == 1 && 1 == 2)
	fmt.Println(1 == 1 || 1 == 2)
	fmt.Println(!(1 == 1 && 1 == 2))

}
