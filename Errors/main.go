package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "README"
	file, err := os.Open(fileName) // change the fileName & run

	if err != nil {
		fmt.Println("File not found")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("File", file.Name(), "Found")
}
