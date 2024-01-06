package main

import (
	"fmt"
	"sync"
)

func sayHello(exec string, wg *sync.WaitGroup) {

	fmt.Println("Hello From", exec)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello("r1", &wg)

	wg.Add(1)
	go sayHello("r2", &wg)

	wg.Add(1)
	go sayHello("r3", &wg)

	fmt.Println("Waiting for Routines to complete")
	wg.Wait()

	fmt.Println("Routines finished")
}
