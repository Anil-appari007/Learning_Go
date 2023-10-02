package main

import "fmt"

func main() {

	a := "teststring"
	// loop through the elements in string
	for index, value := range a {
		fmt.Printf("index_%v->Value_%c\n", index, value)
	}

	var i int = 1
	for i = 2; i < 5; i++ {
		fmt.Println(i)
	}
	// start with first condition and run until second condition reached
	// increment 1 for every execution
	for b := 0; b < 7; b++ {
		fmt.Println(b)
	}

	var testarray1 []int = []int{1, 2, 3, 4}
	for each_array_index, array_element := range testarray1 {
		fmt.Printf("Array_index_%v-Array_element_%v\n", each_array_index, array_element)
	}

	var testmap1 = map[string]int{"apples": 1, "banana": 2}
	for key := range testmap1 {
		fmt.Println(key)
	}
	for each_key, each_value := range testmap1 {
		fmt.Printf("Map_Key_%v-Value_%v\n", each_key, each_value)
	}
}
