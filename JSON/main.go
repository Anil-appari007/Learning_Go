package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Bird struct {
		Name string
		Id   int
	}
	birdJson := `{"name": "eagle","id": 11}`
	var eagle Bird
	json.Unmarshal([]byte(birdJson), &eagle)
	fmt.Println("Name:", eagle.Name)

	birdsJson := `[{"name":"pigeon","id":33},{"name":"peacock","id":22}]`
	var birds []Bird
	json.Unmarshal([]byte(birdsJson), &birds)
	fmt.Println(birds)
	for _, each := range birds {
		fmt.Println(each.Id)
		fmt.Println(each.Name)
	}

	// //
	type Dimension struct {
		Height int
		Weight int
	}
	type Bird1 struct {
		Name  string
		Color string
		Size  Dimension
	}

	newBirdsJson := `{"name":"eagle","color":"black","size":{"height":3,"weight":44}}`
	var newbird1 Bird1
	json.Unmarshal([]byte(newBirdsJson), &newbird1)
	fmt.Println(newbird1)
	fmt.Printf("name: %s", newbird1.Name)
	fmt.Println("\nColor:", newbird1.Color)
}
