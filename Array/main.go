package main

import "fmt"

func main() {
	var arr = [5]string{"a", "b", "c3", "4d"}
	fmt.Println(arr)

	fmt.Println(arr[2])
	fmt.Println(arr[0:1])
	fmt.Println(arr[1:])

	arr[0] = "aaa"
	fmt.Println(arr[0])
	fmt.Println(arr)

	fmt.Println("length of array arr", len(arr))
	fmt.Println("Capacity of arr", cap(arr))
}
