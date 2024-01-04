package main

import "fmt"

func main() {
	type CarsDetails struct {
		Name     string
		Model    int
		Power    int
		MaxSpeed int
	}

	var Verna CarsDetails
	Verna.Name = "test1"
	Verna.Model = 2024
	Verna.Power = 1500
	Verna.MaxSpeed = 210

	fmt.Println(Verna)
}
