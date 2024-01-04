package main

import "fmt"

func main() {
	var name string = "user1"
	var name2 *string = &name

	fmt.Println("name: ", name)
	fmt.Println("name2: ", name2)
	fmt.Println("name2: ", *name2)

	*name2 = "aaaaa"
	fmt.Println("changed name2: ", *name2)
	fmt.Println("name: ", name)

	name = "user2"
	fmt.Println(name, *name2)

	a := 25
	fmt.Println("a:", a)
	doubleTheValue(&a)
	fmt.Println("a:", a)
}

func doubleTheValue(input *int) {
	*input = *input * 2
}
