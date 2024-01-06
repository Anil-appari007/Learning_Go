package main

import "fmt"

func customDefer(exec int) {
	fmt.Println("Defer testing", exec)
}
func main() {
	fmt.Println("Hello")
	defer customDefer(1)
	defer customDefer(2)

	fmt.Println("Hello22")
	defer customDefer(3)

}
