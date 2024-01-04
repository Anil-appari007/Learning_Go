package main

import "fmt"

func main() {
	arr1 := [4]string{"a", "b", "c", "d"}
	for _, each := range arr1 {
		fmt.Println(each)
	}

	string1 := "ssssssss"
	for _, each := range string1 {
		fmt.Println(each)
	}
}
