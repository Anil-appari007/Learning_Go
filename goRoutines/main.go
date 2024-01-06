package main

import (
	"fmt"
	"time"
)

func hello(exec int) {
	fmt.Println("Hello From Go", exec)
}

func main() {
	go hello(1)
	fmt.Println("Hello")
	fmt.Println("Hello")
	go hello(2)

	fmt.Println("Hello")
	fmt.Println("Hello")
	fmt.Println("Hello")
	go hello(3)
	go hello(4)

	time.Sleep(5 * time.Second)
}
