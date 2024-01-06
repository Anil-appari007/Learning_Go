package main

import (
	"fmt"
	"time"
)

func sayHello(exec int, status chan bool) {
	fmt.Println("Sleep 10")
	time.Sleep(10 * time.Second)
	fmt.Println("Hello from Go", exec)
	status <- true
}
func main() {
	status := make(chan bool)
	go sayHello(1, status)
	<-status

	// var newStatus bool
	// newStatus <- status
	go sayHello(2, status)
	<-status

	fmt.Println("Hello")
	// go sayHello(2, status)
	fmt.Println("Status:", status)
}
