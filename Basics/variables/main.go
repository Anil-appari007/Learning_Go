package main

import "fmt"

func main() {
	// Declaring variable along with its type & assign a value
	var a string = "abc"
	fmt.Println("This is a string", a)

	var b int = 1
	fmt.Println("This is a integer", b)

	var c float32 = 2.06
	fmt.Println("This is a float", c)

	var d bool = true
	fmt.Println("This is a boolean", d)

	// Creating variable without mentioning its type & assign a value

	aa := "var_aa"
	fmt.Println("Variable aa declared with shorthand method", aa)
	fmt.Printf("Variable type is %T\n", aa)
	bb := 22
	fmt.Println("Variable bb declared with shorthand method", bb)
	fmt.Printf("Variable type is %T \n", bb)
	cc := 33.33
	fmt.Println("Variable cc declared with shorthand method", cc)
	fmt.Printf("Variable type is %T \n", cc)
	dd := false
	fmt.Println("Variable dd declared with shorthand method", dd)
	fmt.Printf("Variable type is %T \n", dd)

	// Declaring variable first without value
	// Assign value later
	var aaa string
	aaa = "aaa"
	fmt.Println(aaa)

	var bbb int
	bbb = 222
	fmt.Println(bbb)

	var ccc float32
	ccc = 333.33
	fmt.Println(ccc)

	var ddd bool
	ddd = true
	fmt.Println(ddd)
}
