package main

import "fmt"

func sayMyName(name string) {
	fmt.Println("Hello", name)
}

func returnMyName(name string) (reverseName string) {
	fmt.Println("Hello", name)
	reverseName = "randomdata"
	return
}

func main() {
	sayMyName("user1")
	sayMyName("Heisenberg")
	sayMyName("Maximus Decimus Meridius")

	test := returnMyName("test")
	fmt.Println(test)
}
