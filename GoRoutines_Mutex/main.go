package main

import (
	"fmt"
	"sync"
)

func sayHello(user string, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	fmt.Println("Hello from", user)

	mutex.Lock()
	*count = *count + 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	count := 0

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(1)
	go sayHello("user1", &count, &wg, &mutex)

	fmt.Println(count)

	arr := [5]string{"a1", "a2", "a3", "a4", "a5"}
	for _, each := range arr {
		wg.Add(1)
		go sayHello(each, &count, &wg, &mutex)
	}

	wg.Wait()
	fmt.Println(count)
}
