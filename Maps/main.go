package main

import "fmt"

func main() {
	map1 := map[string]string{"name": "test", "data": "1234"}
	fmt.Println(map1)
	fmt.Println(map1["name"])

	var map2 = map[string]int{"id": 1, "data": 1111111}
	fmt.Println(map2)
	fmt.Println(map2["data"])

	emptyMap := map[string]int{}
	fmt.Println(emptyMap)
	fmt.Println(len(emptyMap))

	makeMap1 := make(map[string]int)
	fmt.Println(makeMap1)
	fmt.Println(len(makeMap1))

	makeMap1["id"] = 1
	makeMap1["data"] = 222222222
	fmt.Println(makeMap1)

	delete(makeMap1, "data")
	fmt.Println(makeMap1)

	makeMap1["newdata"] = 44444
	fmt.Println(makeMap1)

	for key, value := range makeMap1 {
		fmt.Println(key, "=", value)
	}
	//
}
