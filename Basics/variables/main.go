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

	// declaring multiple variables
	e, f, g, h := "test", 99, 33.32, false
	fmt.Println(e, f, g, h)

	// wrong way of declaring variables
	// 23var := "variable name should not start with number"
	// 24 test2 := "variable name shoud not have space"
	multiWord := "variable multiple words should be in camel case"
	fmt.Print(multiWord)

	car := "car"
	Car := "Car"
	CAR := "CAR"
	fmt.Println(car, Car, CAR)

	// zero value variables
	var zeroVar1 string
	fmt.Println(zeroVar1)

	var zeroVar2 int
	fmt.Println(zeroVar2) // 0

	var zeroVar3 float32
	fmt.Println(zeroVar3) // 0

	var zeroVar4 bool
	fmt.Println(zeroVar4) // false

	var (
		multiVar1 = "multiVar1"
		multiVar2 = 22222
		multiVar3 = 23.432
		multiVar4 = true
	)
	fmt.Printf("multiVar1: %v\nmultiVar2: %v\nmultiVar3:%v\nmultiVar4:%v", multiVar1, multiVar2, multiVar3, multiVar4)
}
