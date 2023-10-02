package main

import (
	"fmt"
	"strconv"
)

func main() {
	var string1 string = "test1"
	fmt.Println(string1)
	// fmt.Println(strconv.Atoi(string1))	// 0 strconv.Atoi: parsing "test1": invalid syntax
	var string2 string = "100"
	fmt.Println(strconv.Atoi(string2))

	var int1 int = 123
	fmt.Printf("Type of var int1: %T\n", int1)

	string3 := strconv.Itoa(int1)
	fmt.Println(string3)
	fmt.Printf("Type of var string3 is: %T", string3)

	floatVar1 := 23.45
	fmt.Printf("type of floatVar1 is :%T \n", floatVar1)
	// string4 := strconv.Itoa(floatVar1) // cannot use floatVar1 (variable of type float64) as int value in argument to strconv.Itoa

	// string to bool
	string5 := "true"
	fmt.Printf("type of string5: %T", string5)

	bool1, _ := strconv.ParseBool(string5)
	fmt.Printf("Type of bool1 is : %T", bool1)
}
