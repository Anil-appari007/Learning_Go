package main

import "fmt"

func sayHello() {
	fmt.Println("Helllooooo")
}

// Function with params
func sayHelloUser(name string) {
	fmt.Printf("hello user: %v\n", name)
}

func addNumbers(x int, y int) int {
	sum := x + y
	return sum
	// fmt.Println(sum)
}

func checkNumberMatch(x int, y int) bool {
	check := x == y
	if check == true {
		fmt.Printf("Numbers %v & %v are same \n", x, y)
	} else if check == false {
		fmt.Printf("Numbers %v & %v are not same\n", x, y)
	}
	return check
}

// return data without mentioning anything with return
func noReturn() (r_var string) {
	r_var = "Value from No No return"
	return
}

// Anonymous function
var area = func(l int, b int) int {
	return l * b
}

func main() {
	sayHello()
	sayHelloUser("dev")

	fmt.Println("Sum of 2 numbers", addNumbers(1, 2))

	fmt.Println(checkNumberMatch(1, 3))
	fmt.Println(checkNumberMatch(3, 3))

	fmt.Println((noReturn()))

	fmt.Println(area(4, 5))
}
