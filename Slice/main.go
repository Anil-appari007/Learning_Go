package main

import "fmt"

func main() {
	slice1 := []int{1, 3, 5, 7}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := [10]int{1, 3}
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	arr := [5]int{1, 3, 4}
	slice3 := arr[:]
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]int, 1, 3)
	fmt.Println(slice4)
	fmt.Println(slice4[0])
	slice4[0] = 3
	slice4 = append(slice4, 98)
	slice4 = append(slice4, 876, 234)
	fmt.Println(slice4)

	slice4 = append(slice4, slice1...)
	fmt.Println(slice4)
}
